package controllers

import (
	"log"
	"net/http"
	
	"webapp/dao"
	//"webapp/entities"
	"webapp/webutil"
)

func LanguageListController(w http.ResponseWriter, r *http.Request) {
	log.Print("LanguageListController - URL path '" + r.URL.Path )

	switch r.Method {
	case "GET":
	    //languageGetRequest(w,r)
	    languageList(w)
//	case "POST":
//	    doPost(w,r)
	default:
	    webutil.ErrorPage(w, "Method "+r.Method+ " is not supported");
	}
}

//func languageGetRequest(w http.ResponseWriter, r *http.Request) {
//	// nothing to get from request
//    r.ParseForm() // Parse url parameters passed, then parse the response packet for the POST body (request body)
//    code := r.Form.Get("code")
//	if ( code != nil ) {
//		languageList()
//	} else {
//		languageForm(code) 
//	}
//	data := getLanguages()
//	
//}

func languageList(w http.ResponseWriter) {
	// get data
	data := dao.FindAllLanguages()
	// forward to view
	webutil.Forward(w, "templates/languageList.gohtml", data)
}

//func languageForm(code string) {
//	// get data
//	language := dao.FindLanguage(code)
//	if ( language != nil ) {
//		// forward to view
//		webutil.Forward(w, "templates/languageForm.gohtml", data)
//	}
//}
