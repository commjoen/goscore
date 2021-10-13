package main

import (
	"net/http"
)

func (s scoreServer) handExercise1Left() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w,r, "resources/index.html")
	}
}

func (s scoreServer) handExercise1Right() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w,r, "resources/index.html")
	}
}
