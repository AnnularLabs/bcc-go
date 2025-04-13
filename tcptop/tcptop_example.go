
package main

import (
    "net"
    "time"
)

func main() {
    for {
        conn, err := net.Dial("tcp", "google.com:80")
        if err == nil {
            conn.Write([]byte("GET / HTTP/1.0\r\n\r\n"))
            time.Sleep(100 * time.Millisecond)
            conn.Close()
        }
    }
}
