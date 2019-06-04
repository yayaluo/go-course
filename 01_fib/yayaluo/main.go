package main

import (
	"fmt"
	"io"
	"os"
)

var out io.Writer = os.Stdout

func fib(n int) {
	x, y := 0, 1
	if n > 0 {
		for i := 1; i <= n; i++ {
			x, y = y, x+y
			fmt.Fprintf(out, "%d\n", x)
		}
	} else {
		for i := 1; i <= -n; i++ {
			x, y = y, x+y
			if i%2 == 0 {
				fmt.Fprintf(out, "%d\n", -x)
			} else {
				fmt.Fprintf(out, "%d\n", x)
			}
		}
	}
}

func main() {
	fib(7)
}
