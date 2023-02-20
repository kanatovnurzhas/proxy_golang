package delivery

import (
	"github.com/kanatovnurzhas/proxy_golang/internal/service"
	"net/http"
)

type Handler struct {
	proxyService service.IProxyService
}

func NewHandlerInit(proxy service.IProxyService) *Handler {
	return &Handler{proxyService: proxy}
}

func (h *Handler) RoutesRegister() *http.ServeMux {
	router := http.NewServeMux()
	router.HandleFunc("/", h.proxyReverse)
	return router
}
