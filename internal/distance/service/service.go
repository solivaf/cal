package service

import (
	"calindra/internal/geocoding/client"
	"fmt"
	"log"
	"math"
	"sort"
)

type Service struct {
	client.GeoCodingClient
}

func CreateService(codingClient client.GeoCodingClient) *Service {
	return &Service{GeoCodingClient: codingClient}
}

func (s *Service) CalculateDistance(address, destination string) ([]float64, error) {
	distances := make([]float64, 0)
	addressResponse, err := s.GeoCodingClient.FindAddress(address)
	if err != nil {
		log.Println(fmt.Sprintf("Erro ao buscar dados do endereco %s", address))
		return nil, err
	}

	destinationResponse, err := s.GeoCodingClient.FindAddress(destination)
	if err != nil {
		log.Println(fmt.Sprintf("Erro ao buscar dados do endereco %s", address))
		return nil, err
	}

	log.Println(fmt.Sprintf("Iniciando calculo de distance entre %s e %s", address, destination))
	for _, addressResult := range addressResponse.Results {
		addressLatitude := addressResult.Geometry.Location.Latitude
		addressLongitude := addressResult.Geometry.Location.Longitude

		for _, destinationResult := range destinationResponse.Results {
			destinationLatitude := destinationResult.Geometry.Location.Latitude
			destinationLongitude := destinationResult.Geometry.Location.Longitude

			distance := calculateEuclideanDistance(addressLatitude, addressLongitude, destinationLatitude, destinationLongitude)

			log.Println(fmt.Sprintf("distancia entre latitudeX %b longitudeX %b e latitudeY %b longitudeY %b",
				addressLatitude, addressLongitude, destinationLatitude, destinationLongitude))
			distances = append(distances, distance)
		}
	}

	sort.Float64s(distances)

	return distances, nil
}

func calculateEuclideanDistance(latitudeX, longitudeX, latitudeY, longitudeY float64) float64 {
	distance := math.Sqrt(math.Pow(latitudeX-longitudeX, 2) + math.Pow(latitudeY-longitudeY, 2))

	log.Println(fmt.Sprintf("A distancia entre os dois pontos eh %b", distance))

	return distance
}

func (s *Service) toRadians(value float64) float64 {
	return value * math.Pi / 180.0
}
