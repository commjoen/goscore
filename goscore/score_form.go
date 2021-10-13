package main

import (
	"bytes"
	"log"
	"net/http"
	"strings"
)

func (s scoreServer) handleScoreSubmit() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method!= http.MethodPost {
			http.Error(w, "404: only posts live here", http.StatusNotFound)
			return
		}
		if err := r.ParseForm(); err != nil {
			log.Fatal(w, "ParseForm() err: %v", err)
			return
		}
		value:= strings.ToLower(r.FormValue("scoreField"))
		switch value {
		case "alfalfa":
			w.Write(bytes.NewBuffer("This is their favorite food, not the solution"))
		default:
			w.Write(bytes.NewBuffer("Try harder"))
		}
	}
}
