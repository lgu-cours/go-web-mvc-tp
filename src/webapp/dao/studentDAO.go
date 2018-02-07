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

func (dao *StudentDAO) values( m map[int]entities.Student ) []entities.Student  {
	var a = make([]entities.Student, len(m))
	i := 0
    for _, v := range m {
	    a[i] = v
    	i++
    }
	return a
}

func (dao *StudentDAO) FindAll() []entities.Student {
	log.Print("DAO - FindAll() " )
	return dao.values(data.StudentsMap)
}

func (dao *StudentDAO) Find(id int) *entities.Student {
	log.Printf("DAO - Find(%d) ", id )
	student := data.StudentsMap[id]
	return &student
}

func (dao *StudentDAO) Delete(id int) {
	log.Printf("DAO - Delete(%d) ", id )
	delete(data.StudentsMap, id)
}

func (dao *StudentDAO) Create(student entities.Student) {
	log.Printf("DAO - Create(%d) ", student.Id )
	data.StudentsMap[student.Id] = student 
}

func (dao *StudentDAO) Update(student entities.Student) {
	log.Printf("DAO - Update(%d) ", student.Id )
	data.StudentsMap[student.Id] = student 
}
