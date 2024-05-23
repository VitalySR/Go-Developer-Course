package main

import (
	"fmt"
	"time"
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
	//fmt.Println("-- Task6")
	//task6()
}

func task1() {
	go func() {
		fmt.Println("Привет из дочерней горутины!")
	}()

	time.Sleep(1 * time.Second)
	fmt.Println("Основная горутина")
}

func task2() {
	cn := make(chan string)
	go func() {
		cn <- "Привет, строковый канал!"
		fmt.Println("Дочерняя горутина")
	}()

	fmt.Println(<-cn)
}

func task3() {
	cn := make(chan string, 4)
	cn <- "Привет"
	cn <- "буферизованный канал!"
	close(cn)
	for v := range cn {
		fmt.Println(v)
	}
}

func task4() {
	ch := make(chan int)
	stop := make(chan struct{})
	close(ch) //solution
	go func() {
		<-ch
		stop <- struct{}{}
	}()
	<-stop
	fmt.Println("happy end")
}

func task5() {
	//ch := make(chan int)
	var ch chan int //solution
	stop := make(chan struct{}, 2)
	go func() {
	OUT:
		for {
			select {
			case <-stop:
				break OUT
			case v, ok := <-ch:
				if !ok {
					break OUT
				}
				fmt.Println(v)
			default:
				continue
			}
		}
		fmt.Println("завершение работы горутины_1")
	}()
	go func() {
		var i int
	OUT:
		for {
			i++
			select {
			case <-stop:
				break OUT
			default:
				time.Sleep(time.Second)
				if ch == nil {
					continue
				}
				ch <- i
			}
		}
		fmt.Println("завершение работы горутины_2")
	}()
	time.Sleep(5 * time.Second)
	stop <- struct{}{}
	stop <- struct{}{}
	time.Sleep(time.Second)
	fmt.Println("завершение работы главной горутины")
}

func task6() {
	ch := make(chan int)
	stop := make(chan struct{}, 2)
	go func() {
	OUT:
		for {
			select {
			case <-stop:
				break OUT
			case v, ok := <-ch:
				if !ok {
					break OUT
				}
				fmt.Println(v)
			default:
				continue
			}
		}
		fmt.Println("завершение работы горутины_1")
	}()
	go func() {
		var i int
	OUT:
		for {
			i++
			select {
			case <-stop:
				break OUT
			default:
				time.Sleep(time.Second)
				if ch == nil {
					continue
				}
				ch <- i
			}
		}
		fmt.Println("завершение работы горутины_2")
	}()
	time.Sleep(5 * time.Second)
	stop <- struct{}{}
	stop <- struct{}{}
	time.Sleep(time.Second)
	fmt.Println("завершение работы главной горутины")
}
