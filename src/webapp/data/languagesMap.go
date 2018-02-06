package data

import (
	"webapp/entities"
)

var LanguagesMap = map[string]entities.Language {
	"JA" : { "JA", "Java" },
	"JS" : { "JS", "JavaScript" },
	"PH" : { "PH", "PHP" },
	"PY" : { "PY", "Python" },
	"GO" : { "GO", "Go lang" } }

func LanguagesList() []entities.Language  {
	m := LanguagesMap
	var a = make([]entities.Language, len(m))
	i := 0
    for _, v := range m {
	    a[i] = v
    	i++
    }
	return a
}