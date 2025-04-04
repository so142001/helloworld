package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    // Create a buffered reader to read from standard input
    reader := bufio.NewReader(os.Stdin)

    // Prompt the user to press Enter
    fmt.Println("Press Enter to see what happens!")

    // Wait for the user to press Enter
    _, err := reader.ReadString('\n')
    if err != nil {
        fmt.Println("Error reading input:", err)
        return
    }

    // Print the greeting message
    fmt.Println("Hello, World!")
}
