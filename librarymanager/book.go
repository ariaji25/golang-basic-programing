package librarymanager

import (
	"errors"
	"strconv"
	"strings"
)

/* Get Category.
Function to get Book categorie from book code.
- input : book code type string
- output : book categorie, and error if find erro while find the category */
func getCategory(code string) (Categorie, error) {
	// get book categorie in the first index of book code
	bookCategorie, _ := strconv.Atoi(string(code[0:1]))
	// find categorie by call function FindCategorie from categorie.go
	categorie, err := FindCategorie(bookCategorie)
	return categorie, err
}

/* Get The Book Size.
Function to get the Book size form the book code.
- input : book code to find the book size, type string
- output : return size of book, and error if find error while find the book size*/
func getBookSize(book string) (size int, err error) {
	// get the book size in code, book size is in start from 2
	strsize := book[2:len(book)]
	return strconv.Atoi(strsize)
}

/* Is Reach the Maximum Books in the same Bookshelft.
Function to chekc if there is same books with total more than 2 books in the same bookshelft.
- input :
	- books : list of bokks.
	- book : book to check
- output : true if book is reached Books Limit type boolean*/
func isReaachBooksLimit(books []string, book string) (reach bool) {
	// init the maximum same books
	maxBooksLimit := 2
	// counter
	countBook := 0
	// is reached
	reach = false
	// do looping to find the same book
	for i := 0; i < len(books); i++ {
		// chekc the book in index i and index i+1 i same
		if strings.Compare(books[i][0:2], book[0:2]) == 0 {
			// counter + 1
			countBook++
		}
	}
	// if hte counter value is more than 2, set reached is true
	if countBook > maxBooksLimit {
		reach = true
	}
	return
}

/* Decres Books If reach the Maximum Books in Bookshelft.
Function to remove book from a bookshelft if the same books is reached the maximum books in the Bookshelf.
- input : list of books
- output : list of books*/
func decreseBooksByLimit(books []string) []string {
	for i := len(books) - 1; i >= 0; i-- {
		if isReaachBooksLimit(books, books[i]) {
			books = append(books[:i], books[i+1:]...)
		}
	}
	return books
}

/* Short By Categorie.
Function to sorting the books by it categorie. Use bubble sort algorithm, with alphabetic sort.
- inpuy : list of books code
- output : sorted list of books code that sorted by it categorie*/
func shortByCategorie(books []string) []string {
	// do first loop of bubble sort
	for i := 0; i <= len(books); i++ {
		// do second loop of bubble sort
		for j := 0; j < len(books)-1; j++ {
			// find the book categorie of book in index j
			booktempCategorie, _ := getCategory(books[j])
			// find the book categorie of book in index j+1
			bookCategorie, _ := getCategory(books[j+1])
			// compare the both books categorie with alphabetic comparation and ascending sort
			if booktempCategorie.name > bookCategorie.name {
				// swap position of the both bools
				temp := books[j]
				books[j] = books[j+1]
				books[j+1] = temp
			}
		}
	}
	return books
}

/* Sorted By Size.
Function to sort the list with bubble sort algorithm, the list will sorted by it size in the same categorie.
- input : list string that contains list of books code
- output : sorted list string contain list of books code */
func sortBySize(books []string) []string {
	// do first looping for bubble sort
	for i := 0; i <= len(books); i++ {
		// do second looping for bubble sort
		for j := 0; j < len(books)-1; j++ {
			// get the book categories of book in index j
			booktempCategorie, _ := getCategory(books[j])
			// get the book categories of book in index j+1
			bookCategorie, _ := getCategory(books[j+1])
			// compare if book categorie of both books is same
			if booktempCategorie.name == bookCategorie.name {
				// get the book size of book in index j
				booktempsize, _ := getBookSize(books[j])
				// get the book size of book in index j+1
				booksize, _ := getBookSize(books[j+1])
				// compare the both books size, where is the smaalest size (descending sort)
				if booktempsize < booksize {
					// swap position of the both bools
					// save the book in index j to temp variable
					tmp := books[j]
					// set the book in index j to book in index j+1
					books[j] = books[j+1]
					// set the book in index j+1 to book in temp variable
					books[j+1] = tmp
				}
			}
		}
	}
	return books
}

/* BookShorter.
Function to manage the books at a bookshelft, the Books will sorted by Book categories and than the books in the same categorie will sorted by it size.
In a bookshelft, just allowed 2 copies of the same books. If books in the same categorie and the same name, the books recognized as the same book ex. (0A21, 0A14). If books
with the same name but with diferent categorie, the books recognized as diferent books ex. (9A21 and 0A22), */
// - input : string that contain list of books code
// - output : string that contain the sorted list of books code, return with error if there is an error while sorting the books
func BooksShorter(books string) (sortedBook string, err error) {
	if len(books) > 0 {
		booksSplited := strings.Split(books, " ")

		// short by Book Categorie
		shortedByCat := shortByCategorie(booksSplited)

		// short By size in the same categorie
		shortedBySize := sortBySize(shortedByCat)

		// decrease The book if reach the Book Limit
		booksShorted := decreseBooksByLimit(shortedBySize)
		sortedBook = strings.Join(booksShorted, " ")
		err = nil
	} else {
		err = errors.New("Books must a none empty string contain list of books code that split by space")
	}
	return
}
