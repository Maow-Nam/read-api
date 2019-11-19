package api

import (
	"fmt"
	"net/http"

	"github.com/Maow-Nam/read-api/model"
	"github.com/labstack/echo"
)

//PostBook ..
func (db *MongoDB) PostBook(c echo.Context) error {

	book := &model.Book{}
	if err := c.Bind(book); err != nil {
		fmt.Println("In c.Bind Error ", err)
		return err
	}

	return c.JSON(http.StatusOK, book)
}
