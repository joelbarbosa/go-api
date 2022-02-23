package router

import (
	"go-api/src/credit/handler"
	"go-api/src/credit/service"

	"github.com/gorilla/mux"
)

type RouterStruct struct {
	*mux.Router
}

// NewHttpRoute create an instance
func NewHttpRoute(structs RouterStruct) RouterStruct {
	return structs
}

// GetRoute load all methods from credit api
func (router *RouterStruct) GetRoute() {
	creditService := service.NewCreditService()
	creditHandlers := handler.NewHttpHandle(creditService)
	router.HandleFunc("/credit-line", creditHandlers.CreditLine).Methods("POST")
}
