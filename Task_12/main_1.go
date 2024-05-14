package main

import (
	"fmt"
	"log"
)

func main() {
	a := 1
	do(a)
}

func do(v any) {
	a := 10
	/*
	 Здесь нужно увеличить значение a на v.
	 В случае невозможности приведения к int
	 необходимо сообщить об этом и немедленно завершить
	 выполнение программы.
	*/
	vi, ok := v.(int)
	if !ok {
		//fmt.Println("v is not int")
		log.Fatalln("v is not int")
	}

	a += vi

	fmt.Println(a)
}
