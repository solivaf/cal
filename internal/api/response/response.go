package response

import (
	"calindra/internal/api/request/params"
	"encoding/json"
)

type ApiResponse struct {
	Message      string       `json:"message"`
	DistanceInfo DistanceInfo `json:"distance_info"`
}

type DistanceInfo struct {
	Distances []float64 `json:"distances"`
	Closer    float64   `json:"closer"`
}

func (response ApiResponse) CreateMissingRequiredParameters() ([]byte, error) {
	response.Message = "Missing request parameters: " + params.Address + " or " + params.Destination
	return json.Marshal(&response)
}
