package main

import (
	"fmt"
	"math/cmplx"
)

func main() {
	var a, b, c float64

	fmt.Scan(&a, &b, &c)

	if a == 0 {
		if b == 0 && c == 0 {
			fmt.Println("#-1")
		} else if b == 0 {
			fmt.Println("#0")
		} else {
			fmt.Printf("#1: %.3v\n", -c / b)
		}
	} else {
		d  := b * b - 4 * a * c
		x1, x2 := (complex(-b, 0) + cmplx.Sqrt(complex(d, 0))) / complex(2 * a, 0), (complex(-b, 0) - cmplx.Sqrt(complex(d, 0))) / complex(2 * a, 0)
		if d > 0 {
			fmt.Printf("#4: %.3v, %.3v\n", real(x1), real(x2))
		} else if d == 0 {
			fmt.Printf("#3: %.3v\n", real(x1))
		} else {
			fmt.Printf("#5(complex answer): %.3v, %.3v\n", x1, x2)
		}
	}
}