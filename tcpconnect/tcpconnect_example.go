
package main

import (
    "net"
    "time"
)

func main() {
    for {
        conn, _ := net.DialTimeout("tcp", "127.0.0.1:3306", time.Second)
        if conn != nil {
            conn.Close()
        }
        time.Sleep(10 * time.Millisecond)
    }
}
