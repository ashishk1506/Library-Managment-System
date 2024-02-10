package db

import (
	"fmt"
	"github.com/beego/beego/orm"
	"os"
	"time"
)

func dsn() string {
	username := os.Getenv("DB_USERNAME")
	hostname := os.Getenv("DB_HOSTNAME")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")
	return fmt.Sprintf("%s@tcp(%s:%s)/%s", username, hostname, port, dbname)
}

func DB() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	dsn := dsn()

	k := 0
	for k == 0 {
		err := orm.RegisterDataBase("default", "mysql", dsn, 30)
		if err == nil {
			k = 1
			break
		}
		time.Sleep(1000 * time.Millisecond)
	}
	fmt.Println("DB connection success...")
	//orm.RegisterModel(new(model.Book))
	//orm.RegisterModel(new(model.BorrowedBook))
	//orm.RegisterModel(new(model.Transaction))
	//orm.RegisterModel(new(model.Reader))
}
