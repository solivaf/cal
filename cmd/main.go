package cmd

import (
	"calindra/internal/api/handler/distance"
	"calindra/internal/api/route"
	"net/http"
)

func main() {
	http.HandleFunc(route.Distance, distance.CalculateDistance)
	http.ListenAndServe(":8080", nil)
}
