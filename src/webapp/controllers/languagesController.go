package controllers

import (
	"log"
	"net/http"
	
	"webapp/dao"
	"webapp/entities"
	"webapp/webutil"
)

func LanguagesController(w http.ResponseWriter, r *http.Request) {
	log.Print("LanguagesController - URL path '" + r.URL.Path )

	switch r.Method {
	case "GET":
	    processLanguagesList(w,r)
//	case "POST":
//	    doPost(w,r)
	default:
	    webutil.ErrorPage(w, "Method "+r.Method+ " is not supported");
	}
}

func processLanguagesList(w http.ResponseWriter, r *http.Request) {
	// nothing to get from request

	data := getLanguages()
	
	// forward to initial page
	webutil.Forward(w, "templates/languagesList.gohtml", data)
}

func getLanguages() []entities.Language {

	return dao.GetLanguages()
}
