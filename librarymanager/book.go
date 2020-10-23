package librarymanager

import (
	"errors"
	"strconv"
	"strings"
)

func getCategorie(code string) (Categorie, error) {
	bookCategorie, _ := strconv.Atoi(string(code[0:1]))
	categorie, err := FindCategorie(bookCategorie)
	return categorie, err
}

func getBookSize(book string) (size int, err error) {
	strsize := book[2:len(book)]
	return strconv.Atoi(strsize)
}

func isReaachBooksLimit(books []string, book string) (reach bool) {
	maxBooksLimit := 2
	countBook := 0
	reach = false
	for i := 0; i < len(books); i++ {
		if strings.Compare(books[i][0:2], book[0:2]) == 0 {
			countBook++
		}
	}
	if countBook > maxBooksLimit {
		reach = true
	}
	return
}

func decreseBooksByLimit(books []string) []string {
	for i := len(books) - 1; i >= 0; i-- {
		if isReaachBooksLimit(books, books[i]) {
			books = append(books[:i], books[i+1:]...)
		}
	}
	return books
}

func shortByCategorie(books []string) []string {
	for i := 0; i <= len(books); i++ {
		for j := 0; j < len(books)-1; j++ {
			booktempCategorie, _ := getCategorie(books[j])
			bookCategorie, _ := getCategorie(books[j+1])
			if booktempCategorie.name > bookCategorie.name {
				temp := books[j]
				books[j] = books[j+1]
				books[j+1] = temp
			}
		}
	}
	return books
}

func shortBySize(books []string) []string {
	for i := 0; i <= len(books); i++ {
		for j := 0; j < len(books)-1; j++ {
			booktempCategorie, _ := getCategorie(books[j])
			bookCategorie, _ := getCategorie(books[j+1])
			if booktempCategorie.name == bookCategorie.name {
				booktempsize, _ := getBookSize(books[j])
				booksize, _ := getBookSize(books[j+1])
				if booktempsize < booksize {
					tmp := books[j]
					books[j] = books[j+1]
					books[j+1] = tmp
				}
			}
		}
	}
	return books
}

func BooksShorter(books string) (sortedBook string, err error) {
	if len(books) > 0 {
		booksSplited := strings.Split(books, " ")

		// short by Book Categorie
		shortedByCat := shortByCategorie(booksSplited)

		// short By size in the same categorie
		shortedBySize := shortBySize(shortedByCat)

		// decrease The book if reach the Book Limit
		booksShorted := decreseBooksByLimit(shortedBySize)
		sortedBook = strings.Join(booksShorted, " ")
		err = nil
	} else {
		err = errors.New("Books must a none empty string contain list of books code that split by space")
	}
	return
}
