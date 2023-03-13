package main

import (
	"embed"
	"fmt"
	"io/fs"
	"log"
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/go-playground/form"
	"github.com/justinas/nosurf"
	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"main.go/cms/handler"
	"main.go/utility"
)

//go:embed assets
var assetFiles embed.FS

//go:embed migrations
var migrationFiles embed.FS

func main (){
	config := viper.NewWithOptions(
		viper.EnvKeyReplacer(
			strings.NewReplacer(".", "_"),
		),
	)
	config.SetConfigFile("config")
	config.SetConfigType("ini")
	config.AutomaticEnv()
	if err := config.ReadInConfig(); err != nil {
		log.Fatal(err)
	}
	
	decoder := form.NewDecoder()

	postGresStore,err := utility.NewPostgresStorage(config)
	if err != nil {
		log.Fatalln(err)
	}
	goose.SetBaseFS(migrationFiles)
	if err := goose.SetDialect("postgres"); err != nil {
        log.Fatalln(err)
    }
	
    if err := goose.Up(postGresStore.DB.DB, "migrations"); err != nil {
        log.Fatalln(err)
    }


	lifeTime := config.GetDuration("session.lifeTime")
	idleTime := config.GetDuration("session.idleTime")
	sessionManager := scs.New()
	sessionManager.Lifetime = lifeTime * time.Hour
	sessionManager.IdleTimeout = idleTime * time.Minute
	sessionManager.Cookie.Name = "web-session"
	sessionManager.Cookie.HttpOnly = true
	sessionManager.Cookie.Secure = true
	sessionManager.Store = utility.NewSQLXStore(postGresStore.DB)

	var assetFS = fs.FS(assetFiles)
	staticFiles, err1 := fs.Sub(assetFS, "assets/src")
	if err1 != nil {
		log.Fatal(err)
	}

	templateFiles, err := fs.Sub(assetFS, "assets/templates")
	if err != nil {
		log.Fatal(err)
	}
	
    usermgmUrl := config.GetString("usermgm.url")
	usermgmConn,err := grpc.Dial(usermgmUrl,grpc.WithInsecure())
	if err != nil {
		log.Println("the error is in usermgmConn")
		log.Fatalln(err)
	}

	chi := handler.NewHandler(sessionManager,decoder,usermgmConn,staticFiles,templateFiles)
	newChi := nosurf.New(chi)

	p := config.GetInt("server.port")
	lis , err := net.Listen("tcp",fmt.Sprintf(":%d",p))
	if err != nil {
		fmt.Println("the error is in the main function of cms after server config")
		log.Fatalf("unable to listern port %#v",err)
	}

	fmt.Println("cms server running on :",lis.Addr())

	if err := http.Serve(lis,newChi);err != nil {
		log.Fatalf("unable to serve port %#v",err)
	}

}