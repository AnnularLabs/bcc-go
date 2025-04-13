
package main

import (
    "fmt"
    "time"
)

func PlaceOrder(userID int) {
    if userID%3 == 0 {
        time.Sleep(2 * time.Millisecond)
    } else {
        time.Sleep(500 * time.Microsecond)
    }
    fmt.Printf("User %d placed order\n", userID)
}

func main() {
    for i := 0; i < 10000; i++ {
        PlaceOrder(i)
        time.Sleep(10 * time.Millisecond)
    }
}
