package main

import (
	"log"
	"net/http"
	"strings"
)

//struct score data in memory
//1. struct
//2. setup
//3. conifgure
//4. init
//5. run phase
func (s *scoreServer) handleScoreSubmit() http.HandlerFunc {
	s.logger.Printf("Parsing goscore form here")
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "404: only posts live here", http.StatusNotFound)
			return
		}
		if err := r.ParseForm(); err != nil {
			log.Println(w, "ParseForm() err: %v", err)
			http.Error(w, "Error parsing form", http.StatusInternalServerError)
			return
		}
		value := strings.ToLower(r.FormValue("formValue"))
		s.logger.Printf("we got response for scoring, value: " + value)
		switch value {
		case "alfalfa":
			_, err := w.Write([]byte("This is their favorite food, not the solution"))
			if err != nil {
				//todo; error handling
				return
			}
		default:
			_, err := w.Write([]byte("Try harder"))
			if err != nil {
				//todo; error handling
				return
			}
		}
	}
}

func (s *scoreServer) handleGetNugget1() http.HandlerFunc {
	s.logger.Printf("Providing first nuget")
}