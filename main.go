package main

import (
	"axologi/controller"
	"axologi/model/repo"
	"axologi/service"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Main function to execute the program
func main() {
	repo.ConnectDatabase()
	router := mux.NewRouter()

	//Define Services
	provinceService := service.NewProvinceService()
	loginService := service.NewLoginService()

	//Define Controller
	provinceController := controller.NewProvinceController(provinceService)
	loginController := controller.NewLoginController(loginService)

	provinceController.RegisterRouter(router)
	loginController.RegisterRouter(router)

	log.Fatal(http.ListenAndServe(":1010", router))

}
