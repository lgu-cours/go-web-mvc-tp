package dao

import (
	"log"
	
	//entities "../entities"
	"webapp/entities"
	
)

// This type/struct stores no state, it’s just a collection of methods 
// To get an instance of this void struct : var dao = dao.StudentDAO{}
type LanguageDAO struct {
}

func (this *LanguageDAO) values( m map[string]entities.Language ) []entities.Language  {
	var a = make([]entities.Language, len(m))
	i := 0
    for _, v := range m {
	    a[i] = v
    	i++
    }
	return a
}

func (this *LanguageDAO) FindAll() []entities.Language {
	log.Print("DAO - FindAll() " )
	//return this.values(data.LanguagesMap)
	return this.values(languagesMap)
}

func (this *LanguageDAO) Find(code string) *entities.Language {
	log.Printf("DAO - Find(%s) ", code )
	//language := data.LanguagesMap[code]
	language := languagesMap[code]
	return &language

}
