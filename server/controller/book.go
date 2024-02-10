package controller

import (
	"lms/model"
	"net/http"

	"github.com/beego/beego/orm"
	"github.com/gin-gonic/gin"
)

func Ping(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{
		"PING": "PONG",
	})
}

func GetBooks(c *gin.Context) {

	o := orm.NewOrm()
	var mp []model.Book
	res, err := o.Raw("SELECT * FROM books").QueryRows(&mp)
	if err != nil || res == 0 {
		c.IndentedJSON(http.StatusOK, gin.H{
			"message": "No books found",
		})
	}
	c.IndentedJSON(http.StatusOK, gin.H{
		"books": mp,
	})

}

func GetBookById(c *gin.Context) {

	id := c.Param("id")
	o := orm.NewOrm()
	var mp model.Book
	err := o.Raw("SELECT * FROM books WHERE id=?", id).QueryRow(&mp)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"message": "Book not found",
		})
		return
	}
	c.IndentedJSON(http.StatusOK, mp)

}
