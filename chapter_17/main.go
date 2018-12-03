package main

import (
	"fmt"
)

func main() {
	num := 10

	sign := make(chan struct{}, num)
	for i := 0; i < num; i++ {
		go func() {
			fmt.Println(i)
			sign <- struct{}{}
		}()
	}

	for j := 0; j < num; j++ {
		<-sign
	}

	//time.Sleep(time.Millisecond * 500)
}

//10
//10
//10
//10
//10
//5
//10
//10
//10
//10

// 这样只能保证进程退出之前能让所有的goroutine执行完，但还是不能保证执行顺序（其实本质上也不能控制goroutine的执行顺序）只是可以传入参数控制goroutine的输出
