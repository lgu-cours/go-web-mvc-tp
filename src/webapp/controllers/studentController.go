package controllers

import (
	"log"
	"net/http"
	
	"webapp/webutil"
	"webapp/dao"
)

func BuildNewStudentController() StudentController {
	return StudentController{}
}

type StudentController struct {
	dao.StudentDAO // Struct composition
}

func (controller *StudentController) ListHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("ListController - URL path '" + r.URL.Path )

	if r.Method == "GET" {
	    controller.processList(w,r)
	} else {
	    webutil.ErrorPage(w, "Method "+r.Method+ " is not supported");
	}
}

func (controller *StudentController) FormHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("FormController - URL path '" + r.URL.Path )

	switch r.Method {
	case "GET":
	    controller.processForm(w,r)
	case "POST":
	    controller.processPost(w,r)
	default:
	    webutil.ErrorPage(w, "Method "+r.Method+ " is not supported");
	}
}

func (controller *StudentController) processList(w http.ResponseWriter, r *http.Request) {
	// get data
	//var dao = dao.StudentDAO()
	
//	mydao := dao.StudentDAO
//	data := dao.FindAll()
	
	data := controller.StudentDAO.FindAll()
	// forward to view
	webutil.Forward(w, "templates/studentList.gohtml", data)
}

func (controller *StudentController) processForm(w http.ResponseWriter, r *http.Request) {
	
	// forward to initial page
	webutil.Forward(w, "templates/studentForm.gohtml", nil)
}

func (controller *StudentController)  processPost(w http.ResponseWriter, r *http.Request) {
	log.Print("processPost " )
	
    r.ParseForm() // Parse url parameters passed, then parse the POST body (request body)
    submit := r.Form.Get("submit")
    firstname := r.Form.Get("firstname")
    lastname  := r.Form.Get("lastname")

	log.Print("processPost submit = " + submit )
	log.Print("processPost firstname = " + firstname )
	log.Print("processPost lastname  = " + lastname )
    
    switch submit {
    	case "create":
			log.Print("doCreate" )
    	case "delete":
			log.Print("doDelete" )
    	case "update":
			log.Print("doUpdate" )
    	default:
    }
	//data := getStudentsList()
	
	// forward to initial page
	//webutil.Forward(w, "templates/studentsList.gohtml", data)

	webutil.Forward(w, "templates/studentForm.gohtml", nil)
}
