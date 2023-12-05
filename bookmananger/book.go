package main

import "fmt"

type Book struct {
	title string
	author string
	year int
}

type BookManager interface {
	getTitle() string
	getAuthor() string
	getYear() int
	
	setTitle() bool
	setAuthor() bool
	setYear() bool

	toString()
}
func (b Book) getTitle() string {
	return b.title
}

func (b Book) getAuthor() string {
	return b.author
}

func (b Book) getYear() int {
	return b.year
}

func (b *Book) setTitle(title string) Book {
	b.title = title
	return *b;
}

func (b *Book) setAuthor(author string) Book {
	b.author = author
	return *b;
}

func (b *Book) setYear(year int) Book {
	b.year = year
	return *b;
}

func (b Book) print() {
	fmt.Printf("%s was written by %s and released in %d\n", b.title, b.author, b.year)
}

func newBook(title string, description string, year int) Book {
	return Book{ title, description, year}
}