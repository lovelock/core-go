package main

import "fmt"

func main() {
    ch := make(chan int, 3)

    // ch <- 1
    ch <- 2
    ch <- 3
    ch <- 4

    elem1 := <-ch

    fmt.Printf("The first element received from channel ch1: %v\n", elem1)
}
