// challenges/types-composite/begin/main.go
package main

import (
	"fmt"

	"github.com/davecgh/go-spew/spew"
)

// define an author type with a name
// simple struct. we need just specific key and value
type author struct {
	name string
}

// define a book type with a title and author
// struct embedding (instead of inheritance)
type book struct {
	title string
	author
}

// define a library type to track a list of books
// we will use map for it. dynemic keys and unknown length
// value will be a slice of book
type library map[string][]book

// define addBook to add a book to the library
// defining a function reciever for library
func (l library) addBook(b book) {
	l[b.author.name] = append(l[b.author.name], b)
}

// define a lookupByAuthor function to find books by an author's name
func (l library) lookupByAuthor(name string) []book {
	return l[name]
}
func (b book) printBook() {
	fmt.Println("name: " + b.title + ". author: " + b.name + ".")
}

func main() {
	vFrankel, hShapiro := author{name: "Victor Frankel"}, author{name: "Haim Shapiro"}

	// create a new library
	library := make(library)

	// add 2 books to the library by the same author
	library.addBook(book{title: "Man's Search For Meaning", author: vFrankel})
	library.addBook(book{title: "Yes To Life", author: vFrankel})

	// add 1 book to the library by a different author
	library.addBook(book{title: "Go, My Son", author: hShapiro})

	// dump the library with spew
	spew.Dump(library)

	// lookup books by known author in the library
	books := library.lookupByAuthor(vFrankel.name)

	// print out the first book's title and its author's name
	if len(books) != 0 {
		books[0].printBook()
	}

}
