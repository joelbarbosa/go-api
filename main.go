package main

import (
	"fmt"
	"go-api/config"
	credit_routes "go-api/src/credit/router"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	log.Println("starting credit api")

	appPort := config.Get().AppPort

	routers := credit_routes.NewHttpRoute(
		credit_routes.RouterStruct{
			mux.NewRouter(),
		},
	)
	routers.GetRoute()
	http.Handle("/", routers)

	http.ListenAndServe(fmt.Sprintf(":%s", appPort), routers)
}
