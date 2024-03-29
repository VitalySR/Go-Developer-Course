package main

import (
	"fmt"
)

func main() {
	//Задание 1
	task1()

	//Задание 2
	task2()
}

func task1() {
	var symbol rune
	var bt byte

	fmt.Println("Значение по умолчанию:")
	fmt.Printf("Для rune - %v (тип %T)\n", symbol, symbol)
	fmt.Printf("Для byte - %v (тип %T)\n", bt, bt)
	fmt.Println()
	/*
		rune - это алиас для типа int32, а byte - это алиас для типа uint8.
		Так как это числовые типы, то для них значение по умолчанию = 0
	*/
}

func task2() {
	fmt.Println("Деление 16 на 3:")
	rslt1 := 16 / 3.0
	rslt2 := 16 % 3
	fmt.Printf("Результат:%v, остаток от деления: %v, тип результата: %T", rslt1, rslt2, rslt1)
}
