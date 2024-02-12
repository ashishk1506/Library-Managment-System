package controller

import (
	"fmt"
	"lms/helper"
	"lms/model"
	"net/http"

	"github.com/beego/beego/orm"
	"github.com/gin-gonic/gin"
)

const ReaderID = 1

func BorrowBookById(c *gin.Context) {

	//count > 0 --> count -= 1 --> transaction done

	bid := c.Param("id")
	o := orm.NewOrm()

	var bookCount int
	err := o.Raw("SELECT count FROM books WHERE id=?", bid).QueryRow(&bookCount)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error",
		})
		fmt.Println(err)
		return
	}
	if bookCount == 0 {
		c.IndentedJSON(http.StatusOK, gin.H{
			"message": "All books are issued",
		})
		return
	}

	o.Raw("SET autocommit=0").Exec()
	o.Raw("BEGIN TRANSACTION").Exec()

	res, err := o.Raw("UPDATE books SET count=count-1 WHERE id=?", bid).Exec()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error",
		})
		fmt.Println(err)
		o.Raw("ROLLBACK").Exec()
		return
	}

	respp, err := o.Raw("INSERT INTO transaction(readerid,bookid) VALUES(?,?)", ReaderID, bid).Exec()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"Error": "Failed transaction during borrow",
		})
		fmt.Println(err)
		o.Raw("ROLLBACK").Exec()
		return
	}
	println(res.RowsAffected())
	println(respp.RowsAffected())
	o.Raw("COMMIT").Exec()
	o.Raw("END TRANSACTION").Exec()
	c.IndentedJSON(http.StatusOK, gin.H{
		"message": "Book Borrowed",
		"id":      bid,
	})

}

func BorrowBooks(c *gin.Context) {

	bidString := c.Query("ids")
	if bidString == "" {
		c.IndentedJSON(http.StatusNotFound, gin.H{
			"message": "Invalid input data",
		})
	}
	bidArray, err := helper.StringToIntArray(bidString)
	if err != nil {
		return
	}

	o := orm.NewOrm()
	var bookList []*int
	query := "SELECT id FROM books WHERE id IN (" + bidString + ") AND count = 1"
	res, err := o.Raw(query).QueryRows(&bookList)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{
			"message": "Books cannot be borrowed",
			"error":   err,
		})
		fmt.Println(err)
		return
	}
	fmt.Println(res)
	var notAvBookList []int
	var avBookList []int
	for _, y := range bidArray {
		if helper.BookExist(y, bookList) {
			avBookList = append(avBookList, y)
		} else {
			notAvBookList = append(notAvBookList, y)
		}
	}

	//make 2 array id match and for non match
	//return non match and borrow match

	// IF ALL BOOKS NOT BORROWED RETURN THEIR ID
	// ELSE RETURN ALL THE BORROWED BOOKS

	o.Raw("SET autocommit=0").Exec()
	o.Raw("BEGIN TRANSACTION").Exec()

	if len(avBookList) != 0 {
		query := "UPDATE books SET count = count-1 WHERE id IN (" + helper.ArrayToString(avBookList, ",") + ") "
		res, err := o.Raw(query).Exec()
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{
				"message": "Books cannot be borrowed",
				"error":   err,
			})
			fmt.Println(err)
			o.Raw("ROLLBACK").Exec()
			return
		}

		//TRANSACTION
		query = "INSERT INTO transaction(readerid,bookid) VALUE " + helper.ReaderBookValue(ReaderID, avBookList)
		respp, err := o.Raw(query).Exec()
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{
				"message": "Failed transaction during borrow",
			})
			fmt.Println(err)
			o.Raw("ROLLBACK").Exec()
			return
		}
		println(res.RowsAffected())
		println(respp.RowsAffected())
	}

	o.Raw("COMMIT").Exec()
	o.Raw("END TRANSACTION").Exec()

	if len(notAvBookList) == 0 {
		c.IndentedJSON(http.StatusOK, gin.H{
			"message":  "Books Borrowed",
			"borrowed": avBookList,
		})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{
		"message":     "All books cannot be borrowed",
		"borrowed":    avBookList,
		"notborrowed": notAvBookList,
	})

}

