package lib

import "errors"

var ErrBookNotFound error = errors.New("Книга не найдена!")
var ErrBookNotRequest error = errors.New("Не корректно указаны данные о книге!")
