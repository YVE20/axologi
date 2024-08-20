package controller

import (
	"axologi/constant"
	"axologi/helper"
	"axologi/service"
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// Ini untuk declare service
type ProvinceController struct {
	service service.ProvinceService
}

// Construct untuk execute controller ini di main.go
func NewProvinceController(s service.ProvinceService) *ProvinceController {
	return &ProvinceController{
		service: s,
	}
}

func (c ProvinceController) GetAll(writer http.ResponseWriter, request *http.Request) {
	responseData, err := c.service.GetAllProvinceService()

	// Set Content-Type header
	writer.Header().Set("Content-Type", "application/json")

	if err != nil {
		response := helper.GeneralWrapper{
			StatusCode: constant.STATUS_CODE_CONFLICT,
			Message:    constant.MESSAGE_CONFLICT,
			Data:       nil,
			TimeStamp:  time.Now(),
		}
		json.NewEncoder(writer).Encode(response)
		return
	}

	response := helper.GeneralWrapper{
		StatusCode: constant.STATUS_CODE_OK,
		Message:    constant.MESSAGE_OK,
		Data:       responseData,
		TimeStamp:  time.Now(),
	}
	json.NewEncoder(writer).Encode(response)
}

// Registerkan route
func (c ProvinceController) RegisterRouter(router *mux.Router) {
	router.HandleFunc("/provinces/GetAll", c.GetAll).Methods("GET")
}
