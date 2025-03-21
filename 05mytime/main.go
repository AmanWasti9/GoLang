package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to Time study of golang")

	presentTime := time.Now()
	fmt.Println(presentTime)
	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))
	// Standard date time to format in go lang
	// 01-02-2006 15:04:05 Monday

	createdDate := time.Date(2020, time.January, 10, 15, 23, 0, 0, time.UTC)
	fmt.Println(createdDate)
	fmt.Println(createdDate.Format("01-02-2006 15:04:05 Monday"))

}
