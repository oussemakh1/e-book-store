package repositories

import (
	"log"
	"rest-api/database"
	"rest-api/models"
)

func CreateBook(book models.Book) (newBook models.Book, err error) {
	//Get db config
	db, err := database.GetDB("DEV")
	defer db.Close()
	if err != nil {
		return book, err
	}
	//Create a new book
	if err := db.Table("books").Create(&book).Error; err != nil {
		return book, err
	}
	log.Println("----", book)
	return book, nil
}

func FindBooks() (books []models.Book, err error) {
	//Get db config
	db, err := database.GetDB("DEV")
	defer db.Close()
	if err != nil {
		return books, err
	}
	//Find books
	if err := db.Table("books").Find(&books).Error; err != nil {
		return books, err
	}
	return books, nil
}

func FindBook(id int) (book models.Book, err error) {
	//Get db config
	db, err := database.GetDB("DEV")
	defer db.Close()
	if err != nil {
		return book, err
	}
	//Find book by id
	if err := db.Table("books").Where("id = ?", id).Find(&book).Error; err != nil {
		return book, err
	}
	return book, nil
}

func UpdateBook(updateBook models.Book, id int) (book models.Book, err error) {
	//Get db config
	db, err := database.GetDB("DEV")
	defer db.Close()
	if err != nil {
		return book, err
	}
	// Get model if exist
	if err := db.Table("books").Where("id = ?", id).First(&book).Error; err != nil {
		return book, err
	}
	//update book
	if err := db.Model(&book).Updates(&updateBook).Error; err != nil {
		return book, err
	}
	return book, nil
}

func DeleteBook(id int) error {
	//Get db config
	db, err := database.GetDB("DEV")
	defer db.Close()
	if err != nil {
		return err
	}
	// Get model if exist
	var book models.Book
	if err := db.Table("books").Where("id = ?", id).First(&book).Error; err != nil {
		return err
	}
	//delete book
	if err := db.Delete(&book).Error; err != nil {
		return err
	}
	return nil
}

func FindMyBooks(id int)(books []models.Book, err error) {
	db, err := database.GetDB("DEV")
	defer db.Close()
	if err != nil {
		return books, err
	}
	if err := db.Table("books").Where("author_id = ?",id).Find(&books).Error; err != nil {
		return books, err
	}
	return books, nil
}

func FindMyBook(author_id,id int)(book models.Book, err error) {
	db, err := database.GetDB("DEV")
	defer db.Close()
	if err != nil {
		return book, err
	}

	if err := db.Table("books").Where("id = ? AND author_id = ?",id,author_id).Find(&book).Error; err != nil {
		return book, err
	}
	return book, nil
}