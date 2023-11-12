### ФИО: Камаргин Роман Максимович
### Класс: 9Ф

# Golang - основы. Функции (ДЗ №5)
## 1. -
```golang
package main 

import (
	"math"
	"fmt"
)

type Point struct {x, y, z float64}

type Vector struct {x, y, z float64}

func CreateVector(a, b Point) Vector {
	return Vector{b.x - a.x, b.y - a.y, b.z - a.z}
}

func (p Point) Add(v Vector) Point {
	return Point{p.x + v.x, p.y + v.y, p.z + v.z}
}

func (v Vector) Length() float64 {
	return math.Sqrt(math.Pow(v.x, 2) + math.Pow(v.y, 2) + math.Pow(v.z, 2))
}

func (v Vector) IsCollinearTo(other Vector) bool {
	return v.VectorProductWith(other).Length() == 0
}

func (v Vector) Multiply(a float64) Vector {
	return Vector{v.x * a, v.y * a, v.z * a}
}

func (v Vector) Add(other Vector) Vector {
	return Vector{v.x + other.x, v.y + other.y, v.z + other.z}
}

func (v Vector) GetAngleTo(other Vector) float64 {
	scalar := v.ScalarProductWith(other)
	return math.Acos(scalar / math.Sqrt(
		            (math.Pow(v.x, 2) + math.Pow(v.y, 2) + math.Pow(v.z, 2)) * 
	                (math.Pow(other.x, 2) + math.Pow(other.y, 2) + math.Pow(other.z, 2))))
}

func (v Vector) ScalarProductWith(other Vector) float64 {
	return v.x * other.x + v.y * other.y + v.z * other.z
}

func (v Vector) VectorProductWith(other Vector) Vector {
	return Vector{v.y*other.z - v.z*other.y,
	 	    v.x*other.z - v.z*other.x,
	 		v.x*other.y - v.y*other.x}
}

func IsPolygon(arr []Vector) bool {
	var p Point
	for _, v := range arr {
		p = p.Add(v)
	}
	return p == Point{0, 0, 0}
}

```
