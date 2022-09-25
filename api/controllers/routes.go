
package controllers

import "projects/api-court-rental/api/middlewares"

func (s *Server) initializeRoutes() {
	// owners routes
	s.Router.HandleFunc("/owners", middlewares.SetMiddlewareJSON(s.CreateOwner)).Methods("POST")
	s.Router.HandleFunc("/owners/{id}", s.DeleteOwner).Methods("DELETE")
}
