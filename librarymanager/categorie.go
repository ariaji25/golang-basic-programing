package librarymanager

import "errors"

/*Define a struct of Book Categorie model */
type Categorie struct {
	id   int
	name string
}

// Create the list of categories
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

/* Find Categorie.
Function to find the detail of category by given the id of categorie.
- input : id of categoris with integer type
- output : categorie thah found by the id, and the error if there is an error while find the categorie */
func FindCategorie(id int) (cat Categorie, err error) {
	// init the error
	err = errors.New("Categorie not found")
	// do looping in the list categories to find the categorie by the given id
	for _, c := range BookCategories {
		// chekc if this id categorie is equal to given id
		if c.id == id {
			// found the categorie
			cat = c
			// set the error is null
			err = nil
		}
	}
	return
}
