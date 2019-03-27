package core

import "fmt"

func SlicesDemo() {
	// срезы это специальная структура над массивами, которая позволяет им работать с массивами любых
	// размеров указанного типа. можно сказать что срезы это обобщения над массивами своего рода.
	// внутренне срез состоит из трех компонентов: указателя на базовый массив, длины среза и емкости
	// базового массива

	// объявлеям не привязанный срез, он будет равен nil
	var str_slice []string
	fmt.Println("unbound slice is", str_slice, "is nil?", str_slice == nil)

	// но если с пустым литералом то уже срез не будет равен nill
	str_empty := []string{} // его len и cap будут равно 0
	fmt.Println("bound to empty", str_empty, "is nil?", str_empty == nil)

	// поэтому лучший способ проверки это len() !=0 или комб проверка через ИЛИ

	// len(nil) ведет себ корректно, поэтому спокойно можно передавать nil срез в нее
	fmt.Println("len of unbound slice:", len(str_slice))

	months := [...]string{
		1:  "January",
		2:  "Febrary",
		3:  "March",
		4:  "April",
		5:  "May",
		6:  "June",
		7:  "July",
		8:  "August",
		9:  "September",
		10: "October",
		11: "November",
		12: "December",
	}

	// осуществляем привязку массива к срезу
	str_slice = months[:]
	// можно юзать срез для работы с базовым массивом
	fmt.Println(str_slice)

	// срез задается как в python
	firstQuarter := months[1:5] // срез от 1 элемента до 4
	// длина среза говорит нам сколько у нас в срезе элементов
	// емкость сколько от 1-го элемента среза до конца массива,
	fmt.Println(firstQuarter, "len():", len(firstQuarter), "cap():", cap(firstQuarter))

	// на один и тот же массив можно накладывать сколько угодно срезов
	// причем они могут пересекаться
	spring := months[3:6]
	fmt.Println("spring:", spring)

	// срезы можно расширять до cap()
	// end здесь расчитывается от первого индеса в срезе
	firstHalfYear := firstQuarter[:7] // расширим срез до июля
	fmt.Println(firstHalfYear)

	// сжимать
	winter := firstQuarter[:3]
	fmt.Println(winter)

	//!!! срезы нельзя сравнивать между собой, единственное это можно проверить срез на nil

	// таже срез можно создать с помощью литерала, ОПУСТИВ длину массива
	// под капотом создается массив, к кторому можно получить доступ с помощью данного среза
	langs := []string{"Go", "Java", "C", "Python"}
	fmt.Printf("%v type:%T\n", langs, langs)

	fmt.Println()
	fmt.Println("make()")

	// также срез можно сделать с помощью функици make()
	// создает массив из 10 элементов и делает его доступным через срез cap и len = 10
	results := make([]int, 10)
	// выведет массив заполненный нулями
	fmt.Println(results, "len():", len(results), "cap()", cap(results))

	some := make([]int, 5, 10) // len =5 cap =10 эквивалентно make([int],10)[:5]
	fmt.Println(some, "len():", len(some), "cap():", cap(some))

}

func SliceAppendDemo() {
	// добавлять элементы к массиву можно с помощью функции append() - append ведет себя как на структуре
	// ArrayList - т.е если размера результ масиива недостаточно, то юудет воздан и возвращен новый массив

	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	nums = nums[0:6]
	fmt.Println("nums[0:6]:", nums, "len():", len(nums), "cap():", cap(nums))

	//!!! очень интересное поведение функции append, если сделать срез длиной меньше чем cap
	// а потом сделать append() - произойдет перезапись элементов между len() и cap()
	nums = append(nums, 11, 12)
	fmt.Println("nums[0:6] after append 11,12:", nums, "len():", len(nums), "cap():", cap(nums))
	fmt.Println(" nums[:]:", nums[:10]) // это приведет к замене элементов с индексами 6,7

	// с доступным базовым массивом ситуация еще интересней
	nums_arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	int_slice := nums_arr[:8]

	fmt.Println("append with explicit basic array")

	fmt.Println("adding 1")
	int_slice = append(int_slice, 1)

	fmt.Println("slice:", int_slice, "array:", nums_arr)

	fmt.Println("adding 2")
	int_slice = append(int_slice, 2)
	fmt.Println("slice:", int_slice, "array:", nums_arr)

	//!!! здесь будет заметно, что массив отсоединен от слайса, так как
	// к слайсу теперь привязан новый массив
	fmt.Println("adding 2")
	int_slice = append(int_slice, 3)
	fmt.Println("slice:", int_slice, "array:", nums_arr)

	fmt.Println("slice[0] =100")
	int_slice[0] = 100
	fmt.Println("slice:", int_slice, "array:", nums_arr)

	// вобщем лучше создавать и иниц слайсы чем создавать массивы, а потом привязывать их к слайсу
}

func SlicePassToFuncDemo() {
	// !!!нужно быть осторожным когда передаешь срез в функцию, так как срез это не указатель в чистом виде
	// а некая структура и передается что-то вроде алиаса, с которым можно работать in place, но если
	// расширить или реалоцировать базовый массив по этому срезу, то срез будет модифицирован только
	// внутри функции. А снаружи останется таким же, в этом случае возвращают новый срез из функции и
	// присваивают его срезу. ЭТО СТАНДАРТНАЯ ПРАКТИКА, когда ты передаешь срез в функцию которая может
	// реалоцировать его базовый массив

	// эта функция умножает переданный массив на 2 in place
	inplaceMul2 := func(list []int) []int {
		// здесь модифицируется сам массив, а указатель на него остается тем же,
		// поэтому такие изменения отразятся и на переданном массиве
		for i := 0; i < len(list); i++ {
			list[i] *= 2
		}

		return list
	}

	nums := []int{1, 2, 3, 4, 5}

	fmt.Println(nums, "mul2 in place:")
	inplaceMul2(nums)
	fmt.Println(nums)

	// эта функция умножает переданный массив на 2 c realoc
	Mul2 := func(list []int) []int {
		// создаем и копируем значения в новый массив. list извне не изм своего указателя
		new_list := make([]int, len(list))
		copy(new_list, list)
		list = new_list

		for i := 0; i < len(list); i++ {
			list[i] *= 2
		}

		return list
	}

	fmt.Println("nums before mul2:", nums)
	Mul2(nums) // не присвоили новый срез, поэтому он останется прежним
	fmt.Println("after:", nums)
	fmt.Println("with reassignment:", Mul2(nums)) // все норм так как возвр новый срез
}
