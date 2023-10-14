### ФИО: Камаргин Роман Максимович
### Класс: 9Ф

# Golang - основы. Рекурсия, указатели (ДЗ №4)
## 1. -
```golang
package main

import "fmt"

func MyFact(n int) int {
	if n == 1 {
		return 1
	}
	return MyFact(n-1) * n
}

func IsPresent(x int, arr []int) bool {
	if len(arr) == 0 {
		return false
	}
	if arr[0] == x {
		return true
	}
	return IsPresent(x, arr[1:])
}

func GetMax(arr []int) int {
	var max = arr[0]
	var r func(int)
	r = func(i int) {
		if i >= len(arr) {
			return
		}
		if arr[i] > max {
			max = arr[i]
		}
		r(i+1)
	}
	r(1)
	return max
}

func EvenAmount(arr []int) int {
	var retval int
	var r func(int)
	r = func(i int) {
		if i >= len(arr) {
			return
		}
		if arr[i] % 2 == 0 {
			retval++
		}
		r(i+1) 
	}
	r(0)
	return retval
}

func MyRange(start, stop, step int) []int {
	var retarr []int
	var r func(int)
	r = func(el int) {
		if el >= stop {
			return
		}
		retarr = append(retarr, el)
		r(el + step)
	}
	r(start)
	return retarr
}

func IsPowerOfTwo(x float64) bool {
	var r func(float64) bool
	r = func(n float64) bool {
		if n > x {
			return false
		} else if x == n {
			return true
		}
		return r(n*2)
	}
	return r(1)
}

func IsSymmetric(arr []int) bool {
	if len(arr) < 2 {
		return true
	}
	if arr[0] != arr[len(arr)-1] {
		return false
	}
	return IsSymmetric(arr[1:len(arr)-1])
}

func SumOfDigits(x int) int {
	if x == 0 {
		return 0
	}
	return SumOfDigits(x / 10) + x % 10
}

func IsPrime(x int) bool {
	if x == 1 {
		return true
	}
	var r func(int) bool
	r = func(i int) bool {
		if i == 1 {
			return true
		}
		if x % i == 0 {
			return false
		}
		return r(i-1)
	}
	return r(x-1)
}

func MyReverse(arr []int) []int {
	var r func(arr []int)
	r = func(arr []int) {
		if len(arr) < 2 {
			return
		}
		arr[0], arr[len(arr)-1] = arr[len(arr)-1], arr[0]
		r(arr[1:len(arr)-1])
	}
	r(arr)
	return arr
}

func BinSearch_1(arr *[]int, x int) bool {
	if len(*arr) == 0 {
		return false
	}
	low, high := 0, len(*arr)-1
	mid := (low + high) / 2
	if (*arr)[mid] < x {
		*arr = (*arr)[mid+1:]
		return BinSearch_1(arr, x)
	} else if (*arr)[mid] > x {
		*arr = (*arr)[:high]
		return BinSearch_1(arr, x)
	}
	return true
}

func MySum(arr []int, sum *int) {
	for _, v := range arr {
		*sum+=v
	}
}

func MyStrangeArithmeticMean(arr []*int) {
	var arithm int
	for _, v := range arr {
		arithm += *v
	}
	for _, v := range arr {
		*v = arithm / len(arr)
	}
}

func BubbleSort(arr *[10]float64) {
	for i := 0; i < len(*arr); i++ {
		for j := 0; j < len(*arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func BinSearch_2(arr *[]int, x int) bool {
	low, high := 0, len(*arr) - 1
	for low <= high {
		mid := (low + high) / 2
		if x > (*arr)[mid] {
			low = mid + 1
		} else if x < (*arr)[mid] {
			high = mid - 1
		} else {
			return true
		}
	}
	return false
}

```
