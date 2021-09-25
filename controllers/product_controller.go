package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/leeliwei930/go_commerce/ent"
	"github.com/leeliwei930/go_commerce/repository"
	"github.com/leeliwei930/go_commerce/services"
)

func FetchProducts(response http.ResponseWriter, request *http.Request) {
	productRepo := repository.ProductRepository{
		Client: services.GetDefaultDBClient(),
	}

	params := request.URL.Query()

	currentPage, pageErr := strconv.Atoi(params.Get("page"))
	if pageErr != nil {
		currentPage = 1
	}

	products, productsError := productRepo.PaginateProduct(request.Context(), 3, currentPage)

	if productsError != nil {
		response.WriteHeader(http.StatusInternalServerError)
		productErrorResponse, _ := Response().Error(http.StatusInternalServerError, productsError)
		response.Write(productErrorResponse)
		return
	}

	response.WriteHeader(http.StatusOK)

	productsResponse, _ := Response().JSON(http.StatusOK, products)
	response.Write(productsResponse)

}

func CreateProduct(response http.ResponseWriter, request *http.Request) {
	productRepo := repository.ProductRepository{
		Client: services.GetDefaultDBClient(),
	}

	product := &ent.Product{}
	productErr := json.NewDecoder(request.Body).Decode(product)

	if productErr != nil {
		productResponse, _ := Response().Error(http.StatusUnprocessableEntity, productErr)
		response.WriteHeader(http.StatusUnprocessableEntity)
		response.Write(productResponse)
		return
	}

	product, createErr := productRepo.CreateProduct(request.Context(), product)
	if createErr != nil {
		productCreateErrResponse, _ := Response().Error(http.StatusInternalServerError, createErr)
		response.WriteHeader(http.StatusInternalServerError)
		response.Write(productCreateErrResponse)
		return
	}

	productResponse, _ := Response().JSON(http.StatusAccepted, product)

	response.Write(productResponse)

}

func FindProduct(response http.ResponseWriter, request *http.Request) {
	productRepo := repository.ProductRepository{
		Client: services.GetDefaultDBClient(),
	}

	path := mux.Vars(request)

	productID, productIDErr := uuid.Parse(path["id"])

	if productIDErr != nil {
		productIDErrResponse, _ := Response().Error(http.StatusBadRequest, productIDErr)
		response.Write(productIDErrResponse)
		return
	}

	product, productError := productRepo.FindProduct(request.Context(), productID)

	if productError != nil {
		productErrResponse, _ := Response().Error(http.StatusBadRequest, productError)
		response.WriteHeader(http.StatusNotFound)
		response.Write(productErrResponse)
		return
	}

	productResponse, _ := Response().JSON(http.StatusOK, product)
	response.WriteHeader(http.StatusOK)
	response.Write(productResponse)

}
