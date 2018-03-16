package controllers

import (
//	entities "../entities"
	"webapp/entities"
)

type StudentFormData struct {
	CreationMode  bool
    Student       entities.Student 
    Languages     []entities.Language
}

