package main

import (
	"fmt"
)

const (
	cnt0 = 15
	cnt1 = 5
	cnt2 = 2.3
	cnt3 = 'd'
	cnt4 = "string"
)

func main() {
	task1()

	task2()

	task3()

	task4()

	task5()

	task6()

	task7()
}

func task1() {
	fmt.Println(cnt0)
}

func task2() {
	const vrn = "Значение локальное"

	fmt.Println(vrn)
}

func task3() {
	const cnt1 = 10

	fmt.Println(cnt1)
}

func task4() {
	fmt.Println(cnt0, cnt1, cnt2, cnt3, cnt4)
}

func task5() {
	const (
		loc0 = 0
		loc1
		loc2 = 1
		loc3
		loc4 = 2
	)

	fmt.Println(loc0, loc1, loc2, loc3, loc4)
}

func task6() {
	const n int = 5
	m := 3.4 + float32(n)
	fmt.Println(m)
}

func task7() {
	const (
		loc0 = 1 << iota
		loc1
		loc2
		loc3
		loc4
	)
	fmt.Println(loc0, loc1, loc2, loc3, loc4)
}
