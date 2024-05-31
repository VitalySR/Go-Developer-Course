package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	//fmt.Println("-- Task1")
	//task1()
	//fmt.Println("-- Task2")
	//task2()
	//fmt.Println("-- Task3")
	//task3()
	//fmt.Println("-- Task4")
	//task4()
	fmt.Println("-- Task5")
	task5()
}

func task1() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	wg := sync.WaitGroup{}
	for i := 0; i < 5; i++ {
		i := i
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				select {
				case <-ctx.Done():
					fmt.Println("Завершение работы горутины", i)
					return
				default:
					time.Sleep(1 * time.Second)
					fmt.Println("Работает горутина", i)
				}
			}
		}()
	}
	time.Sleep(2 * time.Second)
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("Send cancel")
		cancel()
	}()
	wg.Wait()
}

func task2() {
	start := time.Now()
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()
	wg := sync.WaitGroup{}
	for i := 0; i < 5; i++ {
		i := i
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				select {
				case <-ctx.Done():
					fmt.Println("Завершение работы горутины", i)
					return
				default:
					time.Sleep(1 * time.Second)
					fmt.Println("Работает горутина", i)
				}
			}
		}()
	}

	wg.Wait()
	duration := time.Since(start)
	fmt.Println("Время выполнения:", duration)
}

func task3() {
	start := time.Now()
	ctx := context.Background()
	dlt := time.Now().Add(2 * time.Second)
	ctx, cancel := context.WithDeadline(ctx, dlt)
	defer cancel()
	wg := sync.WaitGroup{}
	for i := 0; i < 5; i++ {
		i := i
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				select {
				case <-ctx.Done():
					fmt.Println("Завершение работы горутины", i)
					return
				default:
					time.Sleep(1 * time.Second)
					fmt.Println("Работает горутина", i)
				}
			}
		}()
	}

	wg.Wait()
	duration := time.Since(start)
	fmt.Println("Время выполнения:", duration)
}

type ctxKay string

func task4() {
	ctx := context.Background()
	var key ctxKay = "some key1"
	ctx = context.WithValue(ctx, key, "some value1")
	key = "some key2"
	ctx = context.WithValue(ctx, key, "some value2")
	task4print(ctx)
}

func task4print(ctx context.Context) {
	var key ctxKay = "some key1"
	fmt.Println(key, "=", ctx.Value(key))
	key = "some key2"
	fmt.Println(key, "=", ctx.Value(key))
}

func task5() {
	mu := sync.Mutex{}
	stop := make(chan struct{})
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		i := i
		wg.Add(1)
		go func() {
			defer wg.Done()
			defer mu.Unlock()
			for {
				mu.Lock()
				fmt.Println("сложные вычисления горутины:", i)
				select {
				case <-stop:
					fmt.Println("stop горутина:", i)
					return
				default:
					mu.Unlock()
					time.Sleep(1 * time.Second)
				}
			}
		}()
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		time.Sleep(3 * time.Second)
		fmt.Println("ой, всё!")
		close(stop)
	}()

	wg.Wait()
}
