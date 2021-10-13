package main

func (s scoreServer) Routes(){
	s.router.HandleFunc("/", s.addServerHeader(s.handleHome()))
	s.router.HandleFunc("/goscore", s.addServerHeader(s.handleScoreSubmit()))
}