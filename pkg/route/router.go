package route

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

// Name2URL 通过路由名称来获取 URL
func Name2URL(routeName string, pairs ...string) string {
	var route *mux.Router
	fmt.Println(route.Get(routeName))
	url, err := route.Get(routeName).URL(pairs...)
	if err != nil {
		// checkError(err)
		return ""
	}

	return url.String()
}

func GetRouteVariable(parameterName string, r *http.Request) string {
	vars := mux.Vars(r)
	return vars[parameterName]
}