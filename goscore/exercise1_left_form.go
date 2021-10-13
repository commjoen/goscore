package main

import (
	"log"
	"net/http"
	"strings"
)

func (s scoreServer) handleLeftExercise1() http.HandlerFunc {
	s.logger.Printf("Parsing exercise1 left form here")
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
		s.logger.Printf("we got response for exercise value 1, value: " + value)
		w.Write([]byte(value))
	}
}
