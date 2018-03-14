package dao

import (
	entities "../entities"
)

var studentsMap = map[int]entities.Student {
	1 : { 1, "Bart", "Simpsons",  19, "JA" },
	2 : { 2, "Luke", "Skywalker", 25, "GO" },
	3 : { 3, "Leia", "Organa",    25, "PH" } }
