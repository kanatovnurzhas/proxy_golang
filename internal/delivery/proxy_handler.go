package delivery

import (
	jsoniter "github.com/json-iterator/go"
	"github.com/kanatovnurzhas/proxy_golang/internal/models"
	"log"
	"net/http"
)

func (h *Handler) proxyReverse(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var proxyReq models.RequestProxy
		if err := jsoniter.NewDecoder(r.Body).Decode(&proxyReq); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Printf("error decode request:%s", err)
			return
		}
		response, err := h.proxyService.ProxyRequest(proxyReq)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Printf("error proxy request func: %s", err)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		if err := jsoniter.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Printf("error encode response: %s", err)
			return
		}
	} else {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		log.Println("error method not allowed")
		return
	}
}
