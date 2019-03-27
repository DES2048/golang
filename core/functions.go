package core

import (
	"fmt"
	"math"
)

// объявление функции возвр значение
func Hypot(x float64, y float64) float64 {
	return math.Sqrt(x*x + y*y)
}

// если функция не имеет возвр значения, то она исп как процедура
func Hello()  {
	fmt.Println("Hello world!")
}

// можно выносить тип параметра, если несколько параметров имеют одинаковый тип
func Add(a,b int) int {
	return a + b
}

// можно создавать именованый параметр возврата
func Sub(a,b int) (z int) {
	z = a - b
	return // при вызове return вернется значение z
}
