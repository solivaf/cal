package response

import (
	"calindra/internal/api/request/params"
	"encoding/json"
)

type ApiResponse struct {
	Message string `json:"message"`
}

func (response ApiResponse) CreateMissingRequiredParameters() ([]byte, error) {
	response.Message = "Missing request parameters: " + params.Address + " or " + params.Destination
	return json.Marshal(&response)
}
