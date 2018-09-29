package main

import (
	"log"
	"net/http"

	"github.com/avtal.chalmers.it/backend/config"
	"github.com/avtal.chalmers.it/backend/domains/auth"
	"github.com/avtal.chalmers.it/backend/domains/util"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

var Secrets struct {
	BetaClientSecret string `json:"client_secret"`
}

func main() {
	config.LoadSecrets()
	rootRouter := CreateRootRouter()
	rootRouter.PathPrefix("/static/").Handler(staticHandler())

	corsHandler := GetCorsHandler()
	log.Fatal(http.ListenAndServe(":3000", corsHandler(rootRouter)))
}

/* The root router maps http requests to domain routes. */
func CreateRootRouter() *mux.Router {
	router := mux.NewRouter()
	for _, route := range getDomainRoutes() {
		handler := route.HandlerFunc
		router.Methods(route.Method).Path(route.Pattern).Name(route.Name).Handler(handler)
	}
	return router
}

func GetCorsHandler() func(http.Handler) http.Handler {
	allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	allowedMethods := handlers.AllowedMethods([]string{"POST", "GET", "PUT", "DELETE"})
	return handlers.CORS(allowedOrigins, allowedMethods)
}

func getDomainRoutes() []util.Route {
	var routes []util.Route
	routes = append(routes, auth.GetRoutes()...)
	return routes
}

func staticHandler() http.Handler {
	return http.StripPrefix("/static/", http.FileServer(http.Dir("./static/")))
}
