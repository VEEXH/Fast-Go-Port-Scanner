package main

import (
    "fmt"
    "net"
    "os"
    "sync"
)

func main() {
    if len(os.Args) != 2 {
        fmt.Println("Usage: go run main.go <ip>")
        return
    }
    ip := os.Args[1]

    var wg sync.WaitGroup

    // Set up the animation
    animation := []string{"|", "/", "-", "\\"}
    animIndex := 0

    for port := 1; port <= 65535; port++ {
        wg.Add(1)
        go func(p int) {
            defer wg.Done()
            address := fmt.Sprintf("%s:%d", ip, p)
            conn, err := net.Dial("tcp", address)
            if err == nil {
                fmt.Printf("Port %d open\n", p)
                conn.Close()
            }
        }(port)

        // Update the animation every 1000 ports
        if port % 1000 == 0 {
            fmt.Printf("\rScanning... %s", animation[animIndex])
            animIndex = (animIndex + 1) % len(animation)
        }
    }

    wg.Wait()
    fmt.Println("\rScan complete")
}
