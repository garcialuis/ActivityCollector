package controllers

import "github.com/garcialuis/ActivityCollector/api/middlewares"

func (s *Server) InitializeRoutes() {
	// Home
	s.Router.HandleFunc("/info", middlewares.SetMiddlewareJSON(s.Info)).Methods("GET")

	// FoodItems
	s.Router.HandleFunc("/activity/{origin_id}", middlewares.SetMiddlewareJSON(s.GetActivity)).Methods("GET")
	s.Router.HandleFunc("/testactivity/{origin_id}", middlewares.SetMiddlewareJSON(s.GetActivitySample)).Methods("GET")
	// s.Router.HandleFunc("/fooditem", middlewares.SetMiddlewareJSON(s.CreateFoodItem)).Methods("POST")
	// s.Router.HandleFunc("/fooditem", middlewares.SetMiddlewareJSON(s.GetAllFoodItems)).Methods("GET")
	// s.Router.HandleFunc("/fooditem/{foodName}", middlewares.SetMiddlewareJSON(s.DeleteFoodItemByName)).Methods("DELETE")
}
