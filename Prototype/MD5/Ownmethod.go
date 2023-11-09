package main

import (
	"bufio"
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter initial input: ")
	input, _ := reader.ReadString('\n')
	input = input[:len(input)-1] // Remove newline character

	// Save the initial input to a file
	err := appendFile("input.txt", input)
	if err != nil {
		fmt.Println("Error saving input to file:", err)
		return
	}

	hash := md5.Sum([]byte(input))
	firstHash := hash
	fmt.Printf("Initial Hash: %x\n", hash)

	// Initialize the file to store hashes
	hashFile := "hashes.txt"
	hashCounter := 0

	// Check if hashes file already exists
	if _, err := os.Stat(hashFile); err == nil {
		// If the file exists, load the previous count
		count, err := loadHashCount(hashFile)
		if err == nil {
			hashCounter = count
		}
	}

	for {
		// Save the hash to the corresponding file
		if hashCounter == 0 {
			err := appendFile("firstHash.txt", fmt.Sprintf("%x\n", firstHash))
			if err != nil {
				fmt.Println("Error saving first hash to file:", err)
				break
			}
		}

		err := appendFile(	, fmt.Sprintf("%x\n", hash))
		if err != nil {
			fmt.Println("Error saving hash to file:", err)
			break
		}

		// Hash the previous hash
		hash = md5.Sum(hash[:])
		hashCounter++

		// Save the hash every 10,000 iterations
		if hashCounter%10000 == 0 {
			err := appendOrChangeLine("10kHash", fmt.Sprintf("%x\n", hash))
			if err != nil {
				fmt.Println("Error saving hash to file:", err)
				break
			}
		}

		fmt.Printf("Hash: %x\n", hash)
	}
}

func writeFile(filename, data string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(data)
	if err != nil {
		return err
	}

	return nil
}

func appendFile(filename, data string) error {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(data)
	if err != nil {
		return err
	}

	return nil
}

func loadHashCount(filename string) (int, error) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return 0, err
	}

	lastLine := string(content)
	lastHash := lastLine[:32]
	count, err := strconv.Atoi(lastHash)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func appendOrChangeLine(filename, data string) error {
    file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0644)
    if err != nil {
        return err
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    var fileContents []string

    addedLine := false

    for scanner.Scan() {
        if !addedLine {
            // If we haven't added a line yet, add the new line
            fileContents = append(fileContents, data)
            addedLine = true
        } else {
            // Otherwise, replace the previously added line
            fileContents = append(fileContents, data)
        }
    }

    if err := scanner.Err(); err != nil {
        return err
    }

    file.Seek(0, 0)
    file.Truncate(0)

    for _, line := range fileContents {
        _, err := file.WriteString(line + "\n")
        if err != nil {
            return err
        }
    }

    return nil
}