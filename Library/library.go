package Library

import (
	"errors"
	"strconv"
)

type Library struct {
	books map[string]Book
}

func (l *Library) GetNextID() int {
	if l.books == nil {
		return 1
	}
	return len(l.books) + 1
}

func (l *Library) AddBook(book Book) {
	if l.books == nil {
		l.books = make(map[string]Book)
	}
	l.books[strconv.Itoa(book.ID)] = book
}

func (l *Library) BorrowBook(ID int) (Book, error) {
	if l.books == nil {
		return Book{}, errors.New("library is empty")
	}

	key := strconv.Itoa(ID)
	book, ok := l.books[key]
	if !ok {
		return Book{}, errors.New("book not found")
	}

	if book.IsBorrowed {
		return Book{}, errors.New("book is already borrowed")
	}

	book.IsBorrowed = true
	l.books[key] = book
	return book, nil
}

func (l *Library) ReturnBook(ID int) error {
	if l.books == nil {
		return errors.New("library is empty")
	}

	key := strconv.Itoa(ID)
	book, ok := l.books[key]
	if !ok {
		return errors.New("book not found")
	}

	if !book.IsBorrowed {
		return errors.New("book was not borrowed")
	}

	book.IsBorrowed = false
	l.books[key] = book
	return nil
}

func (l *Library) ListAvailableBooks() ([]Book, error) {
	if l.books == nil {
		return []Book{}, errors.New("library is empty")
	}

	var availableBooks []Book
	for _, book := range l.books {
		if !book.IsBorrowed {
			availableBooks = append(availableBooks, book)
		}
	}

	return availableBooks, nil
}

