package main

import (
	"fmt"

	v100 "github.com/VitalySR/task10_mod_v1.0.0"
	v110 "github.com/VitalySR/task10_mod_v1.1.0"
)

func main() {
	fmt.Println("Проверка использования модуля")
	v100.Hello()
	v110.Hello()
}
