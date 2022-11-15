package presentation

import (
	"net/http"
)

type Controller interface {
	Handler(res http.ResponseWriter, req *http.Request)
	SetRouter()
	ValidRequest()
}
