package http

type BookDTO struct {
	Name   string
	Author string
	Pages  int
}

type CompleteDTO struct {
	Complete bool
}
