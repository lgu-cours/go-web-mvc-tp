package entities

import "fmt"

type Language struct {
        Code  string
        Name  string
}

func (this Language) ToString() string {
    return fmt.Sprintf(
    	"[%s, %s]", 
    	this.Code, 
    	this.Name )
}