package dao

import (
	"log"
	
	"webapp/data"
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
	return this.values(data.LanguagesMap)
}

func (this *LanguageDAO) Find(code string) *entities.Language {
	log.Printf("DAO - Find(%s) ", code )
//	var language entities.Language 
//	for _, l := range languages {
//		if ( code == l.Code ) {
//			return l
//		}
//	}
//	return language
	
	language := data.LanguagesMap[code]
	return &language

}
