package main

import (
	"bytes"
	"github.com/matryer/is"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPostScoreBoard(t *testing.T) {
	srv := iniTestServer()
	buf := bytes.NewBuffer([]byte("formValue=test"))
	req, err := http.NewRequest(http.MethodPost, "/goscore", buf)
	isVerifier := is.New(t)
	isVerifier.NoErr(err)
	w := httptest.NewRecorder()
	srv.ServeHTTP(w, req)
	isVerifier.Equal(w.Result().StatusCode, http.StatusOK)
	out, err := ioutil.ReadAll(w.Result().Body)
	if err != nil {
		t.Fatal("error in reading response")
	}
	isVerifier.Equal(string(out), "Try harder")
}
