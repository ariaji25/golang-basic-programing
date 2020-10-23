package librarymanager

import "fmt"

/* Libray Manager.
A Go Program to manage books in a bookshelft at a Libray. Books in the bookshelft will sorted by it categories and it size in the same categorie.
The input of this program is a list of book codes. The value of the first index in  the book code is defined the book categorie, the value of second index is
the name of categorie, and the rest of the code is the book size.
- input : a string that contains the list of book codes.
*/
func LibrayManager(input string) {
	fmt.Println("[Libray Manager]")
	fmt.Println("Input : ", input)
	output, _ := BooksShorter(input)
	fmt.Println("Ouput : ", output)
}
