package client

type GeoCodingClient interface {
	FindAddress(address string) (*Response, error)
}

type Response struct {
	Results []Result `json:"results"`
}

type Result struct {
	Geometry Geometry `json:"geometry"`
}

type Geometry struct {
	Location Location `json:"location"`
}

type Location struct {
	Latitude  float64 `json:"lat"`
	Longitude float64 `json:"lng"`
}
