package test11021

import (
	"fmt"
)

func Test11021() {
	var t int
	fmt.Scan(&t)

	var num1 int
	var num2 int

	for i := 1; i <= t; i++ {
		fmt.Scan(&num1, &num2)
		result := num1 + num2
		fmt.Printf("Case #%d: %d\n", i, result)
	}

}
