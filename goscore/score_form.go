package main

import (
	"log"
	"net/http"
	"strings"
)

func (s scoreServer) handleScoreSubmit() http.HandlerFunc {
	s.logger.Printf("Parsing goscore form here")
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "404: only posts live here", http.StatusNotFound)
			return
		}
		if err := r.ParseForm(); err != nil {
			log.Fatal(w, "ParseForm() err: %v", err)
			http.Error(w, "Error parsing form", http.StatusInternalServerError)
			return
		}
		value := strings.ToLower(r.FormValue("formValue"))
		s.logger.Printf("we got response for scoring, value: " + value)
		switch value {
		case "alfalfa":
			w.Write([]byte("This is their favorite food, not the solution"))
		default:
			w.Write([]byte("Try harder"))
		}
	}
}
