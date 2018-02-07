package main

import (
	"net/http"
	
	"webapp/controllers"
)

func initControllers() {

	// Specific Paths with specific controllers 


	languageController := controllers.BuildNewLanguageController() 
	http.HandleFunc("/language/list", languageController.ListHandler)
	http.HandleFunc("/language/form", languageController.FormHandler)

//	http.HandleFunc("/student/list1", controllers.StudentListController)
	
//	studentController := controllers.StudentController{} 
//	http.HandleFunc("/student/list", studentController.ListController )

	studentController := controllers.BuildNewStudentController() 
	http.HandleFunc("/student/list", studentController.ListHandler )
	http.HandleFunc("/student/form", studentController.FormHandler )

//	http.HandleFunc("/student/list", &controllers.BuildNewStudentController().ListController )

}