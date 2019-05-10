package google

import (
	"calindra/internal/geocoding/client"
	"net/http"
	"time"
)

type geoCodingClient struct {
	ApiKey string
	*http.Client
}

func CreateClient(timeoutInMilli int, apiKey string) client.GeoCodingClient {
	httpClient := &http.Client{Timeout: time.Duration(time.Duration(timeoutInMilli) * time.Millisecond)}
	return &geoCodingClient{Client: httpClient, ApiKey: apiKey}
}

func (client *geoCodingClient) FindAddress(address string) (*client.Response, error) {

	return nil, nil
}
