package main

import (
	"net/http"
	
	"webapp/controllers"
)

func initControllers() {

	// Specific Paths with specific controllers 


	http.HandleFunc("/language/list", controllers.LanguageListController)

//	http.HandleFunc("/student/list1", controllers.StudentListController)
	
//	studentController := controllers.StudentController{} 
//	http.HandleFunc("/student/list", studentController.ListController )

	studentController := controllers.BuildNewStudentController() 
	http.HandleFunc("/student/list", studentController.ListHandler )
	http.HandleFunc("/student/form", studentController.FormHandler )

//	http.HandleFunc("/student/list", &controllers.BuildNewStudentController().ListController )

}