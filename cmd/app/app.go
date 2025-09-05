package main

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

type APPServer struct {
	addr string
	db   *sql.DB
}

func NewAppServer(addr string, db *sql.DB) *APPServer {
	return &APPServer{
		addr: addr,
		db:   db,
	}
}

// here we can use gin or mux or anything else.

func (s *APPServer) Run() error {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusAccepted,
			gin.H{"message": "Hello world!"})
	})

	return r.Run(s.addr)
}
