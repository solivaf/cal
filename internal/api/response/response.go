package response

import (
	"calindra/internal/api/request/params"
	"encoding/json"
)

type ApiResponse struct {
	Message      string        `json:"message,omitempty"`
	DistanceInfo *DistanceInfo `json:"distance_info,omitempty"`
}

type DistanceInfo struct {
	Distances []float64 `json:"distances,omitempty"`
	Closer    float64   `json:"closer,omitempty"`
}

func (response ApiResponse) CreateMissingRequiredParameters() ([]byte, error) {
	response.Message = "Missing request parameters: " + params.Address + " or " + params.Destination
	return json.Marshal(&response)
}
