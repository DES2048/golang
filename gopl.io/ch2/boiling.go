package ch2

import "fmt"

const boilingF = 212. // приватная константа уровня пакета
const freezeF  = 32.

func PrintBoilingPoints()  {
	var boilingC = FtoC(boilingF)

	fmt.Printf("Boiling point: %gF or %gC\n", boilingF, boilingC)

}

func PrintFreezePoint() {
	var freezeC = FtoC(freezeF)

	fmt.Printf("Freeze point: %gF or %gC\n", freezeF, freezeC)
}
