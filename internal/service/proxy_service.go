package service

import (
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/kanatovnurzhas/proxy_golang/internal/cache"
	"github.com/kanatovnurzhas/proxy_golang/internal/models"
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
	cacheKey, err := convertKeyForCaching(request) // converting request for save in cache
	if err != nil {
		return models.ResponseProxy{}, fmt.Errorf("proxy service,error convert key for caching: %s", err)
	}
	response, ok := p.Get(cacheKey) // check request in the cache, if request in cache return response
	if ok {
		log.Println("response from cache")
		return response, nil
	}
	newReq, err := http.NewRequest(request.Method, request.Url, nil) // if request is new, create new request
	if err != nil {
		return models.ResponseProxy{}, fmt.Errorf("proxy service,error new request: %s", err)
	}
	for v, k := range request.Headers {
		newReq.Header.Set(k, v)
	}
	client := http.Client{}
	res, err := client.Do(newReq) // sends request and get response
	if err != nil {
		return models.ResponseProxy{}, fmt.Errorf("proxy service, error Do new request: %s", err)
	}
	p.Set(cacheKey, makeResponseProxy(res)) // save in cache request and response
	return makeResponseProxy(res), nil
}

func (p *proxyService) Set(key models.KeyRequest, resp models.ResponseProxy) {
	// если какая то бизнес логика будет можно будет здесь добавить, а так она по умолчанию просто записывает
	p.cacheRequest.SetRequest(key, resp)
}

func (p *proxyService) Get(key models.KeyRequest) (models.ResponseProxy, bool) {
	return p.cacheRequest.GetRequest(key)
}
