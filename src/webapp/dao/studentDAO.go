package dao

import (
	"log"
	
	"webapp/data"
	"webapp/entities"
)

// This type/struct stores no state, it’s just a collection of methods 
// To get an instance of this void struct : var dao = dao.StudentDAO{}
type StudentDAO struct {
}

func (dao *StudentDAO) FindAll() []entities.Student {
	log.Print("DAO - FindAll() " )
	return data.StudentsData
}

func (dao *StudentDAO) Find(id int) *entities.Student {
	log.Printf("DAO - Find(%d) ", id )
	for _, student := range data.StudentsData {
		if ( id == student.Id ) {
			return &student
		}
	}
	return nil
}
