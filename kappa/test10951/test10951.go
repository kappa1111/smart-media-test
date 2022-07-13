package test10951

import (
	"bufio"
	"fmt"
	"os"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func Test10951() {
	defer writer.Flush()

	var a, b int

	for true {
		val, _ := fmt.Fscanln(reader, &a, &b)

		if val != 2 {
			break
		} else {
			fmt.Fprintln(writer, a+b)
		}
	}
}
