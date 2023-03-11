package utility

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
)

type PostGressStorage struct {
	DB *sqlx.DB
}

func NewPostgresStorage(cfg *viper.Viper)(*PostGressStorage,error) {
	db,err := ConnectDatabase(cfg)
	if err != nil {
		log.Fatal(err)
		return nil,err
	}
	return &PostGressStorage{
		DB: db,
	},nil
}

func ConnectDatabase(cfg *viper.Viper) ( *sqlx.DB,error){
    db, err := sqlx.Connect("postgres", fmt.Sprintf("host=%s port=%s user= %s password=%s dbname=%s sslmode=%s",
		cfg.GetString("database.host"),
		cfg.GetString("database.port"),
		cfg.GetString("database.user"),
		cfg.GetString("database.password"),
		cfg.GetString("database.dbname"),
		cfg.GetString("database.sslmode"),
	))
	if err != nil {
		log.Fatal(err)
		return nil,err
	}
	return db,nil	
}