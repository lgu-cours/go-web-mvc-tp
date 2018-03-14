package dao

import (
	entities "../entities"
)

var languagesMap = map[string]entities.Language {
	"JA" : { "JA", "Java" },
	"JS" : { "JS", "JavaScript" },
	"PH" : { "PH", "PHP" },
	"PY" : { "PY", "Python" },
	"GO" : { "GO", "Go lang" } }
