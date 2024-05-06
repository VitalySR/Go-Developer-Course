package main

import (
	"errors"
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
}

func task1() {
	err := errors.New("ошибка1")

	err = fmt.Errorf("ошибка2:%w", err)

	err = fmt.Errorf("ошибка3:%w", err)

	fmt.Println(err)
}

func task2() {
	err := errors.New("ошибка1")

	err = fmt.Errorf("ошибка2:%w", err)

	err = fmt.Errorf("ошибка3:%w", err)

	fmt.Printf("Сформированная ошибка -> %v\n", err)

	err1 := errors.Unwrap(err)

	fmt.Printf("Ошибка до -> %v\n", err1)
}

func task3() {
	err := errors.New("ошибка1")

	err1 := fmt.Errorf("ошибка2:%w", err)

	err2 := fmt.Errorf("ошибка3:%w", err1)

	fmt.Printf("Сформированная ошибка -> %v\n", err2)

	fmt.Println("Ошибка \"ошибка1\" была:", errors.Is(err2, err))
}

type myFirstError struct {
	msg string
}

func (m myFirstError) Error() string {
	return m.msg
}

func task4() {
	err := errors.New("ошибка1")

	err = fmt.Errorf("ошибка2:%w", err)

	err = fmt.Errorf("ошибка3:%w", err)

	fmt.Printf("Сформированная ошибка -> %v\n", err)

	fmt.Println("Ошибка типа myFirstError была:", errors.As(err, new(myFirstError)))
}
