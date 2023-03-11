package main

import (
	"embed"
	"fmt"
	"log"
	"net"
	"strings"

	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	userpb "main.go/gunk/v1/user"
	"main.go/usermgm/service/user"
	cu "main.go/usermgm/core/user"
	"main.go/usermgm/storage/postgres"
)

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
		fmt.Println("the error is in the main function of usermgm after ReadInConfig")
		log.Fatalf("error loading configuration %#v",err)
	}
	p := config.GetInt("server.port")
	lis , err := net.Listen("tcp",fmt.Sprintf(":%d",p))
	if err != nil {
		fmt.Println("the error is in the main function of usermgm after server config")
		log.Fatalf("unable to listern port %#v",err)
	}
    postGresStore,err := postgres.NewPostgresStorage(config)
	if err != nil {
		log.Fatalln(err)
	}
	goose.SetBaseFS(migrationFiles)
	if err := goose.SetDialect("postgres"); err != nil {
		log.Fatalln(err)
	}
	if err := goose.Up(postGresStore.DB.DB,"migrations");err != nil{
		log.Fatalln(err)
	}
	
    grpcServer := grpc.NewServer()
    userCore := cu.NewCoreUser(postGresStore)
	userSvc:= user.NewUserSvc(userCore)
	userpb.RegisterUserServiceServer(grpcServer,userSvc)
	

	fmt.Println("usermgm server running on :",lis.Addr())

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("unable to serve %#v",err)
	}
}