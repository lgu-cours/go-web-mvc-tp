package controllers

import (
	"log"
	"net/http"
	
	"webapp/dao"
	"webapp/webutil"
)

func StudentListController(w http.ResponseWriter, r *http.Request) {
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
	// get data
	var dao = dao.StudentDAO{}
	data := dao.FindAll()
	// forward to view
	webutil.Forward(w, "templates/studentList.gohtml", data)
}
