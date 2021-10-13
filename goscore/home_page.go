package main

import (
	"net/http"
)

func (s scoreServer) handleHome() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w,r, "resources/index.html")
	}
}
