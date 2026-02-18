package main

import "fmt"

type Book struct {
	Name string
}

type BookCollection struct {
	books []Book
}

func (bc *BookCollection) addBook(book *Book) {
	bc.books = append(bc.books, *book)
}

func (bc *BookCollection) getBooks() []Book {
	return bc.books
}

func main() {

	bookCollection := BookCollection{}
	bookCollection.addBook(&Book{Name: "C++ Book"})
	bookCollection.addBook(&Book{Name: "Java Book"})
	bookCollection.addBook(&Book{Name: "Python Book"})

	for idx, _ := range bookCollection.getBooks() {
		fmt.Println(bookCollection.getBooks()[idx])
	}
}