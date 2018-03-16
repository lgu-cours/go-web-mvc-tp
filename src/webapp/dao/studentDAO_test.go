package dao 

import (
	"testing"	
	"fmt"
)

// Student DAO Unit Testing 
// DAO based on a map in memory with initial list of 3 students

func TestFind(t * testing.T ) {
	    fmt.Println(t.Name()) // prints function name
	    
	    dao := NewStudentDAO()
		student := dao.Find(1)
		
		if student == nil {
			t.Error("Student #1 not found")
		} else {
			fmt.Println("OK, Student found")
			fmt.Println(student.String())
		}
}

func TestFindAll(t * testing.T ) {
	    fmt.Println(t.Name()) // prints function name
	    
	    dao := NewStudentDAO()
		students := dao.FindAll()
		
		fmt.Printf("%d students found\n", len(students))
		if len(students) != 3 {
			t.Error("3 students expected")
		} else {
			fmt.Println("OK, 3 student found")
		}
}
