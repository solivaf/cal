package cmd

import (
	"calindra/internal/api/handler"
	"calindra/internal/api/route"
	"net/http"
)

func main() {
	http.HandleFunc(route.Distance, handler.Distance)
	http.ListenAndServe(":8080", nil)
}
