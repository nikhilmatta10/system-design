package main

import "fmt"

type Iterator interface {
	hasNext() bool
	Next() interface{}
}

type Book struct{
	Name string
}

type BookIterator struct {
	books    []Book
	position int
}

func (bi *BookIterator) hasNext() bool {
	if len(bi.books) > bi.position {
		return true
	}
	return false
}

func (bi *BookIterator) Next() interface{} {
	book := bi.books[bi.position]
	bi.position += 1
	return book
}

type BookCollection struct {
	books []Book
}

func (bc *BookCollection) addBook(book Book) {
	bc.books = append(bc.books, book)
}

func (bc *BookCollection) getBooks() []Book {
	return bc.books
}

func (bc *BookCollection) createIterator() Iterator {
	return &BookIterator{
		books: bc.books,
	}
}

func main() {

	bookCollection := BookCollection{}
	bookCollection.addBook(Book{Name: "C++ Book"})
	bookCollection.addBook(Book{Name: "Java Book"})
	bookCollection.addBook(Book{Name: "Python Book"})

	iterator := bookCollection.createIterator()
	
	for iterator.hasNext() {
		book := iterator.Next()
		fmt.Printf("Book Title : %v \n", book)
	}
}
