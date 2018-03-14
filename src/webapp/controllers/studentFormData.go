package controllers

import (
	entities "../entities"
)

type StudentFormData struct {
	CreationMode  bool
    Student       entities.Student 
    Languages     []entities.Language
}

