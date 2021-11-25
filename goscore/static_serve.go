package main

import (
	"net/http"
)

func (s *scoreServer) handleHome() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s.logger.Printf("serving Home")
		http.ServeFile(w,r, "resources/index.html")
	}
}

func (s *scoreServer) handleServeExercise2() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s.logger.Printf("serving exercise2")
		http.ServeFile(w,r, "resources/exercise2.html")
	}
}

func (s *scoreServer) handleServeExercise3() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s.logger.Printf("serving exercise3")
		http.ServeFile(w,r, "resources/exercise3.html")
	}
}