package controllers

import "github.com/garcialuis/ActivityCollector/api/middlewares"

// InitializeRoutes initializes request routes to server
func (s *Server) InitializeRoutes() {
	// Home
	s.Router.HandleFunc("/info", middlewares.SetMiddlewareJSON(s.Info)).Methods("GET")

	s.Router.HandleFunc("/activity/{origin_id}/{activity_id}", middlewares.SetMiddlewareJSON(s.GetActivity)).Methods("GET")
	s.Router.HandleFunc("/activities/{origin_id}", middlewares.SetMiddlewareJSON(s.GetActivities)).Methods("GET")
	s.Router.HandleFunc("/activities/{origin_id}/{starttime}/{endtime}", middlewares.SetMiddlewareJSON(s.GetActivitiesInRange)).Methods("GET")
	s.Router.HandleFunc("/testactivity/{origin_id}", middlewares.SetMiddlewareJSON(s.GetActivitySample)).Methods("GET")
	//TODO: s.Router.HandleFunc("/activity/{origin_id}", middlewares.SetMiddlewareJSON(s.DeleteActivity)).Methods("DELETE")
}
