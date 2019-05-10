package google

import (
	"calindra/internal/api/request/params"
	"calindra/internal/geocoding/client"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"
)

type geoCodingClient struct {
	GeoCodeApiUrl string
	ApiKey        string
	httpClient    *http.Client
}

func CreateClient(timeoutInMilli int, apiKey, geocodeApiUrl string) client.GeoCodingClient {
	httpClient := &http.Client{Timeout: time.Duration(time.Duration(timeoutInMilli) * time.Millisecond)}
	return &geoCodingClient{httpClient: httpClient, ApiKey: apiKey, GeoCodeApiUrl: geocodeApiUrl}
}

func (c *geoCodingClient) FindAddress(address string) (*client.Response, error) {
	encodedUrl := c.encodedUrl(address, c.ApiKey)
	resp, err := c.httpClient.Get(encodedUrl)
	if err != nil {
		log.Println(fmt.Sprintf("Erro ao executar request para url %s - error %s", encodedUrl, err.Error()))
		return nil, err
	}
	reader, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Erro ao realizar leitura do corpo da resposta")
		return nil, err
	}
	defer resp.Body.Close()

	var apiResponse client.Response
	err = json.Unmarshal(reader, &apiResponse)
	if err != nil {
		log.Println(fmt.Sprintf("Erro ao realizar unmarshal do corpo da resposta - error %s", err.Error()))
		return nil, err
	}
	return &apiResponse, nil
}

func (client *geoCodingClient) encodedUrl(address, apiKey string) string {
	encodedUrl, err := url.Parse(client.GeoCodeApiUrl)
	if err != nil {
		log.Fatal(fmt.Sprintf("Erro ao realizar criacao de url %s error %s", client.GeoCodeApiUrl, err.Error()))
	}

	query := encodedUrl.Query()
	query.Add(params.Address, address)
	query.Add(params.Key, apiKey)
	encodedUrl.RawQuery = query.Encode()
	return encodedUrl.String()
}
