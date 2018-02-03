package main

import (
	"net/http"
	
	"webapp/controllers"
)

func initControllers() {

	// Specific Paths with specific controllers 


	http.HandleFunc("/language/list", controllers.LanguageListController)

	http.HandleFunc("/student/list", controllers.StudentListController)
}