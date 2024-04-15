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
	var m = [5]string{"один", "два", "три", "четыре", "пять"}
	fmt.Println(m)
}

func task2() {
	m := [...]string{"яблоко", "груша", "слива", "абрикос"}
	fmt.Println(m)
}

func task3() {
	m := [4]string{"яблоко", "груша", "помидор", "абрикос"}
	fmt.Println(m)
	m[2] = "персик"
	fmt.Println(m)
}

func task4() {
	var n = []int{5, 2, 8, 3, 1, 9}
	fmt.Println(n)
}

func task5() {
	n := make([]int, 0, 10)
	fmt.Printf("Slice n %v, length %d, capacity %d\n", n, len(n), cap(n))
}

func task6() {
	n := make([]int, 0, 10)
	fmt.Printf("Slice n %v, length %d, capacity %d\n", n, len(n), cap(n))
	n = append(n, 4, 1, 8, 9)
	fmt.Printf("Slice n %v, length %d, capacity %d\n", n, len(n), cap(n))
}

func task7() {
	s1 := []int{1, 2, 3}
	s2 := []int{4, 5, 6}
	fmt.Printf("Slice s1 %v, length %d, capacity %d\n", s1, len(s1), cap(s1))
	fmt.Printf("Slice s2 %v, length %d, capacity %d\n", s2, len(s2), cap(s2))
	s1 = append(s1, s2...)
	fmt.Printf("Slice s1 %v, length %d, capacity %d\n", s1, len(s1), cap(s1))
	fmt.Printf("Slice s2 %v, length %d, capacity %d\n", s2, len(s2), cap(s2))
}

func task8() {
	n := []int{1, 2, 3, 4, 5, 6}
	fmt.Printf("Slice n %v, length %d, capacity %d\n", n, len(n), cap(n))
	n = append(n[:3], n[4:]...)
	fmt.Printf("Slice n %v, length %d, capacity %d\n", n, len(n), cap(n))
}