func ReturnBookById(c *gin.Context) {

	bid := c.Param("id")
	o := orm.NewOrm()

	o.Raw("SET autocommit=0").Exec()
	o.Raw("BEGIN TRANSACTION").Exec()

	//CHECK IF BOOK EXIST AND WAS ISSUED AND RETURN ISSUEDATE
	count := 0
	var issueDate string
	err := o.Raw("SELECT COUNT(*), issuedate FROM transaction WHERE returndate IS NULL AND bookid=? AND readerid=?  GROUP BY issuedate ", bid, ReaderID).QueryRow(&count, &issueDate)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{
			"message": "Failed updating transaction or No record found",
		})

		fmt.Println(err)
		return
	}
	if count == 0 {
		c.IndentedJSON(http.StatusNotFound, gin.H{
			"message": "Book was never issued",
		})
		return
	}

	//INCREASE COUNT OF THE BOOK
	res, err := o.Raw("UPDATE books SET count=1 WHERE id=?", bid).Exec()
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{
			"message": "Book cannot be returned",
		})
		fmt.Println(err)
		o.Raw("ROLLBACK").Exec()
		return
	}
	println(res.RowsAffected())

	//GET TRANSACTION AMOUNT
	amount, err := helper.GetAmount(issueDate)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{
			"message": "Failed calculating amount",
		})
		fmt.Println(err)
		o.Raw("ROLLBACK").Exec()
		return
	}

	//SET THE RETURN DATE TO COMPLETE TRANSACTION
	fmt.Printf("issue date is %#v\n", issueDate)
	resp, err := o.Raw("UPDATE transaction SET returndate=CURRENT_TIMESTAMP, amount=? WHERE bookid=? AND readerid=? AND returndate IS NULL", amount, bid, ReaderID).Exec()
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{
			"message": "Failed updating transaction ",
		})
		fmt.Println(err)
		o.Raw("ROLLBACK").Exec()
		return
	}
	println(resp.RowsAffected())
	o.Raw("COMMIT").Exec()
	o.Raw("END TRANSACTION").Exec()

	c.IndentedJSON(http.StatusOK, gin.H{
		"message": "Book Returned",
		"bookid":  bid,
		"amount":  amount,
	})

}

func ReturnBooks(c *gin.Context) {

	// GET BOOKS TO BE RETURNED FROM ID-> CHECK IF ALL THE IDS HAVE NULL RETURN DATE ->
	// MAKE A SEPARATE LIST FOR IDS WHICH CANNOT BE RETURNED
	// INCREASE THE COUNT OF RETURNED BOOKS
	// CALCULATE AMOUNT FOR EACH BOOK
	// SET RETURN DATE FOR EACH BOOK

	bidString := c.Query("ids")
	if bidString == "" {
		c.IndentedJSON(http.StatusNotFound, gin.H{
			"message": "Invalid input data",
		})
	}
	bidArray, err := helper.StringToIntArray(bidString)
	if err != nil {
		return
	}

	o := orm.NewOrm()
	var bookList []model.BorrowedBook
	query := "SELECT bookid,issuedate FROM transaction WHERE bookid IN (" + bidString + ") AND readerid=? AND returndate IS NULL"

	res, err := o.Raw(query, ReaderID).QueryRows(&bookList)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{
			"message": "Books cannot be returned",
			"error":   err,
		})
		fmt.Println(err)
		return
	}
	fmt.Println(res)

	var retBookList []model.BorrowedBook
	var notRetBookList []int
	for _, y := range bidArray {
		x := helper.BooksExist(y, bookList)
		if x != -1 {
			retBookList = append(retBookList, bookList[x])
		} else {
			notRetBookList = append(notRetBookList, y)
		}
	}

	o.Raw("SET autocommit=0").Exec()
	o.Raw("BEGIN TRANSACTION").Exec()

	var retBookIdList []int
	for x, y := range retBookList {
		retBookIdList = append(retBookIdList, y.Bookid)
		amt, err := helper.GetAmount(y.Issuedate)
		if err != nil {
			c.IndentedJSON(http.StatusNotFound, gin.H{
				"message": "Failed calculating amount",
			})
			return
		}
		retBookList[x].Amount = amt
	}

	if len(retBookList) != 0 {

		query := "UPDATE books SET count = 1 WHERE id IN (" + helper.ArrayToString(retBookIdList, ",") + ") "
		res, err := o.Raw(query).Exec()

		if err != nil {
			c.IndentedJSON(http.StatusNotFound, gin.H{
				"message": "Books cannot be returned",
				"error":   err,
			})
			fmt.Println(err)
			o.Raw("ROLLBACK").Exec()
			return
		}
		fmt.Println(res)

		//TRANSACTION
		for _, y := range retBookList {
			query = "UPDATE transaction SET returndate=CURRENT_TIMESTAMP, amount=? WHERE bookid =? AND readerid=? AND returndate IS NULL"
			respp, err := o.Raw(query, y.Amount, y.Bookid, ReaderID).Exec()
			if err != nil {
				c.IndentedJSON(http.StatusNotFound, gin.H{
					"message": "Transaction failed",
					"error":   err,
				})
				fmt.Println(err)
				o.Raw("ROLLBACK").Exec()
				return
			}
			println(respp.RowsAffected())
		}

	}

	o.Raw("COMMIT").Exec()
	o.Raw("END TRANSACTION").Exec()

	if len(notRetBookList) == 0 {
		c.IndentedJSON(http.StatusOK, gin.H{
			"message":  "Books Returned",
			"returned": retBookList,
		})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{
		"message":     "All books cannot be returned",
		"returned":    retBookList,
		"notreturned": notRetBookList,
	})

}

