package models

type Book struct {
	ID        uint   `json:"id" gorm:"primary_key"`
	Title     string `json:"title"`
	Author    string `json:"author"`
	Author_id int    `json:"author_id"`
}

type CreateBookInput struct {
	Title     string `json:"title" binding:"required"`
	Author    string `json:"author" binding:"required"`
	Author_id int    `json:"author_id" binding:"required"`
}
type UpdateBookInput struct {
	Title     string `json:"title"`
	Author    string `json:"author"`
	Author_id int    `json:"author_id"`
}

func NewBook(input CreateBookInput) Book {
	return Book{
		Title:     input.Title,
		Author:    input.Author,
		Author_id: input.Author_id,
	}
}
func UpdateBook(input UpdateBookInput) Book {
	return Book{
		Title:  input.Title,
		Author: input.Author,
	}
}
