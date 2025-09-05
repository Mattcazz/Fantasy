package main

import "database/sql"

func main() {

	var db *sql.DB = nil

	server := NewAppServer(":8080", db)
	server.Run()
}
