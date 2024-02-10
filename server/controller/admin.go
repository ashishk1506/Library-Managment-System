package controller

import (
	"fmt"
	"lms/model"
	"net/http"

	"github.com/beego/beego/orm"
	"github.com/gin-gonic/gin"
)

func AddBook(c *gin.Context) {

	var newBook model.Book
	newBook.Title = c.PostForm("title")
	newBook.Author = c.PostForm("author")
	//bookCount, err := strconv.ParseInt(c.PostForm("count"), 10, 32) //int(bookCount)
	//SET BOOKCOUNT TO 1
	newBook.Count = 1
	if (newBook.Title == "") || (newBook.Author == "") {

		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Incorrect data input",
		})
		return
	}

	//MUTLIPLE BOOKS OF SAME TITLE<AUTHOR ALLOWED
	o := orm.NewOrm()
	resp, errs := o.Raw("INSERT INTO books(title,author,count) VALUES (?,?,?)", newBook.Title, newBook.Author, newBook.Count).Exec()
	if errs != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": "Internal Server Error",
		})
		fmt.Println(errs)
		return
	}
	//fmt.Println(resp)
	var bid int64 = -1
	id, err := resp.LastInsertId()
	if err != nil {
		fmt.Println(err)
	}
	bid = id
	c.JSON(http.StatusOK, gin.H{
		"message": "Book Added",
		"id":      bid,
	})
}

func UpdateBook(c *gin.Context) {

	var newBook model.Book
	bid := c.Param("id")
	newBook.Title = c.PostForm("title")
	newBook.Author = c.PostForm("author")
	//newBook.Id = bid
	if (newBook.Title == "") || (newBook.Author == "") {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Please enter valid Book details",
		})
	}
	//CHECK FOR BOOK
	o := orm.NewOrm()
	res, err := o.Raw("UPDATE books SET title=?, author=? WHERE id=?", newBook.Title, newBook.Author, bid).Exec()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": "Internal Server Error",
		})
		return
	}
	val, err := res.RowsAffected()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "Book update failed",
			"id":      bid,
			"error":   err,
		})
		return
	}
	if val == 0 {
		c.JSON(http.StatusOK, gin.H{
			"message": "Book does not exist",
			"id":      bid,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Book updated",
		"id":      bid,
	})

}
func DeleteBook(c *gin.Context) {

	bid := c.Param("id") //conevrt to int

	//CHECK FOR BOOK
	o := orm.NewOrm()

	res, err := o.Raw("DELETE FROM books WHERE id=? AND count=1", bid).Exec()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error",
		})
		return
	}
	val, err := res.RowsAffected()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "Book cannot be deleted",
			"id":      bid,
		})
		return
	}

	//IF BOOK DOES NOT EXIST
	if val != 1 {
		c.JSON(http.StatusOK, gin.H{
			"message": "Book does not exist or is Borrowed",
			"id":      bid,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Book deleted",
		"id":      bid,
	})

}
func ViewReaders(c *gin.Context) {

	//CHECK FOR BOOK
	o := orm.NewOrm()
	var readers []model.Reader
	res, err := o.Raw("SELECT rname, wallet FROM reader").QueryRows(&readers)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error",
		})
		return
	}
	if res == 0 {
		c.JSON(http.StatusOK, gin.H{
			"message": "No readers found",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"readers": readers,
	})

}
