package dao

import (
	"log"
	"sort"

	entities "../entities"
)

// This type/struct stores no state, it’s just a collection of methods
// To get an instance of this void struct : var dao = dao.StudentDAO{}
type StudentDAO struct {
}

func (this *StudentDAO) values(m map[int]entities.Student) []entities.Student {
	var a = make([]entities.Student, len(m))
	i := 0
	for _, v := range m {
		a[i] = v
		i++
	}
	this.sortById(a)
	return a
}

func (this *StudentDAO) sortById(students []entities.Student) {
	sort.Slice(students, func(i, j int) bool {
		return students[i].Id < students[j].Id
	})
}

func (this *StudentDAO) FindAll() []entities.Student {
	log.Print("DAO - FindAll() ")
	//return this.values(data.StudentsMap)
	return this.values(studentsMap)
}

func (this *StudentDAO) Find(id int) *entities.Student {
	log.Printf("DAO - Find(%d) ", id)
	//student := data.StudentsMap[id]
	student := studentsMap[id]
	return &student
}

func (this *StudentDAO) Delete(id int) {
	log.Printf("DAO - Delete(%d) ", id)
	//delete(data.StudentsMap, id)
	delete(studentsMap, id)
}

func (this *StudentDAO) Create(student entities.Student) {
	log.Printf("DAO - Create(%d) ", student.Id)
	//data.StudentsMap[student.Id] = student
	studentsMap[student.Id] = student
}

func (this *StudentDAO) Update(student entities.Student) {
	log.Printf("DAO - Update(%d) ", student.Id)
	//data.StudentsMap[student.Id] = student
	studentsMap[student.Id] = student
}
