
package main

import (
    "net"
    "time"
)

func main() {
    for {
        conn, _ := net.Dial("tcp", "192.0.2.1:12345") // 不存在的地址模拟丢包
        if conn != nil {
            conn.Close()
        }
        time.Sleep(10 * time.Millisecond)
    }
}
