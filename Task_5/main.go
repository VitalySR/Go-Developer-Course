package main

import (
	"fmt"
)

func main() {
	fmt.Println("-- Task1")
	task1()
	fmt.Println("-- Task2")
	task2()
	fmt.Println("-- Task3")
	task3()
	fmt.Println("-- Task4")
	task4()
	fmt.Println("-- Task5")
	task5()
	fmt.Println("-- Task6")
	task6()
	fmt.Println("-- Task7")
	task7()
	fmt.Println("-- Task8")
	task8()
}

func task1() {
	var st *string
	pr := "Строка"

	st = &pr

	fmt.Println(st, pr, *st)
}

func task2() {
	n := 15

	fmt.Println(n, &n)
}

func task3() {
	var v *int
	b := 22
	v = &b
	fmt.Println(b, v)
	*v = 11
	fmt.Println(b, v)
}

func task4() {
	i, j, x := 10, 15, 20

	fmt.Printf("i = %d, %v\n", i, &i)
	fmt.Printf("j = %d, %v\n", j, &j)
	fmt.Printf("x = %d, %v\n", x, &x)
}

func task5() {
	r := 33
	fmt.Println(r)
	task5Change(&r)
	fmt.Println(r)
}

func task5Change(d *int) {
	*d = 44
}

type square int

func task6() {
	var s square = 25
	fmt.Println(s)
}

func task7() {
	var s square = 30
	fmt.Println(s)
	s += 15
	fmt.Println(s)
}

func task8() {
	var s square = 34
	s += 10
	fmt.Println(fmtSquare(s))
}

func fmtSquare(v square) string {
	return fmt.Sprintf("%d м²", v)
}

/*
// Square - пользовательский тип, обертка для int
type Square struct {
	area int // площадь квадрата
}

// NewSquare - функция для создания нового экземпляра Square
func NewSquare(area int) *Square {
	return &Square{area: area}
}

// String - метод для форматированного вывода значения Square
func (s *Square) String() string {
	return fmt.Sprintf("%d м²", s.area)
}

func task9() {
	// Создаем новый экземпляр Square с площадью 34 м²
	s := NewSquare(34)

	// Увеличиваем площадь на 10 м²
	s.area += 10

	// Выводим значение Square в консоль
	fmt.Println(s)
}
*/
