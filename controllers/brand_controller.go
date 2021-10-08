package controllers

import (
	"net/http"

	"github.com/leeliwei930/go_commerce/inputs"
	"github.com/leeliwei930/go_commerce/repository"
	"github.com/leeliwei930/go_commerce/services"
)

func FetchBrand(response http.ResponseWriter, request *http.Request) {

}

func CreateBrand(responseWriter http.ResponseWriter, request *http.Request) {
	brandRepo := repository.BrandRepository{
		Client: services.GetDefaultDBClient(),
	}

	brandRequest := &inputs.BrandRequest{}
	brandErr := inputs.BindFromJSON(request.Body, brandRequest)

	if brandErr != nil {
		brandResponse, _ := Response().Error(http.StatusUnprocessableEntity, brandErr)
		responseWriter.WriteHeader(http.StatusUnprocessableEntity)
		responseWriter.Write(brandResponse)
		return
	}

	createdBrand, createErr := brandRepo.CreateBrand(request.Context(), brandRequest)
	if createErr != nil {
		brandCreateErrResponse, _ := Response().Error(http.StatusInternalServerError, createErr)
		responseWriter.WriteHeader(http.StatusInternalServerError)
		responseWriter.Write(brandCreateErrResponse)
		return
	}

	brandResponse, _ := Response().JSON(http.StatusAccepted, createdBrand, nil)

	responseWriter.Write(brandResponse)
}

func UpdateBrand() {

}

func DeleteBrand() {

}
