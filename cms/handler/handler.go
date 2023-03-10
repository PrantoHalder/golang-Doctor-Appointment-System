package handler

import (
	"html/template"
	"io/fs"
	"log"
	"net/http"
	"strconv"

	"github.com/Masterminds/sprig"
	"github.com/alexedwards/scs/v2"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-playground/form"
	"google.golang.org/grpc"
	userpb "main.go/gunk/v1/user"
)



type Handler struct {
	sessionManager *scs.SessionManager
	decoder        *form.Decoder
	usermgmService  usermgmService
	Templates      *template.Template
	staticFiles    fs.FS
	templateFiles  fs.FS
}

type usermgmService struct {
	userpb.UserServiceClient
}

func NewHandler(sm *scs.SessionManager, formDecoder *form.Decoder, usermgmConn *grpc.ClientConn, staticFiles, templateFiles fs.FS) *chi.Mux {
	h := &Handler{
		sessionManager: sm,
		decoder:        formDecoder,
		usermgmService: usermgmService{userpb.NewUserServiceClient(usermgmConn)},
		staticFiles:    staticFiles,
		templateFiles:  templateFiles,
	}

	h.ParseTemplates()
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Group(func(r chi.Router) {
		r.Use(sm.LoadAndSave)
		
	})

	r.Handle("/static/*", http.StripPrefix("/static", http.FileServer(http.FS(h.staticFiles))))
    

	r.Group(func(r chi.Router) {
		r.Use(sm.LoadAndSave)
		r.Get("/", h.MainHome)
		r.Get("/login", h.Login)
		r.Post("/loginpost", h.LoginPost)
		r.Get("/register", h.Register)
		r.Post("/registerpost", h.RegisterPost)
	})

	r.Route("/patients", func(r chi.Router) {
		r.Use(sm.LoadAndSave)
		r.Use(h.Authentication)
		r.Get("/home", h.PatientHome)
	})
	r.Route("/facultys", func(r chi.Router) {
		r.Use(sm.LoadAndSave)
		r.Use(h.Authentication)
		
	})
	r.Route("/students", func(r chi.Router) {
		r.Use(sm.LoadAndSave)
		r.Use(h.Authentication)
		

	})
	r.Group(func(r chi.Router) {
		r.Use(sm.LoadAndSave)
		r.Use(h.Authentication)

	})

	return r
}
func (h Handler) Authentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userID := h.sessionManager.GetString(r.Context(), "userID")
		uID, err := strconv.Atoi(userID)
		if err != nil {
			log.Fatalln(err)
		}
		if uID <= 0 {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
		next.ServeHTTP(w, r)
	})
}
func (h *Handler) ParseTemplates() error {
	templates := template.New("web-templates").Funcs(template.FuncMap{
		"calculatePreviousPage": func(currentPageNumber int) int {
			if currentPageNumber == 1 {
				return 0
			}

			return currentPageNumber - 1
		},

		"calculateNextPage": func(currentPageNumber, totalPage int) int {
			if currentPageNumber == totalPage {
				return 0
			}

			return currentPageNumber + 1
		},
	}).Funcs(sprig.FuncMap())

	tmpl := template.Must(templates.ParseFS(h.templateFiles, "*.html,*/*.html"))
	if tmpl == nil {
		log.Fatalln("unable to parse templates")
	}

	h.Templates = tmpl
	return nil
}
