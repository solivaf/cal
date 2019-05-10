package distance_test

import (
	"bytes"
	"calindra/internal/api/handler/distance"
	"calindra/internal/api/request/params"
	"calindra/internal/api/response"
	"calindra/internal/api/route"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCalculateMissingAddressParameter(t *testing.T) {
	resp := &response.ApiResponse{}
	expectedBody, _ := resp.CreateMissingRequiredParameters()

	request := httptest.NewRequest(http.MethodGet, route.Distance, nil)
	query := request.URL.Query()
	query.Add(params.Destination, "Rua da Auvernia, 286, Taúa, Ilha do Governador, Rio de Janeiro, 21920170")
	request.URL.RawQuery = query.Encode()

	recorder := httptest.NewRecorder()

	distance.CalculateDistance(recorder, request)

	assert.Equal(t, http.StatusBadRequest, recorder.Code)
	assert.NotEmpty(t, recorder.Body)
	assert.Equal(t, bytes.NewBuffer(expectedBody), recorder.Body)

}

func TestCalculateMissingDestinationParameter(t *testing.T) {
	resp := &response.ApiResponse{}
	expectedBody, _ := resp.CreateMissingRequiredParameters()

	request := httptest.NewRequest(http.MethodGet, route.Distance, nil)
	query := request.URL.Query()
	query.Add(params.Address, "Rua General Serveriano, 205, Botafogo, Rio de Janeiro, 22290040")
	request.URL.RawQuery = query.Encode()

	recorder := httptest.NewRecorder()

	distance.CalculateDistance(recorder, request)

	assert.Equal(t, http.StatusBadRequest, recorder.Code)
	assert.NotEmpty(t, recorder.Body)
	assert.Equal(t, bytes.NewBuffer(expectedBody), recorder.Body)

}

func TestCalculateDistanceMissingParameters(t *testing.T) {
	resp := &response.ApiResponse{}
	expectedBody, _ := resp.CreateMissingRequiredParameters()

	request := httptest.NewRequest(http.MethodGet, route.Distance, nil)
	recorder := httptest.NewRecorder()

	distance.CalculateDistance(recorder, request)

	assert.Equal(t, http.StatusBadRequest, recorder.Code)
	assert.NotEmpty(t, recorder.Body)
	assert.Equal(t, bytes.NewBuffer(expectedBody), recorder.Body)
}

func TestCalculateDistance(t *testing.T) {
	resp := &response.ApiResponse{}
	expectedBody, _ := resp.CreateMissingRequiredParameters()

	request := httptest.NewRequest(http.MethodGet, route.Distance, nil)
	query := request.URL.Query()
	query.Add(params.Address, "Rua General Serveriano, 205, Botafogo, Rio de Janeiro, 22290040")
	query.Add(params.Destination, "Rua da Auvernia, 286, Taúa, Ilha do Governador, Rio de Janeiro, 21920170")
	request.URL.RawQuery = query.Encode()

	recorder := httptest.NewRecorder()

	distance.CalculateDistance(recorder, request)

	assert.Equal(t, http.StatusBadRequest, recorder.Code)
	assert.NotEmpty(t, recorder.Body)
	assert.Equal(t, bytes.NewBuffer(expectedBody), recorder.Body)
}
