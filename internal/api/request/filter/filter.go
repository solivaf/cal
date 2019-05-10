package filter

import (
	"calindra/internal/api/request"
	"calindra/internal/api/response"
	"net/http"
)

func ValidateRequest(responseWriter http.ResponseWriter, req *http.Request) bool {
	if !request.IsValid(req) {
		resp := &response.ApiResponse{}
		bytesResponse, err := resp.CreateMissingRequiredParameters()
		if err != nil {
			responseWriter.WriteHeader(http.StatusInternalServerError)
			return false
		}
		responseWriter.WriteHeader(http.StatusBadRequest)
		responseWriter.Write(bytesResponse)
		return false
	}

	return true
}
