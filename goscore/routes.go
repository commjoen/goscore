package main

func (s scoreServer) Routes(){
	s.router.HandleFunc("/exercise1/left", s.addServerHeader(s.handleLeftExercise1()))
	s.router.HandleFunc("/goscore", s.addServerHeader(s.handleScoreSubmit()))
	s.router.HandleFunc("/", s.addServerHeader(s.handleHome()))

}