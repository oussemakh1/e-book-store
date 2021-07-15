package services

import (
	"rest-api/models"
	"rest-api/repositories"
)

func CreateBook(input models.CreateBookInput) (newBook models.Book, err error) {
	//New book instance
	book := models.NewBook(input)
	//Create new book
	newBook, err = repositories.CreateBook(book)
	if err != nil {
		return book, err
	}
	return newBook, nil

}

func FindBooks() ([]models.Book, error) {
	var books []models.Book
	//Get all books
	books, err := repositories.FindBooks()
	if err != nil {
		return books, err
	}
	return books, nil

}

func FindBook(id int) (models.Book, error) {
	var book models.Book
	//Get all books
	book, err := repositories.FindBook(id)
	if err != nil {
		return book, err
	}
	return book, nil

}

func UpdateBook(input models.UpdateBookInput, id int) (newBook models.Book, err error) {
	//New book instance
	book := models.UpdateBook(input)

	//update book
	newBook, err = repositories.UpdateBook(book, id)
	if err != nil {
		return book, err
	}
	return newBook, nil

}

func DeleteBook(id int) error {
	//delete book
	if err := repositories.DeleteBook(id); err != nil {
		return err
	}
	return nil

}
