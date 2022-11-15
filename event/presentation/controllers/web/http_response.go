package controllers

import (
	"fmt"
	"net/http"
)

type HttpResponse struct {
	Code        int
	Message     string
	ContentType string
}

func NewHttpResponse(code int, message string, contentType string) HttpResponse {

	httpResponse := new(HttpResponse)
	httpResponse.Code = code
	httpResponse.Message = message
	httpResponse.ContentType = contentType
	return *httpResponse
}

func (hT HttpResponse) SendResponse(res http.ResponseWriter) {

	res.WriteHeader(hT.Code)
	res.Header().Set(
		"Content-Type",
		hT.ContentType,
	)

	fmt.Fprintln(res, hT.Message)
}

func (hT HttpResponse) SendError(res http.ResponseWriter) {

	res.WriteHeader(hT.Code)
	fmt.Fprintf(res, hT.Message)
	//TODO logger
}
