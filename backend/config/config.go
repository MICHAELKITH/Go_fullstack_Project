package config

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

var DBPool *pgxpool.Pool
func InitializeDB( dsn string){
	var err error
	DBPool, err = pgxpool.New(context.Background(), dsn)
	if err != nil{
		log.Fatalf("Unable to connect to the database : %v\n", err)
	}
	log.Println("Database connection is established")
}

func CloseDB(){
	if DBPool != nil{
		DBPool.Close()
	}
}