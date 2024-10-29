package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	// Prompt for the file name
	fmt.Print("Enter file name (be sure to use a valid file extension if needed): ")
	fileName, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading file name:", err)
		return
	}
	fileName = strings.TrimSpace(fileName)

	// Prompt for the file contents
	fmt.Print("Enter file contents: ")
	content, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading file content:", err)
		return
	}

	// Create and write to the file
	err = os.WriteFile(fileName, []byte(content), 0644)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	fmt.Printf("File %s written successfully.\n", fileName)
}
