package Library

import (
	"errors"
	"strconv"
)

type Library struct {
	a map[string]Book
}

func (l *Library) AddBook(book Book) {
	if l.a == nil {
		l.a = make(map[string]Book)
	}
	l.a[strconv.Itoa(book.ID)] = book
}

func (l *Library) BorrowBook(ID int) (Book, error) {
	if l.a == nil {
		return Book{}, errors.New("library is empty")
	}

	key := strconv.Itoa(ID)
	book, ok := l.a[key]
	if !ok {
		return Book{}, errors.New("book not found")
	}

	if book.IsBorrowed {
		return Book{}, errors.New("book is already borrowed")
	}

	book.IsBorrowed = true
	l.a[key] = book
	return book, nil
}

func (l *Library) ReturnBook(ID int) error {
	if l.a == nil {
		return errors.New("library is empty")
	}

	key := strconv.Itoa(ID)
	book, ok := l.a[key]
	if !ok {
		return errors.New("book not found")
	}

	if !book.IsBorrowed {
		return errors.New("book was not borrowed")
	}

	book.IsBorrowed = false
	l.a[key] = book
	return nil
}

func (l *Library) ListAvailableBooks() ([]Book, error) {
	if l.a == nil {
		return []Book{}, errors.New("library is empty")
	}

	var books []Book
	for _, book := range l.a {
		if !book.IsBorrowed {
			books = append(books, book)
		}
	}

	return books, nil
}
