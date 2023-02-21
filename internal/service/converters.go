package service

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
	jsoniter "github.com/json-iterator/go"
	"github.com/kanatovnurzhas/proxy_golang/internal/models"
)

func convertKeyForCaching(req models.RequestProxy) (models.KeyRequest, error) {
	var keyRequest models.KeyRequest
	keyRequest.Method = req.Method
	keyRequest.Url = req.Url
	headersBytes, err := jsoniter.Marshal(req.Headers)
	if err != nil {
		return models.KeyRequest{}, err
	}
	keyRequest.Headers = string(headersBytes)
	fmt.Println(keyRequest.Headers)
	fmt.Println("Convert key:", keyRequest)
	return keyRequest, nil
}

func makeResponseProxy(resp *http.Response) models.ResponseProxy {
	var proxyResponse models.ResponseProxy
	proxyResponse.Headers = make(map[string][]string)
	proxyResponse.Id = int(uuid.New().ID())
	proxyResponse.Status = resp.StatusCode
	proxyResponse.Length = int(resp.ContentLength)
	for k, v := range resp.Header {
		proxyResponse.Headers[k] = v
	}
	return proxyResponse
}
