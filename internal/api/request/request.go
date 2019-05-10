package request

import (
	"calindra/internal/api/request/params"
	"net/http"
)

const empty = ""

func IsValid(request *http.Request) bool {
	query := request.URL.Query()
	return query.Get(params.Address) != empty && query.Get(params.Destination) != empty
}
