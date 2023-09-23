package main

import "fmt"

func main() {
	var h1, m1, s1 int
	var h2, m2, s2 int

	fmt.Scan(&h1, &m1, &s1, &h2, &m2, &s2)

	fmt.Println(3600 * (h2 - h1) + 60 * (m2 - m1) + (s2 - s1))
}