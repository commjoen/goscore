package main

func (s scoreServer) Routes(){
	s.router.HandleFunc("/", s.addServerHeader(s.handleHome()))
}