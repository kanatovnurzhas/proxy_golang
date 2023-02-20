package cache

import (
	"github.com/kanatovnurzhas/proxy_golang/internal/models"
	"log"
	"sync"
)

type ICache interface {
	SetRequest(key models.KeyRequest, resp models.ResponseProxy)
	GetRequest(key models.KeyRequest) (models.ResponseProxy, bool)
}

type cacheRequests struct {
	cache map[models.KeyRequest]models.ResponseProxy
	mutex *sync.RWMutex
}

func NewCacheInit() ICache {
	return &cacheRequests{
		cache: make(map[models.KeyRequest]models.ResponseProxy),
		mutex: &sync.RWMutex{},
	}
}

func (c *cacheRequests) SetRequest(key models.KeyRequest, resp models.ResponseProxy) {
	c.mutex.RLock()
	c.cache[key] = resp
	defer c.mutex.RUnlock()
	log.Println("request and response saved in cache")
}

func (c *cacheRequests) GetRequest(key models.KeyRequest) (models.ResponseProxy, bool) {
	c.mutex.RLock()
	data, ok := c.cache[key]
	defer c.mutex.RUnlock()
	return data, ok
}
