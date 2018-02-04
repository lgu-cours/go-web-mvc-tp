package controllers

import (
	"log"
	"net/http"
	
	"webapp/webutil"
)

func StudentFormController(w http.ResponseWriter, r *http.Request) {
	log.Print("colorController - URL path '" + r.URL.Path )

	switch r.Method {
	case "GET":
	    processStudentsForm(w,r)
	case "POST":
	    doPost(w,r)
	default:
	    webutil.ErrorPage(w, "Method "+r.Method+ " is not supported");
	}
}

func processStudentsForm(w http.ResponseWriter, r *http.Request) {
	
	// forward to initial page
	//webutil.Forward(w, "templates/studentsList.gohtml", data)
}

func doPost(w http.ResponseWriter, r *http.Request) {
//    r.ParseForm() // Parse url parameters passed, then parse the POST body (request body)
//    firstname := r.Form.Get("id")
//    lastname  := r.Form.Get("lastname")

	//data := getStudentsList()
	
	// forward to initial page
	//webutil.Forward(w, "templates/studentsList.gohtml", data)
}
