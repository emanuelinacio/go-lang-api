package controllers

import (
	"errors"
	uc "events/event/use-cases"
	"net/http"
)

type SendEventController struct {
}

func NewSendEventController() SendEventController {
	sendEventController := new(SendEventController)
	return *sendEventController
}

func (sC SendEventController) ValidParameters(res http.ResponseWriter, req *http.Request) error {

	if req.URL.Query().Has("view") == false {
		return errors.New("Invalid parameters view is required")
	}

	if req.URL.Query().Has("source") == false {
		return errors.New("Invalid parameters source is required")
	}

	if req.URL.Query().Has("topic") == false {
		return errors.New("Invalid parameters topic is required")
	}

	if req.URL.Query().Has("referer") == false {
		return errors.New("Invalid parameters referer is required")
	}

	return nil
}

func (sC SendEventController) ValidRequest(res http.ResponseWriter, req *http.Request) (error, *HttpResponse) {
	httpResponse := new(HttpResponse)

	if req.Method == "POST" {
		httpResponse.Code = 401
		httpResponse.Message = "Invalid Method"
		return errors.New("Invalid Method"), httpResponse
	}

	err := sC.ValidParameters(res, req)
	if err != nil {
		httpResponse.Code = 401
		httpResponse.Message = err.Error()
		return errors.New(err.Error()), httpResponse
	}

	return nil, httpResponse
}

func (sC SendEventController) Handler(res http.ResponseWriter, req *http.Request) {

	err, httpResponse := sC.ValidRequest(res, req)

	if err != nil {
		httpResponse.SendError(res)
		return
	}

	localEvent, err := uc.CreateEvent(
		req.URL.Query().Get("view"),
		req.URL.Query().Get("source"),
		req.URL.Query().Get("topic"),
		req.URL.Query().Get("referer"),
	)

	if err != nil {
		httpResponse.Code = 401
		httpResponse.Message = err.Error()
		httpResponse.SendError(res)
		return
	}

	localEventJson, _ := localEvent.ToJson()
	httpResponse.Code = 200
	httpResponse.Message = "evento criado com sucesso" + localEventJson
	httpResponse.ContentType = "application/json"
	httpResponse.SendResponse(res)
}

func (sC SendEventController) SetRouter() {
	http.HandleFunc("/sendevent", sC.Handler)
}
