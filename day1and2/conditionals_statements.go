package main

import (
	"fmt"
	"time"
)

func main() {

	//Using if else Condition
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			fmt.Println(i, " is Even")
		} else {
			fmt.Println(i, " is Odd")
		}
	}

	fmt.Println("Using Switch statement with fallthroug")

	time := time.Now()
	day := time.Day()

	switch day {
	case 1, 2, 3:
		fmt.Println("Day 1 or 2 or 3")
	case 4, 5, 6, 10:
		fmt.Println("Day 4 or 5 or 6 or 7")
		fallthrough
	case 11, 12, 13:
		fmt.Println("Day 11 or 12 or 13")
		fallthrough
	default:
		fmt.Println("Default Day")
	}
}
