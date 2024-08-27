package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
)


func prompt(s string) {
	fmt.Print(s)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Add config to a file. Type Ctrl+D to quit.")

	prompt("\nFile-name? ")
	filename, err := reader.ReadString('\n')
	if err != nil {
		if err == io.EOF {
			fmt.Println("\nExiting the CLI. Goodbye!")
			break
		}
		fmt.Println("Error reading input:", err)
		continue
	}
	filename = strings.TrimSpace(filename)

	if filename == "" {
		return
	}

	// Remove .conf or .tpl suffix if present
	re := regexp.MustCompile(`\.conf$|\.tpl$`)
	filename = re.ReplaceAllString(filename, "")
	filename += ".tpl"

	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening file: %v\n", err)
		return
	}
	defer file.Close()

	for {
		prompt("\nConfig Name? ")
		name, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading config name: %v\n", err)
			return
		}
		name = strings.TrimSpace(name)

		if name == "" {
			return
		}

		prompt("Help text? ")
		help, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading help text: %v\n", err)
			return
		}
		help = strings.TrimSpace(help)

		prompt("export? [yes] ")
		noExp, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading export option: %v\n", err)
			return
		}
		noExp = strings.TrimSpace(noExp)

		prefix := ""
		if noExp != "" {
			prefix = "export "
		}

		line := fmt.Sprintf("%s%s ?= {{%s}}# %s\n", prefix, name, name, help)
		_, err = file.WriteString(line)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error writing to file: %v\n", err)
			return
		}

		fmt.Printf("%s%s ?= {{%s}}# %s >> %s\n", prefix, name, name, help, filename)
	}
}
