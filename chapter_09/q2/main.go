package main

func main() {
    ch1 := make(chan int, 1)

    ch1 <- 1
    ch1 <- 2 // 通道已满，会造成阻塞

    ch2 := make(chan int, 1)
    aaa := <-ch2//通道已空，会造成阻塞
    ch2 <- 1

    var ch3 chan int
    ch3 <- 1 //通道未初始化，nil, 造成永久阻塞
    <- ch3 // 通道未初始化， nil，造成永久阻塞
    _ = ch3
}
