package bootstrap

import (
	"github.com/gorilla/mux"
	"goblog/pkg/route"
	"goblog/routes"
)

var Router *mux.Router

func SetupRoute()  *mux.Router{
	Router = mux.NewRouter()
	routes.RegisterWebRoutes(Router)
	route.SetRoute(Router)
	return Router
}
