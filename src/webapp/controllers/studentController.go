package controllers

import (
	"log"
	"net/http"
	"strconv"
	
	"webapp/webutil"
	"webapp/dao"
	"webapp/entities"
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
	var formData StudentFormData
	formData.CreationMode = true
	formData.Student = entities.Student{} // new Student with default values ( 'zero values' )
	
	id := webutil.GetParameter(r, "id") 
	if  id != "" {
		i, _ := strconv.Atoi(id)
		student := controller.StudentDAO.Find(i)
		if student != nil {
			formData.CreationMode = false
			formData.Student = *student
		}
	} 
	
	// forward to initial page
	webutil.Forward(w, "templates/studentForm.gohtml", formData)
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
	    	controller.processCreate(w,r)
    	case "delete":
	    	controller.processDelete(w,r)
    	case "update":
			log.Print("doUpdate" )
    	default:
	    	webutil.ErrorPage(w, "Unexpected action ")
    }
}

func (controller *StudentController)  processCreate(w http.ResponseWriter, r *http.Request) {
	log.Print("processCreate " )
    r.ParseForm() // Parse url parameters passed, then parse the POST body (request body)
    
    id, _ := strconv.Atoi( r.Form.Get("id") )
    firstname := r.Form.Get("firstname")
    lastname  := r.Form.Get("lastname")
    age, _ := strconv.Atoi( r.Form.Get("age") )
    
    student := entities.Student { Id: id,
    	FirstName: firstname, 
    	LastName: lastname, 
    	Age: age }
    
    // TODO
	// controller.StudentDAO.Create(student) 

	var formData StudentFormData
	formData.CreationMode = false
	formData.Student = student 
	
	webutil.Forward(w, "templates/studentForm.gohtml", formData)
}

func (controller *StudentController)  processDelete(w http.ResponseWriter, r *http.Request) {
	log.Print("processDelete " )
    r.ParseForm() // Parse url parameters passed, then parse the POST body (request body)
    
    id, _ := strconv.Atoi( r.Form.Get("id") )
    
	log.Printf("processDelete id = %d", id )
    // TODO
	// controller.StudentDAO.Delete(id) 

	controller.processList(w, r)
}