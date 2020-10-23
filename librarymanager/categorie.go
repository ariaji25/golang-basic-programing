package librarymanager

import "errors"

type Categorie struct {
	id   int
	name string
}

var BookCategories = []Categorie{
	{id: 6, name: "Applied Sciences (600)"},
	{id: 7, name: "Arts (700)"},
	{id: 0, name: "General (000)"},
	{id: 9, name: "Geography, History (900)"},
	{id: 4, name: "Language (400)"},
	{id: 8, name: "Literature (800)"},
	{id: 1, name: "Philosophy, Psychology (100)"},
	{id: 2, name: "Religion (200)"},
	{id: 5, name: "Science, Math (500)"},
	{id: 3, name: "Social Sciences (300)"},
}

func FindCategorie(id int) (cat Categorie, err error) {
	err = errors.New("Categorie not found")
	for _, c := range BookCategories {
		if c.id == id {
			cat = c
			err = nil
		}
	}
	return
}
