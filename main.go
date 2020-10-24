package main

import "example.com/ariaji25/learninggolang/webservice"

/* Basic Golang Programming.
We created go programs to resolve some case.
case 1 :
	Count polindrom numbers between two range of numbers. Polindrom numbers mostly like string polindrom, if read the number from left or right it alwasy be the same number.
case 2 :
	Shorting books in a Library, the Books will sorted by it categorie and by it size in the same categories.  Maximum 2 copies of the same books in the same bookshelf.
case 3 :
	Find a lost number from a list of sorted numbers.
case 4 :
	Create a web service for run among the cases above.
case 3 :
	Create a Docker service for the web service.
*/

func main() {
	// Run A Go Web Service
	webservice.WebService()
}
