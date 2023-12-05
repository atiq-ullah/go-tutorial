package main

type Library struct {
	books []Book;
	count int;
}

func (l Library) getCount() int {
	return l.count
}

func (l *Library) addBook(b Book) {
	appended := append(l.books, b)
	l.books = appended
}
