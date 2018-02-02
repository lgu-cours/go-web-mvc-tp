package dao

import (
	"log"
	
	"webapp/entities"
)

var languages = []entities.Language {
	entities.Language{ "JA", "Java" },
	entities.Language{ "JS", "JavaScript" },
	entities.Language{ "PH", "PHP" },
	entities.Language{ "PY", "Python" },
	entities.Language{ "GO", "Go lang" } }


func FindAllLanguages() []entities.Language {
	log.Print("DAO - FindAllLanguages() " )
	return languages
}
func FindLanguage(code string) entities.Language {
	log.Print("DAO - FindLanguage() " )
	var language entities.Language 
	for _, l := range languages {
		if ( code == l.Code ) {
			return l
		}
	}
	return language
}
