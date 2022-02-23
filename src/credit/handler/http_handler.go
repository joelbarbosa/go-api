package handler

import (
	"encoding/json"
	"go-api/src/credit/entities"
	"go-api/src/credit/service"
	"io/ioutil"
	"log"
	"net/http"
)

type CreditHandler interface {
	CreditLine(w http.ResponseWriter, r *http.Request)
}

type creditHandler struct {
	CreditService service.CreditService
}

// NewHttpHandle create an instance handle
func NewHttpHandle(creditService service.CreditService) CreditHandler {
	return &creditHandler{
		CreditService: creditService,
	}
}

// CreditLine handle with request and response api
func (service *creditHandler) CreditLine(w http.ResponseWriter, r *http.Request) {
	log.Println("handle CreditLine...")
	reqBody, _ := ioutil.ReadAll(r.Body)
	var post entities.Credit
	json.Unmarshal(reqBody, &post)

	creditLine := service.CreditService.CreditLine(post)

	json.NewEncoder(w).Encode(creditLine)

}
