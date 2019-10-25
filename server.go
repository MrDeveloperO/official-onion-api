package main

import "net/http"

//Server is a wrapper around http.Server
type Server struct {
	addr string
	mux  *http.ServeMux
}

//NewServer creates new server
func NewServer(addr string) *Server {
	return &Server{
		addr: addr,
		mux:  http.NewServeMux(),
	}
}

//Register routes for handlers
func (s *Server) Register(ictrl IController) {
	s.mux.HandleFunc("/", ictrl.Links)
}

//ListenAndServe starts the server
func (s *Server) ListenAndServe() error {
	return http.ListenAndServe(s.addr, s.mux)
}
