package core

import "fmt"

func StatementDemo(){
	// цикл for в форме while
	i:=1
	sum:=0
	for i <= 10 {
		sum+=i
		i++
	}
	print("sum from 1 to 10:", sum, "\n")

	// for в своей каноничной форме
	for j:=10; j>=0; j-- {
		fmt.Print(j, " ")
	}
	fmt.Println()

	// условный оператор, он условный и есть
	for i:=1; i<=20; i++ {
		if i % 2 == 0 {
			fmt.Println(i, "is even")
		} else { // else, надо писать так, иначе блять ошибка!
			fmt.Println(i, "is odd")
		}
	}
}

func ExpressionSwitchCaseDemo(){
	// switch case довольно развит, есть 2 формы
	// здесь мы рассмотрим expresion switch case
	choice:=5
	switch choice {
	case 1: fmt.Println("main") // здесь не нужны break, так как в одной ветви можно указывать несколько условий
	case 2: fmt.Println("edit")
	case 3: fmt.Println("view")
	case 4,5: fmt.Println("future options")
	default:
		fmt.Println("not found")
	}
}