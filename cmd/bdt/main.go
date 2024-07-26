package main

import (
    "fmt"
    "os"

    "github.com/derickit/blockchain-dev-tools/pkg/utils"
)

func main() {
    if len(os.Args) < 2 {
        fmt.Println("Usage: bdt <command> [arguments]")
        utils.PrintVersion()
        os.Exit(1)
    }

    command := os.Args[1]

    switch command {
    case "version":
        utils.PrintVersion()
    case "hex":
        if len(os.Args) < 3 {
            fmt.Println("Usage: bdt hex <number>")
            os.Exit(1)
        }
        result, err := utils.ToHex(os.Args[2])
        if err != nil {
            fmt.Printf("Error: %v\n", err)
            os.Exit(1)
        }
        fmt.Println(result)
    case "fromhex":
        if len(os.Args) < 3 {
            fmt.Println("Usage: bdt fromhex <hex_value>")
            os.Exit(1)
        }
        result, err := utils.FromHex(os.Args[2])
        if err != nil {
            fmt.Printf("Error: %v\n", err)
            os.Exit(1)
        }
        fmt.Println(result)
    case "hash":
        if len(os.Args) < 3 {
            fmt.Println("Usage: bdt hash <string>")
            os.Exit(1)
        }
        result := utils.Hash(os.Args[2])
        fmt.Println(result)
    case "ping":
        if len(os.Args) < 3 {
            fmt.Println("Usage: bdt ping <address>")
            os.Exit(1)
        }
        result := utils.Ping(os.Args[2])
        fmt.Println(result)
    default:
        fmt.Printf("Unknown command: %s\n", command)
        os.Exit(1)
    }
}