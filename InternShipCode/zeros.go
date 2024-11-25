package main

import (
    "fmt"
    "strings"
)

func main() {
    var userInput int
    fmt.Print("Enter the number of zeros you want: ")
    _, err := fmt.Scan(&userInput)
    
    if err != nil {
        fmt.Println("Invalid input. Please enter a valid number.")
        return
    }

    if userInput < 0 {
        fmt.Println("Number of zeros cannot be negative.")
        return
    }

    zeros := strings.Repeat("0", userInput)
    fmt.Println("Resulting string of zeros:", zeros)
}
