package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/leeliwei930/go_commerce/controllers"
)

func StartServer() {
	muxRouter := mux.NewRouter()

	muxRouter.HandleFunc("/product", controllers.FetchProducts).Methods(http.MethodGet).Queries()
	muxRouter.HandleFunc("/product", controllers.FetchProducts).Methods(http.MethodPost)
	muxRouter.HandleFunc("/product/{id}", controllers.FindProduct).Methods(http.MethodGet)
	muxRouter.Use(controllers.InjectJSONResponseHeader)
	server := &http.Server{
		Addr:    fmt.Sprintf("0.0.0.0:%s", os.Getenv("APP_PORT")),
		Handler: muxRouter,
	}
	fmt.Printf("Gorilla Mux running on http://%s \n", server.Addr)
	if listenError := server.ListenAndServe(); listenError != nil {
		log.Fatalln(listenError)
	}

}
