package service

import (
	"github.com/kanatovnurzhas/proxy_golang/internal/models"
	"net/http"
)

func convertKeyForCaching(req models.RequestProxy) (models.KeyRequest, error) {
	return models.KeyRequest{}, nil
}

func makeResponseProxy(resp *http.Response) models.ResponseProxy {
	return models.ResponseProxy{}
}
