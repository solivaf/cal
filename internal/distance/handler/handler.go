package handler

import (
	_request "calindra/internal/api/request"
	"calindra/internal/api/request/filter"
	"calindra/internal/api/request/params"
	"calindra/internal/api/response"
	"calindra/internal/distance"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Handler struct {
	DistanceService *distance.Service
}

func CreateHandler(service *distance.Service) *Handler {
	return &Handler{DistanceService: service}
}

func (h *Handler) CalculateDistance(responseWriter http.ResponseWriter, req *http.Request) {
	log.Println("Recebendo request " + req.URL.String())
	if isValid := filter.ValidateRequest(responseWriter, req); !isValid {
		log.Println("Request invalido. Faltando parametros obrigatorios")
		return
	}

	address := _request.GetQueryParamFromRequest(params.Address, *req)
	destination := _request.GetQueryParamFromRequest(params.Destination, *req)

	_, err := h.DistanceService.CalculateDistance(address, destination)
	if err != nil {
		responseWriter.WriteHeader(http.StatusInternalServerError)
		return
	}

	resp := &response.ApiResponse{DistanceInfo: response.DistanceInfo{}}
	responseBytes, err := json.Marshal(resp)
	if err != nil {
		log.Println(fmt.Sprintf("Erro ao realizar marshal - error %s", err.Error()))
		responseWriter.WriteHeader(http.StatusInternalServerError)
		return
	}

	responseWriter.WriteHeader(http.StatusOK)
	responseWriter.Write(responseBytes)

}
