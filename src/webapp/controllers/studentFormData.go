package controllers

import (
	"webapp/entities"
	"webapp/data"
)

type StudentFormData struct {
	CreationMode  bool
    Student       entities.Student 
    Languages     []entities.Language
}

func NewStudentFormData(creationMode bool, student entities.Student ) StudentFormData {
	// New structure
	var formData StudentFormData
	// Init structure fields
	formData.CreationMode = creationMode
	formData.Student      = student 
	formData.Languages    = data.LanguagesList()  // The current list of languages
	// Return structure
	return formData
}