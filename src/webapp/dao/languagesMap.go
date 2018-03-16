package dao

import (
	//entities "../entities"
	"webapp/entities"
)

// map of LANGUAGES 
// private ( visibility restricted to the current package )

var languagesMap = map[string]entities.Language {
	"JA" : { "JA", "Java" },
	"JS" : { "JS", "JavaScript" },
	"PH" : { "PH", "PHP" },
	"PY" : { "PY", "Python" },
	"GO" : { "GO", "Go lang" } }
