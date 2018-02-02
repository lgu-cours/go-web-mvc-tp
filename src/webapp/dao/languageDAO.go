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


func GetLanguages() []entities.Language {
	log.Print("DAO - GetLanguages() " )
	return languages
}
