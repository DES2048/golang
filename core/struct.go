package core

import (
	"fmt"
	"time"
)

// структура исп для объединения нескольких полей в одно целое
// поля начинающиеся с большой буквы явл экспортируемыми
// в функциии структура передается ПО ЗНАЧЕНИЮ
type Employee struct {
	ID       int
	Name     string
	Address  string
	DoB      time.Time
	Position string
	Salary   int
	Manager  *Employee // для агрегации лучше использовать указатели
}

type Point2d struct {
	X, Y int
}

func DemoStruct() {

	fmt.Println("*** demo structs ***")

	// при объявлении перем с типом структуры, поля структуры зап значениями по умолчанию
	var dilbert Employee
	fmt.Println("dilbert zero value:", dilbert)

	// обращение к свойствам происходит через точку
	dilbert.ID = 3456
	dilbert.Name = "Dilbert Adams"
	dilbert.Position = "developer"
	fmt.Println("dilbert non zero:", dilbert)

	// можно получать адреса на конкретные поля структуры и изменять их
	pPos := &dilbert.Position
	*pPos = "Senior " + *pPos
	fmt.Println("Became a senior over poiner to pos:")
	fmt.Println(dilbert)

	// через точку можно обр к полям если это указатель на структуру
	var pEmployer *Employee = &dilbert
	pEmployer.Salary = 3000
	fmt.Println("over pointer:", pEmployer)

	// можно создавать структуру с помощью литерала
	p1 := Point2d{1, 1} // без указания имен полей, необх заполнять ВСЕ поля в порядке их следования
	fmt.Println(p1)
	// с помощью имен можно устанавливать только опр поля
	mine := &Employee{Name: "Oleg Podzorov", Salary: 5000, Position: "Senior Potato"}
	fmt.Println(mine)

	// использование & перед литералом выше поз сразу создать указатель на структуру,
	// т.е эквивалентно
	var emp *Employee = new(Employee)
	*emp = Employee{Name: "Super hr", Salary: 1000}
	///!!! интересно что при таком создании выводится данные структуры, без необходимости разыменования
	// как было выше
	fmt.Println(emp)

	// 2 структуры равны ттт когда равны их поля
	p2 := Point2d{1, 1}
	fmt.Println(p1 == p2)
}

type Person struct {
	Name string
	Age  int
}

type User struct {
	Person // встраивание структуры, теперь к полям Person можно обращатся через точку
	// минуя имя поля, которое мы опустили
	Email    string
	UserName string
	Password string
}

// также можно встаивать и указатель на структуру, тогда в литерале
// перед заданием этой структуры необ поставить &
type Circle struct {
	*Point2d
	Radius int
}

func DemoEmbeddedStruct() {
	var admin User
	admin.Name = "Oleg"   // обращаемся к встроенной структуре через точку
	admin.Person.Age = 32 // а можно так же как если бы мы указали имя поля
	admin.UserName = "DES2048"

	fmt.Println(admin)

	// однако данный сахар не работает при использовании стрктуррных литералов
	// нужно указывать имена втроенных полей
	user := User{Person{"Adrew", 31}, "aaa@mail.com", "andy", "111"}
	fmt.Println(user)

	user2 := User{
		Person: Person{ // нужно указать имя типа структуры, а потом имя поля, так можно исп имена полей
			Name: "Oleg",
			Age:  31},
		UserName: "Oleg123", // нужно завершать запятной, если исп литералы с указанием полей
	}
	fmt.Println(user2)

	// если встриваем указатель на структуру то просто нужно поствить & перед опр встр структуры
	var circ = Circle{
		&Point2d{1, 1},
		4,
	}

	fmt.Printf("%#v\n", circ)
}
