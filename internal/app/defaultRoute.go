package app

import (
	"fmt"
	"net/http"
)

type DefaultRoute struct{}

func defaultRoute(router *http.ServeMux) {
	handler := &DefaultRoute{}
	router.HandleFunc("GET /", handler.apiReady)
}

func (h *DefaultRoute) apiReady(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "API is ready. v1")
}
