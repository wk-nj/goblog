package bootstrap

import (
	"github.com/gorilla/mux"
	"goblog/routes"
)

var Router *mux.Router

func SetupRoute()  *mux.Router{
	Router = mux.NewRouter()
	routes.RegisterWebRoutes(Router)
	return Router
}
