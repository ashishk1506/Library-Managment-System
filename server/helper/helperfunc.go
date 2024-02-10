package helper

import (
	"fmt"
	"lms/model"
	"strconv"
	"strings"
	"time"
)

func BookExist(ele int, arr []*int) bool {

	for _, y := range arr {
		if *y == ele {
			return true
		}
	}

	return false
}
func BooksExist(ele int, arr []model.BorrowedBook) int {

	for x, y := range arr {
		if y.Bookid == ele {
			return x
		}
	}

	return -1
}
func ArrayToString(a []int, delim string) string {
	return strings.Trim(strings.Replace(fmt.Sprint(a), " ", delim, -1), "[]")
}

func StringToIntArray(bidString string) ([]int, error) {
	bidStringArray := strings.Split(bidString, ",")
	bidArray := make([]int, len(bidStringArray))

	for x, y := range bidStringArray {
		val, err := strconv.ParseInt(y, 10, 32)
		if err != nil {
			return bidArray, err
		}
		bidArray[x] = int(val)
	}
	return bidArray, nil
}

func ReaderBookValue(rid int, bids []int) string {

	str := "(" + strconv.Itoa(rid) + ","
	var finalStr string
	for x, y := range bids {

		finalStr = finalStr + str + strconv.Itoa(y) + ")"
		if x != len(bids)-1 {
			finalStr = finalStr + ","
		}

	}
	return finalStr

}
func GetAmount(isdt string) (int, error) {

	issueDate, err := time.Parse("2006-01-02 15:04:05", isdt)
	if err != nil {
		return 0, err
	}
	currDate := time.Now().Format("2006-01-02 15:04:05")
	returnDate, err := time.Parse("2006-01-02 15:04:05", currDate)
	if err != nil {
		return 0, err
	}

	difference := returnDate.Sub(issueDate)
	//fmt.Printf("Days: %d\n", int64(difference.Hours()/24))
	//fmt.Printf("Minutes: %.f\n", difference.Minutes())
	fmt.Printf("Seconds: %.f\n", difference.Seconds())
	return int(difference.Seconds()), nil
}