func GetBorrowedBooks(c *gin.Context) {

	o := orm.NewOrm()
	rid := c.Param("rid")
	var bookList []model.BorrowedBook
	//var issdt []string
	res, err := o.Raw("SELECT * FROM transaction INNER JOIN books ON books.id = transaction.bookid WHERE transaction.returndate IS NULL AND transaction.readerid=? ", rid).QueryRows(&bookList)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{
			"message": "Failed to fetch borrowed books",
		})
		return
	}
	fmt.Println(res)
	fmt.Printf("%+v\n", bookList)
	for i := range bookList {
		amt, err := helper.GetAmount(bookList[i].Issuedate)
		if err != nil {
			return
		}
		bookList[i].Amount = amt
		fmt.Println("amount is : ", bookList[i])
		//fmt.Printf("%+v\n", val)
	}
	c.IndentedJSON(http.StatusOK, gin.H{
		"message": "Borrowed Books",
		"books":   bookList,
	})

}
func MakePayment(c *gin.Context) {

	o := orm.NewOrm()
	bid := c.Param("id")
	//ID TID PAYEDAMOUNT -> SELECT * FROM transaction LEFT JOIN payment ON transaction.tid = payment.tid
	//WHERE readrerid=? AND returndate IS NOT NULL

	//INSERT
	// payment.tid WHERE

	//CHECK WALLET BALANCE
	walletBalance := 0
	err := o.Raw("SELECT wallet FROM reader WHERE rid=?", ReaderID).QueryRow(&walletBalance)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{
			"message": "Unable to load wallet",
		})
		return
	}
	//CHECK IF PAYMENT IS DUE
	var amount []int
	var txid []int

	_, errs := o.Raw("SELECT tid,amount FROM transaction LEFT JOIN payment ON transaction.tid = payment.transid WHERE payment.transid IS NULL AND readerid=? AND bookid=?", ReaderID, bid).QueryRows(&txid, &amount)
	if errs != nil || len(amount) == 0 {
		c.IndentedJSON(http.StatusNotFound, gin.H{
			"message": "No dues",
		})
		return
	}

	totalamt := 0
	for _, y := range amount {
		totalamt += y
	}

	if totalamt > walletBalance {
		c.IndentedJSON(http.StatusOK, gin.H{
			"message": "insuffiecient balance",
		})
		return
	}
	//MAKE PAYMENT
	o.Raw("SET autocommit=0").Exec()
	o.Raw("BEGIN TRANSACTION").Exec()

	for _, y := range txid {
		_, err = o.Raw("INSERT INTO payment(transid) VALUES(?)  ", y).Exec()
		if errs != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{
				"message": "internal server error",
			})
			o.Raw("ROLLBACK").Exec()
			return
		}
	}

	//update wallet.
	_, err = o.Raw("UPDATE reader SET wallet=? WHERE rid=?", walletBalance-totalamt, ReaderID).Exec()
	if errs != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"message": "internal server error",
		})
		o.Raw("ROLLBACK").Exec()
		return
	}
	o.Raw("COMMIT").Exec()
	o.Raw("END TRANSACTION").Exec()
	c.IndentedJSON(http.StatusOK, gin.H{
		"message":           "Payment success",
		"Balance Reaminign": walletBalance - totalamt,
	})
	return

}

func PendingPayment(c *gin.Context) {
	o := orm.NewOrm()
	var bookList []model.BorrowedBook
	res, err := o.Raw("SELECT bookid,title,amount,issuedate FROM mydb.transaction AS T LEFT JOIN mydb.payment as P ON T.tid = P.transid INNER JOIN mydb.books AS B ON T.bookid = B.id WHERE T.readerid=? AND P.transid IS NULL AND T.returndate IS NOT NULL ;",
		ReaderID).QueryRows(&bookList)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{
			"message": "Failed to fetch pending payment",
		})
		return
	}
	if res == 0 {
		c.IndentedJSON(http.StatusOK, gin.H{
			"message": "No dues",
		})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{
		"booklist": bookList,
	})

}
