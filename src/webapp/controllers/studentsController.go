package controllers

import (
	"log"
	"net/http"
	
	"webapp/webutil"
)

func StudentsController(w http.ResponseWriter, r *http.Request) {
	log.Print("colorController - URL path '" + r.URL.Path )

	switch r.Method {
	case "GET":
	    processStudentsList(w,r)
//	case "POST":
//	    doPost(w,r)
	default:
	    webutil.ErrorPage(w, "Method "+r.Method+ " is not supported");
	}
}

func processStudentsList(w http.ResponseWriter, r *http.Request) {
	// nothing to get from request
	//data := getStudentsList()
	
	// forward to initial page
	//webutil.Forward(w, "templates/studentsList.gohtml", data)
}

//func getStudentsList() [] {
//	// TODO
//	return nil
//}
