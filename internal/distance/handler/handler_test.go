package handler_test

import (
	"bytes"
	"calindra/internal/api/request/params"
	"calindra/internal/api/response"
	"calindra/internal/api/route"
	"calindra/internal/distance/handler"
	distanceService "calindra/internal/distance/service"
	"calindra/internal/geocoding/client"
	"calindra/internal/geocoding/client/google"
	"calindra/internal/geocoding/client/mocks"
	"encoding/json"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

const (
	AuverniaResponse     = `{"results":["geometry":{"location":{"lat":-22.798986,"lng":-43.193803}},"status":"OK"}`
	GalSeverianoResponse = `{"results":["geometry":{"location":{"lat":-22.9542736,"lng":-43.180108}}`
)

func TestCalculateMissingAddressParameter(t *testing.T) {
	client := google.CreateClient(3000, "foo")
	service := distanceService.CreateService(client)
	distanceHandler := handler.CreateHandler(service)

	resp := &response.ApiResponse{}
	expectedBody, _ := resp.CreateMissingRequiredParameters()

	request := httptest.NewRequest(http.MethodGet, route.Distance, nil)
	query := request.URL.Query()
	query.Add(params.Destination, "Rua da Auvernia, 286, Taúa, Ilha do Governador, Rio de Janeiro, 21920170")
	request.URL.RawQuery = query.Encode()

	recorder := httptest.NewRecorder()

	distanceHandler.CalculateDistance(recorder, request)

	assert.Equal(t, http.StatusBadRequest, recorder.Code)
	assert.NotEmpty(t, recorder.Body)
	assert.Equal(t, bytes.NewBuffer(expectedBody), recorder.Body)

}

func TestCalculateMissingDestinationParameter(t *testing.T) {
	client := google.CreateClient(3000, "foo")
	service := distanceService.CreateService(client)
	distanceHandler := handler.CreateHandler(service)

	resp := &response.ApiResponse{}
	expectedBody, _ := resp.CreateMissingRequiredParameters()

	request := httptest.NewRequest(http.MethodGet, route.Distance, nil)
	query := request.URL.Query()
	query.Add(params.Address, "Rua General Serveriano, 205, Botafogo, Rio de Janeiro, 22290040")
	request.URL.RawQuery = query.Encode()

	recorder := httptest.NewRecorder()

	distanceHandler.CalculateDistance(recorder, request)

	assert.Equal(t, http.StatusBadRequest, recorder.Code)
	assert.NotEmpty(t, recorder.Body)
	assert.Equal(t, bytes.NewBuffer(expectedBody), recorder.Body)

}

func TestCalculateDistanceMissingParameters(t *testing.T) {
	client := google.CreateClient(3000, "foo")
	service := distanceService.CreateService(client)
	distanceHandler := handler.CreateHandler(service)

	resp := &response.ApiResponse{}
	expectedBody, _ := resp.CreateMissingRequiredParameters()

	request := httptest.NewRequest(http.MethodGet, route.Distance, nil)
	recorder := httptest.NewRecorder()

	distanceHandler.CalculateDistance(recorder, request)

	assert.Equal(t, http.StatusBadRequest, recorder.Code)
	assert.NotEmpty(t, recorder.Body)
	assert.Equal(t, bytes.NewBuffer(expectedBody), recorder.Body)
}

func TestCalculateDistance(t *testing.T) {
	ctrl := gomock.NewController(t)
	geocodingClientMock := mock_client.NewMockGeoCodingClient(ctrl)

	var galSeverianoResponse client.Response
	json.Unmarshal([]byte(GalSeverianoResponse), &galSeverianoResponse)
	geocodingClientMock.EXPECT().FindAddress("Rua General Serveriano, 205, Botafogo, Rio de Janeiro, 22290040").
		Return(&galSeverianoResponse, nil)

	var auverniaResponse client.Response
	json.Unmarshal([]byte(AuverniaResponse), &auverniaResponse)
	geocodingClientMock.EXPECT().FindAddress("Rua da Auvernia, 286, Taúa, Ilha do Governador, Rio de Janeiro, 21920170").
		Return(&auverniaResponse, nil)

	service := distanceService.CreateService(geocodingClientMock)
	distanceHandler := handler.CreateHandler(service)

	request := httptest.NewRequest(http.MethodGet, route.Distance, nil)
	query := request.URL.Query()
	query.Add(params.Address, "Rua General Serveriano, 205, Botafogo, Rio de Janeiro, 22290040")
	query.Add(params.Destination, "Rua da Auvernia, 286, Taúa, Ilha do Governador, Rio de Janeiro, 21920170")
	request.URL.RawQuery = query.Encode()

	recorder := httptest.NewRecorder()

	distanceHandler.CalculateDistance(recorder, request)

	assert.Equal(t, http.StatusOK, recorder.Code)
}

func TestHandler_CalculateDistanceServiceError(t *testing.T) {
	ctrl := gomock.NewController(t)
	geocodingClientMock := mock_client.NewMockGeoCodingClient(ctrl)

	geocodingClientMock.EXPECT().FindAddress("Rua General Serveriano, 205, Botafogo, Rio de Janeiro, 22290040").
		Return(nil, errors.New("connection i/o timeout"))

	service := distanceService.CreateService(geocodingClientMock)
	distanceHandler := handler.CreateHandler(service)

	request := httptest.NewRequest(http.MethodGet, route.Distance, nil)
	query := request.URL.Query()
	query.Add(params.Address, "Rua General Serveriano, 205, Botafogo, Rio de Janeiro, 22290040")
	query.Add(params.Destination, "Rua da Auvernia, 286, Taúa, Ilha do Governador, Rio de Janeiro, 21920170")
	request.URL.RawQuery = query.Encode()

	recorder := httptest.NewRecorder()

	distanceHandler.CalculateDistance(recorder, request)

	assert.Equal(t, http.StatusInternalServerError, recorder.Code)
}
