package main

import (
	"log"

	"github.com/Mattcazz/Fantasy.git/db"
)

func main() {

	db := db.ConnectDB()

	server := NewAppServer(":8080", db)
	if err := server.Run(); err != nil {
		log.Fatal("Error en establecer la conexion")
	}
}
