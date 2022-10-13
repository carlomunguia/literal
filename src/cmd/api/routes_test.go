package main

import (
	"net/http"
	"testing"

	"github.com/go-chi/chi/v5"
)

func Test_If_Routes_exist(t *testing.T) {
	testRoutes := testApp.routes()
	chiRoutes := testRoutes.(chi.Router)

	doesRouteExist(t, chiRoutes, "/users/login")
	doesRouteExist(t, chiRoutes, "/users/logout")
	doesRouteExist(t, chiRoutes, "/admin/users/get/{id}")
	doesRouteExist(t, chiRoutes, "/admin/users/save")
	doesRouteExist(t, chiRoutes, "/admin/users/delete")
}

func doesRouteExist(t *testing.T, routes chi.Router, route string) {
	found := false

	_ = chi.Walk(routes, func(method string, foundRoute string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		if route == foundRoute {
			found = true
		}

		return nil
	})

	if !found {
		t.Errorf("Route %s not found", route)
	}
}
