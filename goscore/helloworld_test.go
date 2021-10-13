package main

import (
	"bytes"
	"github.com/matryer/is"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestHelloWorld(t *testing.T){
	var buf bytes.Buffer
	srv:= &scoreServer{
		version: "",
		router: http.NewServeMux(),
		logger: log.New(os.Stdout, "logger: ", log.Lshortfile),
	}
	srv.Routes()
	req, err := http.NewRequest(http.MethodGet, "/", &buf)
	is:= is.New(t)
	is.NoErr(err)
	w:= httptest.NewRecorder()
	srv.ServeHTTP(w, req)
	is.Equal(w.Result().StatusCode, http.StatusOK)
}