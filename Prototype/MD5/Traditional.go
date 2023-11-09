package main

import (
    "crypto/md5"
    "encoding/hex"
    "fmt"
    "io/ioutil"
    "os"
    "strings"
)

const dataFileName = "data.txt"

func calculateMD5(input string) string {
    hasher := md5.New()
    hasher.Write([]byte(input))
    return hex.EncodeToString(hasher.Sum(nil))
}

func saveData(input, hash string) {
    f, err := os.OpenFile(dataFileName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666)
    if err != nil {
        fmt.Println("Error opening file:", err)
        return
    }
    defer f.Close()

    _, err = fmt.Fprintf(f, "%s,%s\n", input, hash)
    if err != nil {
        fmt.Println("Error writing to file:", err)
    }
}

func main() {
    if _, err := os.Stat(dataFileName); os.IsNotExist(err) {
        file, err := os.Create(dataFileName)
        if err != nil {
            fmt.Println("Error creating data file:", err)
            return
        }
        file.Close()
    }

    data, err := ioutil.ReadFile(dataFileName)
    if err != nil {
        fmt.Println("Error reading data file:", err)
        return
    }

    existingData := string(data)
    lines := strings.Split(existingData, "\n")

    collisionFound := false
    lastProcessedInput := "input0"

    for _, line := range lines {
        parts := strings.Split(line, ",")
        if len(parts) == 2 {
            lastProcessedInput = parts[0]
            if parts[1] == "collision" {
                collisionFound = true
                break
            }
        }
    }

    lastProcessedIndex := 0
    fmt.Sscanf(lastProcessedInput, "input%d", &lastProcessedIndex)

    for i := lastProcessedIndex + 1; !collisionFound; i++ {
        input := fmt.Sprintf("input%d", i)
        hash := calculateMD5(input)

        for _, line := range lines {
            parts := strings.Split(line, ",")
            if len(parts) == 2 {
                if parts[1] == hash {
                    collisionFound = true
                    fmt.Printf("Collision Found: %s -> %s\n", input, hash)
                    saveData(input, "collision")
                    break
                }
            }    
        }

        if !collisionFound {
            fmt.Printf("No Collision Found: %s -> %s\n", input, hash)
            saveData(input, hash)
        }
    }
}
