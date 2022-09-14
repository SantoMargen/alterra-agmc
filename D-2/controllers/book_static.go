package controllers

import (
	"D-2/models"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

var books = []models.Book{
	{ID: 1, Title: "Book 1", Page: 254, Category: "Sains", Author: "Author 1", Price: 56000},
	{ID: 2, Title: "Book 2", Page: 540, Category: "Technology", Author: "Author 2", Price: 680000},
	{ID: 3, Title: "Book 3", Page: 311, Category: "Filsafat", Author: "Author 3", Price: 120000},
}

func CreateBook(c echo.Context) error {
	var bodyBytes []byte
	if c.Request().Body != nil {
		bodyBytes, _ = ioutil.ReadAll(c.Request().Body)
	}
	var data models.Book
	err := json.Unmarshal(bodyBytes, &data)
	if err != nil {
		return err
	}
	if len(books) == 0 {
		data.ID = 1

	} else {
		data.ID = books[len(books)-1].ID + 1
	}

	c.Bind(&data)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"StatusCode": 201,
		"status":     "succes",
		"message":    "Success create new Book",
		"book":       data,
	})
}

func GetBooks(c echo.Context) error {
	return c.JSON(http.StatusOK, books)
}
func GetBooksById(c echo.Context) error {
	bookId, _ := strconv.Atoi(c.Param("id"))
	for _, book := range books {
		if book.ID == uint(bookId) {
			return c.JSON(http.StatusOK, map[string]interface{}{
				"code":    http.StatusOK,
				"status":  "succes",
				"message": "Success get book",
				"book":    book,
			})
		}
	}
	return c.JSON(http.StatusBadRequest, map[string]interface{}{
		"status":  "error",
		"message": "Book not Found",
		"code":    http.StatusBadRequest,
	})
}
func EditPriceBook(c echo.Context) error {
	bookId, err := strconv.Atoi(c.Param("id"))

	if err != nil || bookId > len(books) {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":   "Error",
			"code":     http.StatusBadRequest,
			"messsage": "Book not found",
		})
	}
	body, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  "error",
			"code":    http.StatusBadRequest,
			"message": "Cannot parse payload",
		})
	}
	bookInput := models.Book{}
	json.Unmarshal(body, &bookInput)
	bookUpdated := models.Book{}

	for i := 0; i < len(books); i++ {
		if books[i].ID == uint(bookId) {
			if bookInput.Price != 0 {
				books[i].Price = bookInput.Price
			}
			bookUpdated = books[i]
		}
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  "error",
		"code":    http.StatusOK,
		"message": "Success",
		"book":    bookUpdated,
	})
}
func DeleteBook(c echo.Context) error {
	bookId, _ := strconv.Atoi(c.Param("id"))
	for _, book := range books {
		if book.ID == uint(bookId) {
			return c.JSON(http.StatusOK, map[string]interface{}{
				"code":    http.StatusOK,
				"status":  "succes",
				"message": "Book has been deleted",
			})
		}
	}
	return c.JSON(http.StatusBadRequest, map[string]interface{}{
		"status":  "error",
		"message": "Book not Found",
		"code":    http.StatusBadRequest,
	})
}
