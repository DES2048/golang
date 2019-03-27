package core

import "fmt"

func PointersDemo() {
	// указатели похожи на таковые в C с некоторыми отличиями
	// При объявлении указатель имеет значение nil, а не непоняно какое значение в памяти
	var intPtr *int
	fmt.Println("IntPtr not init:", intPtr)

	// чтобы писвоить знаечние указателю исп оператор адреса
	myVar := 5
	myVar++
	intPtr = &myVar
	// а чтобы получить значение(разыменовать указатель) исп *ptr
	fmt.Println("intPtr->myVar:", intPtr, "value:", *intPtr)
	*intPtr++
	fmt.Println("myVar after ++ from intPtr:", myVar)

	// указатели равны ттт когда оба ссылаются на один и тот же адресс или оба равны nil

	// Функция совершенно спокойно может возвращать указатель на свою лок переменную
	// и он при каждом вызове естественно будет разным
	fmt.Println("creating 5 int_ptrs...")
	for i:=0; i<5; i++ {
		fmt.Println("int_ptr ", createIntPtr())
	}
	fmt.Println()

	// передаем адрес на переменную в функцию принимающую указатель
	incInt(&myVar)
	fmt.Println("myVar after incr(*int) passing address:", myVar)

	// ну или сам указатель
	incInt(intPtr)
	fmt.Println("myVar after incr(*int) passing pointer:", myVar)

	// также можно создавать перемнную с помощью new()
	// !!! Однако семантика тут иная чем в C++, new как и var решают где выделить
	// память на куче или на стеке, исходя из того "сбегает" ли переменная из функции или нет
	intPtr = new(int) // будет создана на стеке, так как не передается вовне
	// если бы мы возвращали адрес или указатель на переменную из функции то, память для нее выделилась
	// бы стеке
	fmt.Println("intPtr after new():", intPtr, "val:", *intPtr)

}

// Из функции спокойно можно вернуть указатель на лок переменную
func createIntPtr() *int {
	val:=1
	return &val // возвращаем адрес что и будет указателем
}

// функция может и принимать указатель, чтобы модифицировать переданное значение
func incInt(val *int){
	*val++ // инкрементируем значение
}