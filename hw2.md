# Основы (ДЗ №1) - шаблон для заполнения
### ФИО: Камаргин Роман Максимович
### Класс: 9Ф

# Golang - основы (ДЗ № 1)
## 1. Линейное уравнение
```golang
package main

import "fmt"

func main() {
	var n, c int
	fmt.Scan(&n)
	for i := 1; i <= n; i++ {
		if n % i == 0 {
			c++
		}
	}
	fmt.Println(c)
}
// второй вариант в разы быстрее для больших чисел
// 	var n int
// 	fmt.Scan(&n)
// 	var power = make(map[int]int)
// loop:
// 	for {
// 		for i := 2; ; i++ {
// 			if i <= n {
// 				if n % i == 0 {
// 					power[i]++
// 					n /= i
// 					continue loop
// 				}
// 			} else {
// 				break loop
// 			}
// 		}
// 	}
// 	var c = 1
// 	for _, v := range power {
// 		c *= v + 1
// 	}
// 	fmt.Println(c)
```

## 2. Часы - 1
```golang
package main

import "fmt"

func main() {
	var n, x int

	fmt.Scan(&n)

	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	fmt.Scan(&x)

	t := "NO"
	for _, v := range arr {
		if v == x {
			t = "YES"
			break
		}
	}
	fmt.Println(t)
}
```

## 3. Часы - 2
```golang
package main

import "fmt"

func main() {
	var n, num int

	fmt.Scan(&n)

	var arr []int
	for i := 0; i < n; i++ {
		fmt.Scan(&num)
		if num < 0 {
			arr = append(arr, num)
		}
	}
	fmt.Println(arr)
}
```

## 4. Квадратное уравнение
```golang
package main

import "fmt"

func main() {
	var c, max, n, num int
	fmt.Scan(&n)
	
	for i := 0; i < n; i++ {
		fmt.Scan(&num)
		if i == 0 {
			max = num
			c++
		}
		if num > max {
			max = num
			c = 1
		} else if max == num {
			c++
		}
	}
	fmt.Println(c)
}
```

## 5. Простой for
```golang
package main

import "fmt"

func main() {
	var c, n int

	fmt.Scan(&n)

	var arr = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	for i := 1; i < n-1; i++ {
		if arr[i-1] < arr[i] && arr[i] > arr[i+1] {
			c++
		}
	}

	fmt.Println(c)
}
```

## 6. Степени двойки
```golang
package main

import "fmt"

func main() {
	var n, min1, min2, num int
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Scan(&num)
		if i == 0 {
			min1 = num
		} else if num <= min1 {
			min1, min2 = num, min1
		} else if min1 < num && (num < min2 || i == 1) {
			min2 = num
		}
	}
	fmt.Println(min1, min2)
}
```

## 7. Количество делителей
```golang
package main

import "fmt"

func main() {
	var min, max, n int
	fmt.Scan(&n)

	var arr = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	for i := 0; i < n; i++ {
		if arr[i] < arr[min] {
			min = i
		} else if arr[i] > arr[max] {
			max = i
		}
	}
	arr[max], arr[min] = arr[min], arr[max]
	fmt.Println(arr)
}
```

## 8. Кубическое уравнение
```golang
package main

import "fmt"

func main() {
	var n1, n2 int

	fmt.Scan(&n1)
	var arr1 = make([]int, n1)
	for i := 0; i < n1; i++ {
		fmt.Scan(&arr1[i])
	}

	fmt.Scan(&n2)
	var arr2 = make([]int, n2)
	for i := 0; i < n2; i++ {
		fmt.Scan(&arr2[i])
	}

	var arr3 []int
	for _, v1 := range arr1 {
		for _, v2 := range arr2 {
			if v1 == v2 {
				arr3 = append(arr3, v1)
				break
			}
		}
	}

	fmt.Println(arr3)
}
```

## 9. Числа Фибоначчи
```golang
package main

import "fmt"

func main() {
	var n, num int
	fmt.Scan(&n)

	
	var counts = make(map[int]int)
	for i := 0; i < n; i++ {
		fmt.Scan(&num)
		counts[num]++
	}

	s := "YES"
	for i := 1; i <= n; i++ {
		if !(counts[i] == 1) {
			s = "NO"
			break
		}
	}
	fmt.Println(s)
}
```
