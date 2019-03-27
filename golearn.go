package main // любая программа на go начинается с указания имени пакета

import (
	"fmt"
	"golearn/core"

	"golearn/gopl.io/ch2/tempconv"
	"golearn/gopl.io/ch4/treesort"
) // импортируем модуль fmt для вывода

var global_number = 42 // это глобальная переменная

const plank = 6.626070040e-34 // так объявляется константа, константу менять нельзя

func helloWorld() {
	fmt.Println("Hello world from Golang!")
}

func demoCoreLearn() {
	//fmt.Println()
	//core.VariablesDemo()
	//fmt.Println()
	//statements_demo()
	//fmt.Println()
	//expression_switch_case_demo()
	//fmt.Println()

	//core.ArraysDemo()

	//core.SlicesDemo()
	//core.SliceAppendDemo()
	core.SlicePassToFuncDemo()
	//core.PointersDemo()
	//core.DemoNamedTypes()

	//core.HideF()
	//core.DemoHidingInBlock()
	//core.DemoNumberOps()
	//core.DemoBool()
	//core.StringsDemo()
	//core.Hello()
	//core.DemoStruct()
	//core.DemoEmbeddedStruct()
}

func demoGlobals() {
	fmt.Println("global before +1", global_number)
	global_number++ // изменять глобальные переменные вполне можно
	fmt.Println("global++", global_number)
}

func demoTempconv() {
	fmt.Println("boiling point", tempconv.WaterBoilingPointC, "in f", tempconv.CtoF(tempconv.WaterBoilingPointC))
	fmt.Println("zero point", tempconv.AbsoluteZeroF, "in c", tempconv.FtoC(tempconv.AbsoluteZeroF))
	fmt.Println("water boiling", tempconv.WaterBoilingPointC, "in K", tempconv.CtoK(tempconv.WaterBoilingPointC))
}

func basename(s string) string {

	for i := len(s) - 1; i >= 0; i-- {
		// ищем последний слэш и отбрасываем все что перед ним как путь
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}

	// теперь ищем посл. точку как расширение
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' && i > 0 { // вторая часть чтобы сохранять имена файлов нач с точки
			s = s[:i]
			break
		}
	}

	return s
}

func sortDemo() {
	nums := []int{4, 5, 3, 2, 8, 9, 1, 0, 7, 6}

	treesort.Sort(nums)
	fmt.Println(nums)
}

// это точка входа в нашу программу
func main() {

	demoCoreLearn()

	// ch2.PrintBoilingPoints()
	// ch2.PrintFreezePoint()

	//demoTempconv()
	//fmt.Println(basename("/home/.gitignore"))
	//core.PrintfDemo()
	//fmt.Printf("%.3f\n",core.Hypot(3,4))
	//sortDemo()

	str := "Hello world!"
	fmt.Printf("%#x of %[1]T\n", []byte(str))
}
