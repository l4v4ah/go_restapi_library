package lib

import (
	"time"
)

type Book struct {
	Name      string
	Author    string
	Pages     int
	Completed bool

	CreatedAt   time.Time
	CompletedAt *time.Time
}

func NewBook(name string, author string, pages int) Book {
	return Book{
		Name:      name,
		Author:    author,
		Pages:     pages,
		Completed: false,
		CreatedAt: time.Now(),
	}
}

func (b *Book) Complete() {
	completeTime := time.Now()

	b.Completed = true
	b.CompletedAt = &completeTime
}

func (b *Book) Uncomplete() {
	b.Completed = false
	b.CompletedAt = nil
}
