package controller

import (
	"axologi/constant"
	"axologi/helper"
	"axologi/model"
	"axologi/service"
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type LoginController struct {
	service service.LoginService
}

// Construct
func NewLoginController(s service.LoginService) *LoginController {
	return &LoginController{
		service: s,
	}
}

func (c LoginController) Login(writter http.ResponseWriter, request *http.Request) {
	defer request.Body.Close()
	// Set Content-Type header
	writter.Header().Set("Content-Type", "application/json")

	//Decode JSONRequest
	var loginInput model.LoginModel
	decoder := json.NewDecoder(request.Body)
	if err := decoder.Decode(&loginInput); err != nil {
		//Error decode
		response := helper.GeneralWrapper{
			StatusCode: constant.STATUS_CODE_INTERNAL_SERVER_ERROR,
			Message:    constant.MESSAGE_INTERNASL_SERVER_ERROR,
			Data:       err,
			TimeStamp:  time.Now(),
		}
		json.NewEncoder(writter).Encode(response)
		return
	}

	responseData, err := c.service.LoginService(loginInput)

	if err != nil {
		//Error
		response := helper.GeneralWrapper{
			StatusCode: constant.STATUS_CODE_NOT_FOUD,
			Message:    constant.MESSAGE_NOT_FOUND,
			Data:       nil,
			TimeStamp:  time.Now(),
		}
		json.NewEncoder(writter).Encode(response)
		return
	}

	//Success
	response := helper.GeneralWrapper{
		StatusCode: constant.STATUS_CODE_OK,
		Message:    constant.MESSAGE_OK,
		Data:       responseData,
		TimeStamp:  time.Now(),
	}
	json.NewEncoder(writter).Encode(response)
}

func (c LoginController) Register(writter http.ResponseWriter, request *http.Request) {
	defer request.Body.Close()

	// Set Content-Type header
	writter.Header().Set("Content-Type", "application/json")

	//Decode data req ke JSON
	var registerInput model.RegisterModel
	decoder := json.NewDecoder(request.Body)

	if err := decoder.Decode(&registerInput); err != nil {
		//Error decode
		response := helper.GeneralWrapper{
			StatusCode: constant.STATUS_CODE_INTERNAL_SERVER_ERROR,
			Message:    constant.MESSAGE_INTERNASL_SERVER_ERROR,
			Data:       err,
			TimeStamp:  time.Now(),
		}
		json.NewEncoder(writter).Encode(response)
		return
	}

	registerInput.IpAddress = helper.GetIPAddress(request)
	registerInput.DeviceID = uuid.New().String()
	result, err := c.service.RegisterService(registerInput)

	if err != nil {
		//Error data sudah ada
		response := helper.GeneralWrapper{
			StatusCode: constant.STATUS_CODE_CONFLICT,
			Message:    constant.MESSAGE_CONFLICT,
			Data:       err.Error(),
			TimeStamp:  time.Now(),
		}
		json.NewEncoder(writter).Encode(response)
		return
	}

	//Success
	response := helper.GeneralWrapper{
		StatusCode: constant.STATUS_CODE_OK,
		Message:    constant.MESSAGE_OK,
		Data:       result,
		TimeStamp:  time.Now(),
	}
	json.NewEncoder(writter).Encode(response)
}

func (c LoginController) LogOut(writter http.ResponseWriter, request *http.Request) {
	return
}

func (c LoginController) RegisterRouter(router *mux.Router) {
	router.HandleFunc("/axologi/login", c.Login).Methods("POST")
	router.HandleFunc("/axologi/register", c.Register).Methods("POST")
}
