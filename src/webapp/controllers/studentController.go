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

	log.Print("processPost submit = " + submit )
    
    switch submit {
    	case "create":
	    	controller.processCreate(w,r)
    	case "delete":
	    	controller.processDelete(w,r)
    	case "update":
			controller.processUpdate(w,r)
    	default:
	    	webutil.ErrorPage(w, "Unexpected action ")
    }
}

func (controller *StudentController)  processCreate(w http.ResponseWriter, r *http.Request) {
	log.Print("processCreate " )
//    r.ParseForm() // Parse url parameters passed, then parse the POST body (request body)
//    
//    id, _ := strconv.Atoi( r.Form.Get("id") )
//    firstname := r.Form.Get("firstname")
//    lastname  := r.Form.Get("lastname")
//    age, _ := strconv.Atoi( r.Form.Get("age") )
//    
//    student := entities.Student { Id: id,
//    	FirstName: firstname, 
//    	LastName: lastname, 
//    	Age: age }
    
    student := controller.buildStudent(r)
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

func (controller *StudentController)  processUpdate(w http.ResponseWriter, r *http.Request) {
	log.Print("processUpdate " )
    student := controller.buildStudent(r)
    
    // TODO
	// controller.StudentDAO.Delete(id) 

	var formData StudentFormData
	formData.CreationMode = false
	formData.Student = student 
	
	webutil.Forward(w, "templates/studentForm.gohtml", formData)
}

func (controller *StudentController)  buildStudent(r *http.Request) entities.Student {
	log.Print("buildStudent " )
    r.ParseForm() // Parse url parameters passed, then parse the POST body (request body)
    
    id, _ := strconv.Atoi( r.Form.Get("id") )
    firstname := r.Form.Get("firstname")
    lastname  := r.Form.Get("lastname")
    age, _ := strconv.Atoi( r.Form.Get("age") )
    
    student := entities.Student { Id: id,
    	FirstName: firstname, 
    	LastName: lastname, 
    	Age: age }
    
	return student
}
