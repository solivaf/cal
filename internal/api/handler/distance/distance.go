package distance

import (
	_request "calindra/internal/api/request"
	"calindra/internal/api/response"
	"log"
	"net/http"
)

func CalculateDistance(responseWriter http.ResponseWriter, req *http.Request) {
	log.Println("Recebendo request em /distance ")
	if !_request.IsValid(req) {
		resp := &response.ApiResponse{}
		bytesResponse, err := resp.CreateMissingRequiredParameters()
		if err != nil {
			responseWriter.WriteHeader(http.StatusInternalServerError)
			return
		}
		responseWriter.WriteHeader(http.StatusBadRequest)
		responseWriter.Write(bytesResponse)
		return
	}
}
