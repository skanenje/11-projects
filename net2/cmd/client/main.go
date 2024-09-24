package main

import (
    "fmt"
    "os"
    "net/internal/client"
)

func main() {
    if len(os.Args) < 3 {
        fmt.Println("[USAGE]: ./TCPChat $port $IP")
        return
    }

    port := os.Args[1]
    ip := os.Args[2]

    client := client.NewClient(ip, port)
    if err := client.Run(); err != nil {
        fmt.Println("Error:", err)
    }
}