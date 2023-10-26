package main

import (
"fmt"
"sync"
)

func generateNumbers(total int, wg *sync.WaitGroup, ch chan int) {
    defer wg.Done()

    sum := 0
    for idx := 1; idx <= total; idx++ {
        fmt.Printf("Generating number %d\n", idx)
        sum = sum + idx
        ch <- sum
    }
}

func printNumbers(wg *sync.WaitGroup, ch chan int) {
    defer wg.Done()

    fmt.Printf("Sum is now:\n")
    for idx := 1; idx <= 3; idx++ {
        sum := <-ch
        fmt.Printf("Printing number %d %d\n", idx, sum)
    }
}

func main() {
    var wg sync.WaitGroup

    ch1 := make(chan int)

    wg.Add(2)
    go printNumbers(&wg, ch1)
    go generateNumbers(3, &wg, ch1)

    fmt.Println("Waiting for goroutines to finish...")
    wg.Wait()
    fmt.Println("Done!")
}

