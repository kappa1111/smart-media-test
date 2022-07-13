package test2439

import (
	"fmt"
)

func Test2439() {
	var t int
	fmt.Scan(&t)

	for i := 1; i <= t; i++ {
		for k := t - i; k > 0; k--{
			fmt.Print(" ")
		}

		for j := 1; j <= i; j++ {
		fmt.Print("*")
		}
		fmt.Println("")
	}
}
