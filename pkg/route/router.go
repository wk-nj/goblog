package route

import (
	"github.com/gorilla/mux"
	"goblog/pkg/config"
	"goblog/pkg/logger"
	"net/http"
)

var Route *mux.Router

func SetRoute(r *mux.Router)  {
	Route = r
}

// Name2URL 通过路由名称来获取 URL
func Name2URL(routeName string, pairs ...string) string {
	url, err := Route.Get(routeName).URL(pairs...)
	if err != nil {
		logger.LogError(err)
		return ""
	}

	return config.GetString("app.url") + url.String()
}

func GetRouteVariable(parameterName string, r *http.Request) string {
	vars := mux.Vars(r)
	return vars[parameterName]
}