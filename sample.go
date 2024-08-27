package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Welcome to the interactive CLI! Type 'exit' to quit.")

	for {
		fmt.Print("> ")
		input, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				fmt.Println("\nExiting the CLI. Goodbye!")
				break
			}
			fmt.Println("Error reading input:", err)
			continue
		}

		// Trim the input to remove newline characters
		input = strings.TrimSpace(input)

		// Check for exit command
		if input == "exit" {
			fmt.Println("Exiting the CLI. Goodbye!")
			break
		}

		// Process the input (for demonstration, we just echo it back)
		fmt.Println("You entered:", input)
	}
}
