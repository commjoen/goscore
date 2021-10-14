package main

import (
	"log"
	"net/http"
	"os"
)

func iniTestServer() *scoreServer {
	srv := &scoreServer{
		version: "",
		router:  http.NewServeMux(),
		logger:  log.New(os.Stdout, "logger: ", log.Lshortfile),
	}
	srv.Routes()
	return srv
}
