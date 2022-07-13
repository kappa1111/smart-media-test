package test1110

import (
	"fmt"
)

func Test1110() {
	var N int
	fmt.Scan(&N)

	num1 := N / 10
	num2 := N % 10

	var nToList []int
	var goalList []int
	nToList = append(nToList, num1, num2)
	goalList = append(goalList, num1, num2)

	cnt := 1

	for true {
		newNum := nToList[0] + nToList[1]

		var newList []int

		newList = append(newList, newNum/10, newNum%10)

		var result []int

		result = append(result, nToList[1], newList[1])

		if (goalList[0] == result[0]) && (goalList[1] == result[1]) {
			break
		} else {
			nToList = result
			cnt = cnt + 1
		}
	}

	fmt.Print(cnt)

}
