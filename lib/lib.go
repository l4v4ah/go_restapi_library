package lib

import (
	"math/rand"
)

type Lib struct {
	books map[int]Book
}

func NewLib() *Lib {
	return &Lib{
		books: make(map[int]Book),
	}
}

func (l *Lib) AddBook(book Book) error {
	var numId int
	for {
		numId = rand.Intn(10000) + 1
		if _, ok := l.books[numId]; !ok {
			break
		}
	}

	if book.Name == "" || book.Author == "" || book.Pages <= 0 {
		return ErrBookNotRequest
	}

	l.books[numId] = book

	return nil
}

func (l *Lib) ListBook() map[int]Book {
	tmp := make(map[int]Book, len(l.books))

	for k, v := range l.books {
		tmp[k] = v
	}

	return tmp
}

func (l *Lib) ListCompletedBooks() map[int]Book {
	tmp := make(map[int]Book)

	for k, v := range l.books {
		if v.Completed {
			tmp[k] = v
		}
	}

	return tmp
}

func (l *Lib) ListUncompletedBooks() map[int]Book {
	tmp := make(map[int]Book)

	for k, v := range l.books {
		if !v.Completed {
			tmp[k] = v
		}
	}

	return tmp
}

func (l *Lib) ListAuthorBooks(author string) map[int]Book {
	tmp := make(map[int]Book)

	for k, v := range l.books {
		if v.Author == author {
			tmp[k] = v
		}
	}

	return tmp
}

func (l *Lib) GetBook(id int) (Book, error) {
	tmp := Book{}
	flag := false

	for k, v := range l.books {
		if k == id {
			tmp = v
			flag = true
			break
		}
	}
	if flag {
		return tmp, nil
	} else {
		return tmp, ErrBookNotFound
	}
}

func (l *Lib) DeleteBook(id int) error {
	for k, _ := range l.books {
		if k == id {
			delete(l.books, k)
			return nil
		}
	}
	return ErrBookNotFound
}

func (l *Lib) CompleteBook(id int) (Book, error) {
	v, ok := l.books[id]
	if !ok {
		return Book{}, ErrBookNotFound
	}

	v.Complete()

	l.books[id] = v

	return v, nil
}

func (l *Lib) UncompleteBook(id int) (Book, error) {
	v, ok := l.books[id]
	if !ok {
		return Book{}, ErrBookNotFound
	}

	v.Uncomplete()

	l.books[id] = v

	return v, nil
}
