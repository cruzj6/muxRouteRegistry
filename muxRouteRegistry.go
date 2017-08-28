package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Handler func(w http.ResponseWriter, r *http.Request)

type RegisteredRoute struct {
	Method string
	Path string
	HandleRoute Handler
	SubRoutes []RegisteredRoute
}

var registeredRoutes = []RegisteredRoute{}

func SetupRoutes(muxRouter *mux.Router) *mux.Router{
	r := muxRouter

	for _, route := range registeredRoutes {

		s := r.PathPrefix(route.Path).Subrouter()
		s.HandleFunc("/", route.HandleRoute).Methods(route.Method)

		for _, subroute := range route.SubRoutes {
			s.HandleFunc(subroute.Path, subroute.HandleRoute).Methods(subroute.Method)
		}
	}

	return r
}

func RegisterRoute(route RegisteredRoute) {
	registeredRoutes = append(registeredRoutes, route)
}

func MakeRoute(path string, method string, handler Handler) RegisteredRoute{
	return RegisteredRoute{
		Method: method,
		Path: path,
		HandleRoute: handler,
	}
}

func (r *RegisteredRoute) AddSubRoute(route RegisteredRoute) {
	r.SubRoutes = append(r.SubRoutes, route)
}
