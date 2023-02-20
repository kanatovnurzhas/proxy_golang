package main

import (
	"fmt"
	"github.com/kanatovnurzhas/proxy_golang/internal/cache"
	"github.com/kanatovnurzhas/proxy_golang/internal/delivery"
	server2 "github.com/kanatovnurzhas/proxy_golang/internal/server"
	"github.com/kanatovnurzhas/proxy_golang/internal/service"
	"log"
)

const port = ":8888"

func main() {
	cacheInit := cache.NewCacheInit()
	proxyService := service.NewProxyServiceInit(cacheInit)
	handler := delivery.NewHandlerInit(proxyService)

	server := new(server2.Server)

	fmt.Printf("Starting server at port %s\nhttp://localhost%s/\n", port, port)
	if err := server.Run(port, handler.RoutesRegister()); err != nil {
		log.Fatal(err)
	}

}
