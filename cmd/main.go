package main

import (
	"calindra/internal/api/route"
	"calindra/internal/distance/handler"
	distanceService "calindra/internal/distance/service"
	"calindra/internal/geocoding/client/google"
	"fmt"
	"net/http"
	"os"
)

func main() {
	client := google.CreateClient(3000, os.Getenv("GEOCODING_API_KEY"), os.Getenv("GEOCODING_API_URL"))
	service := distanceService.CreateService(client)
	distanceHandler := handler.CreateHandler(service)

	http.HandleFunc(route.Distance, distanceHandler.CalculateDistance)
	fmt.Println(http.ListenAndServe(":"+os.Getenv("SERVER_PORT"), nil))
}
