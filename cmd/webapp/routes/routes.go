package routes

import (
	"net/http"

	"github.com/MakMoinee/go-mith/pkg/goserve"
)

type routeHandler struct {
}

func Set(httpService *goserve.Service) {
	handler := routeHandler{}

	httpService.Router.Get("/", handler.Home)
}

func (h *routeHandler) Home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}
