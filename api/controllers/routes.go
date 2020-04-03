package controllers

import "github.com/zhuangalbert/app-go/api/middlewares"

func (s *Server) initializeRoutes() {
	s.Router.HandleFunc("/", middlewares.SetMiddlewareJSON(s.Home)).Methods("GET")

	s.Router.HandleFunc("/api/login", middlewares.SetMiddlewareJSON(s.Login)).Methods("POST")

	s.Router.HandleFunc("/api/users", middlewares.SetMiddlewareJSON(s.GetUsers)).Methods("GET")
	s.Router.HandleFunc("/api/users/profile", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.GetProfile))).Methods("GET")
	s.Router.HandleFunc("/api/users", middlewares.SetMiddlewareJSON(s.CreateUser)).Methods("POST")

	s.Router.HandleFunc("/api/commodity", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.Getcommodity))).Methods("GET")
	s.Router.HandleFunc("/api/area", middlewares.SetMiddlewareJSON(s.GetArea)).Methods("GET")
	s.Router.HandleFunc("/api/size", middlewares.SetMiddlewareJSON(s.GetSize)).Methods("GET")
}
