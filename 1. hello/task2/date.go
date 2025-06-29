package main

import (
	"fmt"
	"time"
)

func main() {
	name := "Garik"
	date := time.Now()

	msg := fmt.Sprintf(
		"Hello, my name is %s. Current date is %s",
		name,
		date.Format("June, 2005"),
	)

	fmt.Println(msg)
}
