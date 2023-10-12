### ФИО: Камаргин Роман Максимович
### Класс: 9Ф

# Golang - основы. Функции (ДЗ №3)
## 1. -
```golang
package main

import "fmt"

func MyRange(start, stop, step int) []int {
	var arr []int
	for i := start; i < stop; i+=step {
		arr = append(arr, i)
	}
	return arr
}

func MyPower(a float64, e int) float64 {
	var retval float64 = 1
	if e >= 0 {
		for i := 0; i < e; i++ {
			retval *= a
		}
	} else {
		for i := 0; i > e; i-- {
			retval /= a
		}
	}
	return retval
}

func MyFactorial(n int) int {
	var retval = 1
	for i := 1; i <= n; i++ {
		retval *= i
	}
	return retval
}

func MyGcd(a, b int) int {
	for b != 0 {
		a, b = b, a % b
	}
	return a
}

func IsSymmetric(arr []int) bool {
	for i := 0; i < len(arr) / 2; i++ {
		if arr[i] != arr[len(arr)-1-i] {
			return false
		}
	}
	return true
}

func MyReverse(arr []int) []int {
	for i := 0; i < len(arr) / 2; i++ {
		arr[i], arr[len(arr)-1-i] = arr[len(arr)-1-i], arr[i]
	}
	return arr
}

func MyShorten(a, b int) (int, int) {
	p := a / MyGcd(a, b)
	q := b / MyGcd(a, b)
	if q < 0 {
		p, q = -p, -q
	}
	return p, q
}

func IsPrime(a int) bool {
	for i := 2; i < a; i++ {
		if a % i == 0 {
			return false
		}
	}
	return true
}

func MySet(arr []int) []int {
	var m = make(map[int]bool)
	var retval []int
	for _, v := range arr {
		if !m[v] {
			retval = append(retval, v)
		}
		m[v] = true
	}
	return retval
}

func IsSet(arr []int) bool {
	var m = make(map[int]bool)
	for _, v := range arr {
		if m[v] {
			return false
		}
		m[v] = true
	}
	return true
	// return len(MySet(arr)) == len(arr)
}


```
