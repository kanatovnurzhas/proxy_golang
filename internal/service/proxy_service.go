package service

import (
	"fmt"
	"github.com/kanatovnurzhas/proxy_golang/internal/cache"
	"github.com/kanatovnurzhas/proxy_golang/internal/models"
	"log"
	"net/http"
	"net/url"
)

type IProxyService interface {
	Set(key models.KeyRequest, resp models.ResponseProxy)
	Get(key models.KeyRequest) (models.ResponseProxy, bool)
	ProxyRequest(request models.RequestProxy) (models.ResponseProxy, error)
}

type proxyService struct {
	cacheRequest cache.ICache
}

func NewProxyServiceInit(cache cache.ICache) IProxyService {
	return &proxyService{cacheRequest: cache}
}

func (p *proxyService) ProxyRequest(request models.RequestProxy) (models.ResponseProxy, error) {
	if request.Method != "GET" {
		return models.ResponseProxy{}, fmt.Errorf("proxy service,use only GET method")
	}
	_, err := url.Parse(request.Url)
	if err != nil {
		return models.ResponseProxy{}, fmt.Errorf("proxy service,error parse url: %s", err)
	}
	cacheKey, err := convertKeyForCaching(request)
	if err != nil {
		return models.ResponseProxy{}, fmt.Errorf("proxy service,error convert key for caching: %s", err)
	}
	response, ok := p.Get(cacheKey)
	if ok {
		log.Println("response from cache")
		return response, nil
	}

	newReq, err := http.NewRequest(request.Method, request.Url, nil)
	if err != nil {
		return models.ResponseProxy{}, fmt.Errorf("proxy service,error new request: %s", err)
	}
	return models.ResponseProxy{}, nil
}

func (p *proxyService) Set(key models.KeyRequest, resp models.ResponseProxy) {
	//если какая то бизнес логика будет можно будет здесь добавить, а так она по умолчанию просто записывает
	p.Set(key, resp)
}

func (p *proxyService) Get(key models.KeyRequest) (models.ResponseProxy, bool) {
	return p.Get(key)
}
