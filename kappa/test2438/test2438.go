package test2438

import (
	"fmt"
)

func Test2438() {
	var t int
	fmt.Scan(&t)

	for i := 1; i <= t; i++ {
		for j := 1; j <= i; j++ {
		fmt.Print("*")
		}
		fmt.Println("")
	}
}
