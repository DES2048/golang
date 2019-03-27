package core

import "fmt"

// Именованые типы, одна из фишек языка, с одной стороны оно создает алиас для уже существующего типа
// !но, вносит определенные ограничения

// Обявим 2 именнованых типа, хотя базовый тип у них 1 - float64
// СЕМАНТИКА у них разная, то есть мы не можем напрмер в одном выражении сочетать 2 этих
// типа без ЯВНОГО преобразования
type Fahrengeit float64
type Celcius float64

// ОБъявим константы температур в Цельсиях
const (
	AbsoluteZeroC Celcius= -273.5
	FreezePointC Celcius = 0
	BoilingPointC Celcius = 100
	)

// Функция принимает Celcius, чтобы передать значение базового типа
// или другого нужно явное преобразование
func CtoF(c Celcius) Fahrengeit {
	// явно приводим значения к нужному типу
	return Fahrengeit(c*9/5 + 32)
}

func FtoC(f Fahrengeit) Celcius {
	return Celcius((f-32)*5/9)
}

func DemoNamedTypes()  {
	// используем наши типы
	var absZeroF Fahrengeit = CtoF(AbsoluteZeroC)
	fmt.Println("asb zero in F", absZeroF)

	bodytemp := 36.6

	// для вызова необходимо явное преобразование, так как семантика у типов разная
	fmt.Println("body temp in F", CtoF(Celcius(bodytemp)))

	var f451 Fahrengeit = 451
	fmt.Println("451F in C:", FtoC(f451))

	// !!! операции сравнения можно проводить между именнованым типом и его базовым типом
	// НО для сравнения типов имеющих один базовый необходимо явное преобразование либо в базовый
	// либо к одному из именованых
	}


