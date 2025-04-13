
package main

import (
    "fmt"
)

func busyLoop() {
    sum := 0
    for i := 0; i < 1e9; i++ {
        sum += i
    }
    fmt.Println("Done:", sum)
}

func main() {
    for {
        busyLoop()
    }
}
