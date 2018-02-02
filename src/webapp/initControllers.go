package main

import (
	"net/http"
	
	"webapp/controllers"
)

func initControllers() {

	// Specific Paths with specific controllers 

	//http.HandleFunc("/students", controllers.StudentsController)

	http.HandleFunc("/languages", controllers.LanguagesController)

}