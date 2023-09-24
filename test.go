package main

import (
	"fmt"
	"time"
)

func f1(n int) {
	var c int
	for i := 1; i <= n; i++ {
		if n % i == 0 {
			c++
		}
	}
	fmt.Println(c)
}

func f2(n int) {
	var power = make(map[int]int)
loop:
	for {
		for i := 2; ; i++ {
			if i <= n {
				if n % i == 0 {
					power[i]++
					n /= i
					continue loop
				}
			} else {
				break loop
			}
		}
	}
	var c = 1
	for _, v := range power {
		c *= v + 1
	}
	fmt.Println(c)
}

func main() {
	var n int
	fmt.Scan(&n)

	now2 := time.Now()
	f2(n)
	fmt.Printf("2: %f\n", time.Since(now2).Seconds())

	fmt.Println()

	now1 := time.Now()
	f1(n)
	fmt.Printf("1: %f\n", time.Since(now1).Seconds())
}