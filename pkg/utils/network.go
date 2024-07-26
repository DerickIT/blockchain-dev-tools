package utils

import (
    "fmt"
    "net"
    "time"
)

func Ping(address string) string {
    conn, err := net.DialTimeout("tcp", address, time.Second*5)
    if err != nil {
        return fmt.Sprintf("Failed to connect to %s: %v", address, err)
    }
    defer conn.Close()
    return fmt.Sprintf("Successfully connected to %s", address)
}