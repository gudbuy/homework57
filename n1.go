package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	if a == 0 && b == 0 {
		fmt.Println("INF")
	} else if a == 0 {
		fmt.Println("NO")
	} else if b % a == 0 {
		fmt.Println(-b / a)
	} else {
		fmt.Println("NO")
	}
}