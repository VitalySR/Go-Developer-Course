package main

import (
	"fmt"
)

func main() {
	// task1()

	// task2()

	// task3()

	// task4()

	// task5()

	task6()
}

/*------------------------------------------------------*/

func task1() {
	hello1()
}

func hello1() {
	fmt.Println("Task1. Hello, Go!")
}

/*------------------------------------------------------*/

func task2() {
	func() {
		fmt.Println("Task2. Hello, Go!")
	}()
}

/*------------------------------------------------------*/

func task3() {
	afn := func() {
		fmt.Println("Task3. Hello, Go!")
	}
	hello3(afn)
}

func hello3(fn func()) {
	fn()
}

/*------------------------------------------------------*/

func task4() {
	hello4()()
}

func hello4() func() {
	return func() {
		fmt.Println("Task4. Hello, Go!")
	}
}

/*------------------------------------------------------*/

func task5() {
	defer fmt.Println("Task5. Завершение работы")
	hello5()
}

func hello5() {
	fmt.Println("Task5. Hello, Go!")
}

/*------------------------------------------------------*/

var cnt uint8

func task6() {
	//fibonacci(0, 1)
	fibonacci2(23)
}

func fibonacci(one int, two int) {
	if cnt >= 23 {
		return
	}
	cnt++
	fmt.Printf("%d ", one)
	fibonacci(two, one+two)
}

func fibonacci2(n int) (int, int) {
	if n == 0 {
		return 0, 1
	}
	one, two := fibonacci2(n - 1)
	fmt.Printf("%d ", one)
	return two, one + two
}
