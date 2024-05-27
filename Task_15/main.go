package main

import (
	"fmt"
	"sync"
	"sync/atomic"
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
	wg := sync.WaitGroup{}
	for i := 0; i < 5; i++ {
		i := i
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println("Горутина", i)
		}()
	}
	wg.Wait()
	fmt.Println("Завершение работы программы")
}

func task2() {
	var j uint32
	wg := sync.WaitGroup{}
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			atomic.AddUint32(&j, 1)
		}()
	}
	wg.Wait()
	fmt.Println(j)
}

func task3() {
	var n int
	var m sync.Mutex
	for i := 0; i < 50; i++ {
		go func() {
			m.Lock()
			defer m.Unlock()
			n++
		}()
	}
	time.Sleep(3 * time.Second)
	m.Lock()
	defer m.Unlock()
	fmt.Println(n)
}

func task4() {
	var loadOnce sync.Once
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		i := i
		wg.Add(1)
		go func() {

			defer wg.Done()
			fmt.Println("Run goroutine", i)
			loadOnce.Do(start)
		}()
	}
	wg.Wait()
}

func start() {
	fmt.Println("Run func start")
}

type Meteo interface {
	ReadTemp() string
	ChangeTemp(v string)
}

type temperature struct {
	mu   sync.RWMutex
	temp int16
}

func (t *temperature) ReadTemp() int16 {
	t.mu.RLock()
	defer t.mu.RUnlock()
	return t.temp
}

func (t *temperature) ChangeTemp(v int16) {
	t.mu.Lock()
	defer t.mu.Unlock()
	t.temp = v
}

func task5() {
	wg := sync.WaitGroup{}
	tp := temperature{
		temp: 16,
	}

	for i := 0; i < 50; i++ {
		i := i
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Printf("Читает горутина %d - Текущая температура %d\n", i, tp.ReadTemp())
		}()
		if i != 0 && i%5 == 0 {
			wg.Add(1)
			go func() {
				defer wg.Done()
				fmt.Println("Изменяем температуру")
				tp.ChangeTemp(int16(i))
			}()
		}
	}
	wg.Wait()
}
