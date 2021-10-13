package main

import (
	"bytes"
	"github.com/matryer/is"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestGetHomePage(t *testing.T) {
	srv := iniTestServer()
	var buf bytes.Buffer
	req, err := http.NewRequest(http.MethodGet, "/", &buf)
	is := is.New(t)
	is.NoErr(err)
	w := httptest.NewRecorder()
	srv.ServeHTTP(w, req)
	is.Equal(w.Result().StatusCode, http.StatusOK)
	out, err := ioutil.ReadAll(w.Result().Body)
	if err != nil {
		t.Fatal("error in reading response")
	}
	is.True(strings.Contains(string(out), "This is our first welcome page"))
}
