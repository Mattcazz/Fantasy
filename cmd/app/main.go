package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Mattcazz/Fantasy.git/db"
)

func main() {

	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("PG_USER")
	dbPassword := os.Getenv("PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", dbUser, dbPassword, dbHost, dbPort, dbName) //postgres://USER:PASSWORD@HOST:PORT/DATABASE?OPTIONS

	db, err := db.New(dsn, 30, 30, "15m")

	if err != nil {
		log.Fatal(err)
	}

	server := NewAppServer(":8080", db)
	if err := server.Run(); err != nil {
		log.Fatal("Error en establecer la conexion")
	}
}
