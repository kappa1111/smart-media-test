package test10871

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Test10871() {
	br := bufio.NewReader(os.Stdin)
	str, _ := br.ReadString(('\n'))
	str2, _ := br.ReadString(('\n'))
	
	blank := strings.Fields(str)
	blank2 := strings.Fields(str2)
	
	var blink1 = []int{}
	var a = []int{}

	for i :=0; i<=len(blank)-1;i++{
		b1, _ := strconv.Atoi(blank[i])
		blink1 = append(blink1, b1)
	}

	for i :=0; i<=len(blank2)-1;i++{
		b2, _ := strconv.Atoi(blank2[i])
		a = append(a, b2)
	}
	
	n := blink1[0]
	x := blink1[1]

	for i := 0; i <= n-1; i++ {
		if x > a[i]	{
			fmt.Printf("%d ", a[i])
		}
	}
}
