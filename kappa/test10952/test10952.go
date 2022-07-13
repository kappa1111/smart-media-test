package test10952

import (
	"fmt"
)

func Test10952() {
	var num1 int
	var num2 int
	var result int

	for true {
		fmt.Scan(&num1, &num2)

		result = num1 + num2

		if result == 0 {
			break
		} else {
			fmt.Println(result)
		}
	}
}
