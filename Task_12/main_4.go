package main

import (
	"errors"
	"fmt"
	"reflect"
)

/*
Паника возникает при вызове b.Sing(), т.к. интерфесу в данном случае присвоен тип указателя на Duck, но без значения.
В результате при отсутствии значения (объекта) вызов его метода невозможен.
*/

/*
//Вариант решения 1
//Перменной d присвоить конкретное значение структуры
//var d *Duck = &Duck{voice: "кря"}
//В этом случае у интерфеса будет и тип и значение, т.е. вызов b.Sing() пройдет без проблем

type Bird interface {
	Sing() string
}

type Duck struct {
	voice string
}

func (d *Duck) Sing() string {
	return d.voice
}

func main() {
	var d *Duck = &Duck{voice: "кря"}
	song, err := Sing(d)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(song)
}

func Sing(b Bird) (string, error) {
	if b != nil {
		return b.Sing(), nil
	}
	return "", errors.New("Ошибка пения!")
}

*/

/*
//Вариант решения 2
//Отказаться от ссылочного типа, тогда переменной присвоится значение по умолчанию

type Bird interface {
	Sing() string
}

type Duck struct {
	voice string
}

func (d Duck) Sing() string {
	return d.voice
}

func main() {
	var d Duck
	song, err := Sing(d)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(song)
}

func Sing(b Bird) (string, error) {
	if b != nil {
		return b.Sing(), nil
	}
	return "", errors.New("Ошибка пения!")
}
*/

//Вариант решения 3
//В функции Sing добавить условие проверки на преобразование к типу Duck (b.(*Duck) != nil) и только в случае с не nil вызывать метод Sing.
//или
//Использовать пакет reflect и метод проверяющий значение на nil (!reflect.ValueOf(b).IsNil())

type Bird interface {
	Sing() string
}

type Duck struct {
	voice string
}

func (d *Duck) Sing() string {
	return d.voice
}

func main() {
	var d *Duck
	song, err := Sing(d)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(song)
}

func Sing(b Bird) (string, error) {
	//if b != nil && b.(*Duck) != nil {
	//	return b.Sing(), nil
	//}
	if !reflect.ValueOf(b).IsNil() {
		return b.Sing(), nil
	}
	return "", errors.New("Ошибка пения!")
}
