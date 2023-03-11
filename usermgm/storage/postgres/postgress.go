package postgres

import (
	"fmt"
	"log"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
)
const (
  NotFound = "sql: no rows in result set"
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
func NewTestStorage(dbstring string, migrationDir string) (*PostGressStorage, func()) {
	db, teardown := MustNewDevelopmentDB(dbstring, migrationDir)
	db.SetMaxOpenConns(5)
	db.SetConnMaxLifetime(time.Hour)

	return &PostGressStorage{
		DB: db,
	}, teardown
}