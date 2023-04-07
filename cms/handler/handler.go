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
	adminpb "main.go/gunk/v1/admin"
	doctorpb "main.go/gunk/v1/doctor"
	doctortypepb "main.go/gunk/v1/doctortype"
	loginpb "main.go/gunk/v1/login"
	userpb "main.go/gunk/v1/user"
)

type Handler struct {
	sessionManager *scs.SessionManager
	decoder        *form.Decoder
	usermgmService usermgmService
	Templates      *template.Template
	staticFiles    fs.FS
	templateFiles  fs.FS
}

type usermgmService struct {
	userpb.UserServiceClient
	adminpb.AdminServiceClient
	doctorpb.DoctorServiceClient
	doctortypepb.DoctorTypeServiceClient
	loginpb.LoginServiceClient
}

func NewHandler(sm *scs.SessionManager, formDecoder *form.Decoder, usermgmConn *grpc.ClientConn, staticFiles, templateFiles fs.FS) *chi.Mux {
	h := &Handler{
		sessionManager: sm,
		decoder:        formDecoder,
		usermgmService: usermgmService{userpb.NewUserServiceClient(usermgmConn),
			adminpb.NewAdminServiceClient(usermgmConn),
			doctorpb.NewDoctorServiceClient(usermgmConn),
			doctortypepb.NewDoctorTypeServiceClient(usermgmConn),
			loginpb.NewLoginServiceClient(usermgmConn)},
		staticFiles:   staticFiles,
		templateFiles: templateFiles,
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
		r.Get("/inactive", h.Inactive)
		r.Post("/registerpost", h.RegisterPost)
	})

	r.Route("/patients", func(r chi.Router) {
		r.Use(sm.LoadAndSave)
		r.Use(h.Authentication)
		r.Get("/{id:[0-9]+}/home", h.PatientHome)
		r.Get("/{id:[0-9]+}/showdoctortype", h.ShowDoctorType)
		r.Get("/{id:[0-9]+}/showAppointmentstatus", h.AppointmentStatus)
		r.Get("/{id:[0-9]+}/fixAppointment", h.FixAppointment)
		r.Get("/{id:[0-9]+}/searcdoctors", h.ShowDoctorPatient)
		r.Get("/logout", h.LogoutPatienthandler)
	})
	r.Route("/admin", func(r chi.Router) {
		r.Use(sm.LoadAndSave)
		r.Use(h.Authentication)
		r.Get("/home", h.AdminHome)
		r.Get("/registerpatient", h.PatientRegister)
		r.Get("/registeradmin", h.AdminRegister)
		r.Get("/registerdoctor", h.DoctorRegister)
		r.Get("/registerdoctortype", h.DoctorTypeRegister)
		r.Post("/registerpatientpost", h.PatientRegisterPost)
		r.Post("/registeradminpost", h.AdminRegisterPost)
		r.Post("/registerdoctorpost", h.DoctorRegisterPost)
		r.Post("/registerdoctortypepost", h.DoctorTypeRegisterPost)
		r.Get("/showpatient", h.Show)
		r.Get("/showadmin", h.ShowAdmin)
		r.Get("/showdoctor", h.ShowDoctor)
		r.Get("/showdoctortype", h.ShowDoctorType)
		r.Get("/logout", h.LogoutAdminhandler)
		r.Get("/{id:[0-9]+}/inputdeatails", h.InputDoctorDeatails)
		r.Post("/{id:[0-9]+}/InputDoctorDeatailspost", h.InputDoctorDeatailspost)
		r.Get("/{id:[0-9]+}/listDoctorDeatails", h.ListDoctorDetails)
		r.Get("/{id:[0-9]+}/editpatient", h.EditPatient)
		r.Get("/{id:[0-9]+}/userstatusedit", h.EditUserStatus)
		r.Get("/{id:[0-9]+}/editdoctortype", h.EditDoctorType)
		r.Get("/{id:[0-9]+}/editadmin", h.EditAdmin)
		r.Get("/{id:[0-9]+}/editdoctor", h.EditDcotor)
		r.Post("/{id:[0-9]+}/updatepatient", h.UpdatePatient)
		r.Post("/{id:[0-9]+}/updateadmin", h.UpdateAdmin)
		r.Post("/{id:[0-9]+}/updatedoctor", h.UpdateDoctor)
		r.Post("/{id:[0-9]+}/updatedoctortype", h.UpdateDoctorType)
		r.Get("/{id:[0-9]+}/deletepatient", h.DeletePatient)
		r.Get("/{id:[0-9]+}/deleteadmin", h.DeleteAdmin)
		r.Get("/{id:[0-9]+}/deletedoctor", h.DeleteDoctor)
		r.Get("/{id:[0-9]+}/deletedoctortype", h.DeleteDoctorType)
	})
	r.Route("/doctor", func(r chi.Router) {
		r.Use(sm.LoadAndSave)
		r.Use(h.Authentication)
		r.Get("/{id:[0-9]+}/home", h.DoctorHome)
		r.Get("/{id:[0-9]+}/manageschedule",h.MangeSchedule)
		r.Get("/{id:[0-9]+}/listschedule",h.ListSchedule)
		r.Post("/{id:[0-9]+}/manageschedulepost",h.MangeSchedulePost)
		r.Get("/{id:[0-9]+}/appointmentlist",h.AppointmentList)
		r.Get("/logout", h.LogoutDoctorhandler)
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

	tmpl := template.Must(templates.ParseFS(h.templateFiles, "*.html"))
	if tmpl == nil {
		log.Fatalln("unable to parse templates")
	}

	h.Templates = tmpl
	return nil
}
