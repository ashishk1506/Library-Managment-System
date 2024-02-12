package model

import "github.com/dgrijalva/jwt-go"

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

type Book struct {
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Count  int    `json:"count"`
}

type BorrowedBook struct {
	Bookid    int    `json:"bookid"`
	Issuedate string `json:"issuedate"`
	Amount    int    `json:"amount"`
}

type Transaction struct {
	Tid        int    `json:"tid"`
	ReaderId   int    `json:"readerid"`
	BookId     int    `json:"bookid"`
	IssueDate  string `json:"issuedate"`
	ReturnDate string `json:"returndate"`
	Amount     int    `json:"amount"`
}

type Reader struct {
	Rid    int    `json:"rid"`
	Rname  string `json:"rname"`
	Wallet int    `json:"wallet"`
}
