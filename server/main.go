package main

import (
	controller "lms/controller"
	"lms/db"
	"lms/middlewares"
	model2 "lms/model"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	//CUSTOMER
	//LOGIN ()
	//REGISTER ()

	//GET ALL BOOKS -> SEARCH IN FRONTEND , HIGHLIGTH ONLY NOT ISSUED ONES
	// SELECT * FROM books LIMIT 10 OFFSET 0

	//ISSUE BOOK BY ID (ARRAY, CHECK IF CAN BE ISSUED THEN ISSUE ONE BY ONE)-> ADDED TO CART
	//SELECT * FROM books WHERE isISSUED == false
	//UPDATE books SET isISSUED == true WHERE id IN []
	//INSERT INTO transaction VALUES('CUSTID','BOOKID','TIMESTAMP')

	//RETURN BOOK BY ID -> SAME AS ISSUE BOOK BY ID
	//UPDATE book SET isISSUE = false
	//UPDATE transaction SET return date = timestamp

	db.DB()
	//model2.ConnectDataBase()
	router := gin.Default()
	//router.Use(middlewares.JwtAuthMiddleware())

	public := router.Group("/api")
	public.POST("/register", controller.Register)
	public.POST("/login", controller.Login)

	//UN-SECURED
	router.GET("/", controller.Ping)
	router.GET("/books", controller.GetBooks)
	router.GET("/books/:id", controller.GetBookById)
	router.GET("/books/borrow/:id", controller.BorrowBookById)
	router.GET("/books/borrow-m", controller.BorrowBooks)
	router.GET("/books/return/:id", controller.ReturnBookById)
	router.GET("/books/return-m", controller.ReturnBooks)
	router.GET("/books/pending-payment", controller.PendingPayment)
	router.GET("/books/payment/:id", controller.MakePayment)
	router.GET("/reader/borrowedBooks/:rid", controller.GetBorrowedBooks)
	//router.GET("/reader/returnedBooks/:rid" getReturnedBooks) //null are deleted books

	//PROTECTED
	protected := router.Group("/admin")
	protected.Use(middlewares.AuthMiddleware)
	protected.POST("/add", controller.AddBook)
	protected.GET("/delete/:id", controller.DeleteBook)
	protected.GET("/update/:id", controller.UpdateBook)
	protected.GET("/view-readers", controller.ViewReaders)

	router.GET("/reader/register", registerReader)

	router.Run(":8080")

}

func registerReader(c *gin.Context) {
	var newBook model2.Book
	newBook.Title = c.PostForm("name")
	//o := orm.NewOrm()
}

// HELPER
