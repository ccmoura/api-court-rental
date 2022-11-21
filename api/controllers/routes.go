package controllers

import "api-court-rental/api/middlewares"

func (s *Server) initializeRoutes() {
	// owners routes
	s.Router.HandleFunc("/owners", middlewares.SetMiddlewareJSON(s.CreateOwner)).Methods("POST")
	s.Router.HandleFunc("/owners/{id}", middlewares.SetMiddlewareAuthentication(s.DeleteOwner)).Methods("DELETE")
	s.Router.HandleFunc("/owners/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.UpdateOwner))).Methods("PUT")
	s.Router.HandleFunc("/owners/{id}", middlewares.SetMiddlewareAuthentication(s.getOwnerById)).Methods("GET")

	// login route
	s.Router.HandleFunc("/login", middlewares.SetMiddlewareJSON(s.Login)).Methods("POST")
}
