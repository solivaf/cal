package service_test

import (
	"calindra/internal/distance/service"
	"calindra/internal/geocoding/client"
	"calindra/internal/geocoding/client/mocks"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestService_CalculateDistanceConnectionTimeout(t *testing.T) {
	ctrl := gomock.NewController(t)
	geocodingClientMock := mock_client.NewMockGeoCodingClient(ctrl)
	geocodingClientMock.EXPECT().FindAddress("Rua General Serveriano, 205, Botafogo, Rio de Janeiro, 22290040").
		Return(nil, errors.New("connection i/o timeout"))

	service := service.CreateService(geocodingClientMock)
	_, err := service.CalculateDistance("Rua General Serveriano, 205, Botafogo, Rio de Janeiro, 22290040", "Rua da Auvernia, 286, Taúa, Ilha do Governador, Rio de Janeiro, 21920170")

	assert.Error(t, err)
	assert.Equal(t, "connection i/o timeout", err.Error())
}

func TestService_CalculateDistance(t *testing.T) {
	ctrl := gomock.NewController(t)
	geocodingClientMock := mock_client.NewMockGeoCodingClient(ctrl)

	var galSeverianoResponse client.Response
	if err := json.Unmarshal([]byte(`{"results":[{"geometry":{"location":{"lat":-22.798986,"lng":-43.193803}}}]}`), &galSeverianoResponse); err != nil {
		fmt.Println(err.Error())
	}
	geocodingClientMock.EXPECT().FindAddress("Rua General Serveriano, 205, Botafogo, Rio de Janeiro, 22290040").
		Return(&galSeverianoResponse, nil)

	var auverniaResponse client.Response
	if err := json.Unmarshal([]byte(`{"results":[{"geometry":{"location":{"lat":-22.9542736,"lng":-43.180108}}}]}`), &auverniaResponse); err != nil {
		fmt.Println(err.Error())
	}
	geocodingClientMock.EXPECT().FindAddress("Rua da Auvernia, 286, Taúa, Ilha do Governador, Rio de Janeiro, 21920170").
		Return(&auverniaResponse, nil)

	service := service.CreateService(geocodingClientMock)
	distances, err := service.CalculateDistance("Rua General Serveriano, 205, Botafogo, Rio de Janeiro, 22290040", "Rua da Auvernia, 286, Taúa, Ilha do Governador, Rio de Janeiro, 21920170")

	assert.NoError(t, err)
	assert.Equal(t, 1, len(distances))

}
