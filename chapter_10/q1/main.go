package main

import (
    "fmt"
    "time"
    "math/rand"
)

func main() {
    intChannels := [3]chan int {
        make(chan int, 1),
        make(chan int, 1),
        make(chan int, 1),
    }

    rand.Seed(int64(time.Now().UnixNano()))
    index := rand.Intn(3)
    fmt.Printf("The index: %d\n", index)
    intChannels[index] <- index * 800

    select {
    case <- intChannels[0]:
        fmt.Println("The first candidate case is selected.")
    case <- intChannels[1]:
        fmt.Println("The second candidate case is selected.")
    case <- intChannels[2]:
        fmt.Println("The third candidate case is selected.")
    default:
        fmt.Println("No candidate case is selected.")
    }
}
