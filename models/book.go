package models

type Book struct {
	ID     uint   `json:"id" gorm:"primary_key"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

type CreateBookInput struct {
	Title  string `json:"title" binding:"required"`
	Author string `json:"author" binding:"required"`
}
type UpdateBookInput struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}

func NewBook(input CreateBookInput) Book {
	return Book{
		Title:  input.Title,
		Author: input.Author,
	}
}
func UpdateBook(input UpdateBookInput) Book {
	return Book{
		Title:  input.Title,
		Author: input.Author,
	}
}
