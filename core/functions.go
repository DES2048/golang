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
func Hello() {
	fmt.Println("Hello world!")
}

// можно выносить тип параметра, если несколько параметров имеют одинаковый тип
func Add(a, b int) int {
	return a + b
}

// можно создавать именованый параметр возврата, он не должен совпадать с одним
// из параметров функции
func Sub(a, b int) (z int) {
	z = a - b
	return // при вызове return вернется значение z
}

func triplet(v int) (int, int, int) {
	return v - 1, v, v + 1
}

func DemoFunctions() {
	// в сигнатуру функции входят типы и число паметров и типы возвр значений
	// т.е след функции эквиваленты по типу

	funcA := func(a, b int) int {
		return 1
	}

	funcb := func(a int, b int) (c int) {
		c = a + b
		return
	}

	// так объявлется переменная хран адрес функции
	var fpoint func(int, int) int

	fpoint = funcA
	fmt.Println("fpoint is", fpoint)
	fpoint = funcb
	fmt.Println("fpoint is", fpoint)

	// ссылки на функцию можно сравнивать только на nil

	// функция может возвращать несколько значений, ненужные возвраты можно
	// отбрасывать с помощью _

	left, _, right := triplet(4)
	fmt.Println("left", left, "right", right)

	// можно отборосить хоть все возвр значения, такое исп например когда нужно вызвать
	// функцию, как процедуру а возврат не нужен совсем

}

// демонстрация замыкания, с помощью которого можно создать
// что то вроде генератора
func createGenerator(initial int) func() int {

	return func() int {
		retval := initial // адрес initial замыкается на локальную переменную
		// и становится атрибутом функции, как например поле у класса
		// каждый вызов изменяет initial
		initial++
		return retval
	}
}

// демонстрация некорректного захвата перменной цикла
func DemoClosures() {
	addrs := []string{"156.98.78.9", "123.0.0.10", "10.56.07.6"}

	var conns []func() string

	// некорректный захват
	for _, addr := range addrs {
		// так как захватывается адрес переменной, то в итоге после окончания цикла
		// в захваченной переменной будет последнее значение
		conns = append(conns, func() string { return addr })
	}

	fmt.Println("Incorrect loop closure")
	for _, f := range conns {
		fmt.Println(f())
	}

	conns = nil

	// корректный захват способ первый
	for _, addr := range addrs {
		// переприсваем addr, делая ее лок переменной, для каждой итерации
		// и тогда замыкание будет работать как надо
		addr := addr
		// теперь замыкается адрес лок переменной каждой итерации цикла
		conns = append(conns, func() string { return addr })

	}

	fmt.Println("Correct loop closure way 1")
	for _, f := range conns {
		fmt.Println(f())
	}

	conns = nil

	for i := 0; i < len(addrs); i++ {
		conns = append(conns, func() string {
			// здесь последнее значение i = 3, что за границей массива,
			// хороший способ выстрелить себе в ногу
			return addrs[i]
		})
	}

	conns = nil

	// исправляем
	for i := 0; i < len(addrs); i++ {
		i := i
		conns = append(conns, func() string {
			return addrs[i]
		})
	}

	fmt.Println("Correct loop closure way 2")
	for _, f := range conns {
		fmt.Println(f())
	}

	conns = nil

	// причем переменная динамически меняется поэтому вызывая созд функцию
	// еще в цикле, значения будут нужными
	fmt.Println("addr in loop")

	for i, addr := range addrs {
		conns = append(conns, func() string { return addr })
		fmt.Println(conns[i]())
	}

	//  а вот потом уже нет!
	fmt.Println("after loop")
	for _, f := range conns {
		fmt.Println(f())
	}

	// как бороться, если такое "динамическое" поведение ни к чему, замыкаемся через
	// Лок переменную, можно прям в самой функции

}

// можно создавать функции с переменным количеством аргументов
func sum(vals ...int) (sum int) {
	// внтури vals представлен как срез
	// sum иниц неявно в самом начале функции своим zero value
	for _, v := range vals {
		sum += v
	}
	return
}

func DemoVariableFunc() {

	// вызываем функцию
	fmt.Println("sum of 3,4,5:", sum(3, 4, 5))

	// функцию можно вызывать и БЕЗ аргументов, и такую ситуацию нужно
	// корректно обрабатывать в своих функциях
	fmt.Println("sum of 0", sum())

	// можно забрасывать в такую функцию и срез
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println("sum of", nums, "is", sum(nums...))

	//!!! НО функция с переменным количеством аргументов отличается по типу и сигнатуре
	// От функции которая принимает срез
	// из этого следует что их типы не совместимы друг с другом

	sum2 := func(val []int) (sum int) {
		for _, v := range val {
			sum += v
		}
		return
	}

	fmt.Printf("variable func: %T, slice func %T\n", sum, sum2)

	//!!!НО если в вариабельную функцию передать срез, то копирования, как можно ожидать
	// не происходит, т.е изменения отразятся на исходном срезе. Хотя по ... можно подумать
	// что это что-то типа распаковки в python

	add2 := func(vals ...int) []int {
		for i := range vals {
			///!!! изменения будут видны и на исходном срезе
			vals[i] += 2
		}

		return vals
	}

	fmt.Println("nums before add2", nums)
	_ = add2(nums...)
	fmt.Println("nums after add2", nums)

}
