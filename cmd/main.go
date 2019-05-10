package main

import (
	"calindra/internal/api/route"
	distanceService "calindra/internal/distance"
	"calindra/internal/distance/handler"
	"calindra/internal/geocoding/client/google"
	"net/http"
	"os"
)

func main() {
	client := google.CreateClient(3000, os.Getenv("GEOCODING_API_KEY"))
	service := distanceService.CreateService(client)
	distanceHandler := handler.CreateHandler(service)

	http.HandleFunc(route.Distance, distanceHandler.CalculateDistance)
	http.ListenAndServe(":8080", nil)
}
