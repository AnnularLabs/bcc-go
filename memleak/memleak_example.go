
package main

import "time"

var leaks [][]byte

func main() {
    for {
        b := make([]byte, 1024*100) // 100KB
        leaks = append(leaks, b)
        time.Sleep(100 * time.Millisecond)
    }
}
