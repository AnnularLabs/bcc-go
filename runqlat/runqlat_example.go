
package main

import (
    "time"
    "runtime"
)

func sleeper() {
    for {
        time.Sleep(10 * time.Millisecond)
    }
}

func main() {
    runtime.GOMAXPROCS(1)
    for i := 0; i < 1000; i++ {
        go sleeper()
    }
    select {}
}
