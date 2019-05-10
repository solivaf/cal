package handler

import (
	_request "calindra/internal/api/request"
	"calindra/internal/api/request/filter"
	"calindra/internal/api/request/params"
	"calindra/internal/api/response"
	"calindra/internal/distance/service"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Handler struct {
	DistanceService *service.Service
}

func CreateHandler(service *service.Service) *Handler {
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

	distances, err := h.DistanceService.CalculateDistance(address, destination)
	if err != nil {
		responseWriter.WriteHeader(http.StatusInternalServerError)
		return
	}

	resp := &response.ApiResponse{DistanceInfo: &response.DistanceInfo{Closer: distances[0], Distances: distances}}
	responseBytes, err := json.Marshal(resp)
	if err != nil {
		log.Println(fmt.Sprintf("Erro ao realizar marshal - error %s", err.Error()))
		responseWriter.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.Println(fmt.Sprintf("Calculo realizado com sucesso, corpo de resposta %s", string(responseBytes)))

	responseWriter.Header().Add("Content-Type", "application/json")
	responseWriter.WriteHeader(http.StatusOK)
	responseWriter.Write(responseBytes)

}
