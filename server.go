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
	// product routes
	muxRouter.HandleFunc("/product", controllers.CreateProduct).Methods(http.MethodPost)
	muxRouter.HandleFunc("/product", controllers.FetchProducts).Methods(http.MethodGet).Queries()
	muxRouter.HandleFunc("/product/{id}", controllers.FindProduct).Methods(http.MethodGet)
	muxRouter.HandleFunc("/product/{id}", controllers.FindProduct).Methods(http.MethodPut)
	muxRouter.HandleFunc("/product/{id}/delete", controllers.DeleteProduct).Methods(http.MethodPut)
	muxRouter.HandleFunc("/product/{id}/restore", controllers.RestoreProduct).Methods(http.MethodPut)
	muxRouter.HandleFunc("/product/{id}/destroy", controllers.DestroyProduct).Methods(http.MethodDelete)
	muxRouter.Use(controllers.InjectJSONResponseHeader)

	muxRouter.HandleFunc("/brand", controllers.FetchBrand).Methods(http.MethodGet)

	server := &http.Server{
		Addr:    fmt.Sprintf("0.0.0.0:%s", os.Getenv("APP_PORT")),
		Handler: muxRouter,
	}
	fmt.Printf("Gorilla Mux running on http://%s \n", server.Addr)
	if listenError := server.ListenAndServe(); listenError != nil {
		log.Fatalln(listenError)
	}

}
