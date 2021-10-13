package main

import (
	"bytes"
	"github.com/matryer/is"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
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
	out, err := ioutil.ReadAll(w.Result().Body)
	if err != nil {
		t.Fatal("error in reading response")
	}
	is.True(strings.Contains(string(out), "This is our first welcome page, let's see if it works"))
}