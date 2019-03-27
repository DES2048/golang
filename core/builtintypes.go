package core

import (
	"fmt"
	"math"
	"strings"
)

func BuiltinTypesDemo(){
	// встроенные типы
	// 1. Целые числа

	// int платформозависимый знаковый такого же размера что и uint,
	// т.е на x64 - int64
	var age int = 9223372036854775807 // 2^64/2-1
	fmt.Println("signed int:", age)

	// uint платформозависмый беззнаковый x64 -> uiint64
	var ptr uint = 18446744073709551615 // 2^64-1
	fmt.Println("uint:", ptr)

	// uintptr - видимо для 64-битных указателей платформозависимый
	var large_ptr uintptr = 18446744073709551615 // 2^64-1 -maxval
	fmt.Println("uintptr:", large_ptr)

	// платформонезависимые целые
	// {u}int{bit} -
	// 		где u - признак беззнакового типа
	//		bit - количество бит 8,16,32 или 64
	//		например uint8 - беззнаковый байтовый тип, еще этот тип называется byte
	//		int32 - знаковый 4б целый
	//		int64 - соответсвует знаковому long в java

	var space_code uint8 = 20
	var symbol byte = 64 // то же что и uint64
	fmt.Println("uint8 and byte:", space_code, symbol)

	// 2. floats

	var pi float32 = 3.14159
	// лучше использовать float64, на 64x
	var electron_mass float64 = 9.10938356e-31 // инженерный формат

	fmt.Println("pi is", pi)
	fmt.Println("electron mass is", electron_mass)

	//3. есть еще комплексные числа complex64 и complex128

	// 4. bool - true or false
	var isTrue bool = true
	var isFalse bool = false

	fmt.Println("bools are", isTrue, isFalse)

	// 5. string - строковый тип, неизменяемый, по сути хранит массив байт
	var name string = "Oleg"
	var surname = "Podzorov"

	// строки можно конкатенировать
	var fullname = name + " " + surname

	// можно исп втроенную функцию len() для определения длины
	var fname_len = len(fullname)

	fmt.Println("My name is", fullname, "len() is", fname_len)
}

// операции с числами
func DemoNumberOps() {
	// арифметические операции
	// как обычно, сложение, вычитание, умножение, деление(ЦЕЛОчисленное) для целых
	// взятие остатка
	fmt.Println("1+2:", 1+2)
	fmt.Println("3-4:", 3-4)
	fmt.Println("6*6:", 6*6)
	fmt.Println("5/2", 5/2)
	fmt.Println("5%2", 5%2)

	// для чисел все арифметические операции выполняются по модулю, таким образом при
	// переполнении происходит wrap around
	max_uint := uint32(math.Pow(2, 32) - 1)
	var uint32_zero uint32 = 0
	fmt.Println("uint32 max val:", max_uint, "+1, overflow, wraps around:", max_uint+1)
	fmt.Println("uint 0-1:", uint32_zero-1)

	// побитовые операторы
	// | побитовое или, исп. в основном чтобы устновить биты в шаблоне
	// установим 32  бит в числе
	var bitfield uint = 0
	bitfield = bitfield | (1 << 32)
	fmt.Println("set 32", bitfield)

	// & побитовое И, используется чтобы проверить установлены конкретные биты или нет
	var bit32 uint = 1 << 32
	var bit4 uint = 1 << 4
	fmt.Println("bit32 set:", bitfield&bit32 == bit32, "bit4 set", bitfield&bit4 == bit4)

	// ^ xor исключающее или, позволяет сбрасывать нужный бит
	// !!! как для арифметических так и побитовых опероров доступна форма +=
	bitfield |= bit4            // установим 4 бит и сбросим 32
	bitfield = bitfield ^ bit32 // сбросим 32 бит
	fmt.Println("bit32 set:", bitfield&bit32 == bit32, "bit4 set", bitfield&bit4 == bit4)

	// ~ побитовый not, инвертирует все биты
}

func DemoBool(){
	// 4. bool - true or false
	var isTrue bool = true
	var isFalse bool = false

	fmt.Println("bools are", isTrue, isFalse)

	// для булева типа существует стандартный набор лог опертаторов !, && ||
	// бинарные операторы вычисляются по короткому циклу

	str := "Welcome"

	// короткий цикл исп. для подобных проверок
	isCapitalize := str != "" && strings.ToUpper(str)[0] == str[0];
	fmt.Println("is ", str, " capitalized?", isCapitalize)

	// нет неяного преобразования bool to int поэтому если часто исп то можно напиманть
	// спец функции
	fmt.Println("true is", BtoI(true))
	fmt.Println("0 is", ItoB(0))
}

func BtoI(value bool) int{
	if value {
		return 1
	} else {
		return 0
	}
}

func ItoB(value int) bool {
	return value != 0
}

func StringsDemo() {
	// string - тип строки, строка в unicode и неизменяемая
	fmt.Println("*** String demo ***")

	//!!! прикол намба уан. len() возвращает не длину строки, а КОЛИЧЕСТВО БАЙТ
	// ЧТО МОЖЕТ ПРИВЕСТИ К ТОМУ ЧТО ДЛЯ unicod длина мб больше
	hello := "Привет!"
	fmt.Println("len of", hello, "is", len(hello)) // даст 13 символов

	//!!! прикол намба ту, если обратится к строке по индексу то мы получим знаечние БАЙТА по
	// этому индексу
	fmt.Println("П это 2 байта:", hello[0], hello[1])

	// обращение за границы диапазона приведет к ошибке

	// однако подстроку можно получить с помощью такой вот операции получения подстроки
	//!!! НО ВСЕ РАВНО ИНДЕКСЫ ПРИВЯЗЫВАЮТСЯ К БАЙТАМ
	fmt.Println(hello[0:2]) // семантика как в py первая граница включается, посл. исключется,
	// если опустить границы то они счит соотв за 0 символ и за всю длину строки

	w := "Welcome!"
	fmt.Println(w[0:1], w[0:], w[:3])

	//  в строковых литералах можно передавать байты в hex и oct
	fmt.Println("hex symbol", "\x54", "oct symbol", "\122")

	// есть также raw строки они закл. в `` и выводятся как есть
	couplet :=
		`Hello
Dolly!
	Make me sun
	Dolly`

	fmt.Println(couplet)

	// можно включать коды символов unicode в hex в виде \uFFFF или \UFFFFFFFF

	myname := "\u041e\u043b\u0435\u0433";
	fmt.Println("My name is", strings.ToUpper(myname))
}


