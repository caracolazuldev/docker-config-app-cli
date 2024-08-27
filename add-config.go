package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
)


func prompt(s string) {
	fmt.Print(s)
}

func errorMsg(tpl string, err error) {
	if tpl == "" {
		tpl = "Error reading input: %v\n"
	}

	if err != nil {
		if errors.Is(err,io.EOF){
			fmt.Println("\nExiting. Goodbye!")
			return
		}
		fmt.Fprintf(os.Stderr, tpl, err)
		return
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Add config to a file. Type Ctrl+D to quit.")

	prompt("\nFile-name? ")
	filename, err := reader.ReadString('\n')
	errorMsg("Error reading input: %v\n", err)

	filename = strings.TrimSpace(filename)

	if filename == "" {
		return
	}

	// Remove .conf or .tpl suffix if present
	re := regexp.MustCompile(`\.conf$|\.tpl$`)
	filename = re.ReplaceAllString(filename, "")
	filename += ".tpl"

	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	errorMsg("Error opening file: %v\n", err)
	defer file.Close()

	for {
		prompt("\nConfig Name? ")
		name, err := reader.ReadString('\n')
		errorMsg("Error reading config name: %v\n", err)
		name = strings.TrimSpace(name)

		if name == "" {
			return
		}

		prompt("Help text? ")
		help, err := reader.ReadString('\n')
		errorMsg("Error reading help text: %v\n", err)
		help = strings.TrimSpace(help)

		prompt("export? [yes] ")
		noExp, err := reader.ReadString('\n')
		errorMsg("Error reading export option: %v\n", err)
		noExp = strings.TrimSpace(noExp)

		prefix := ""
		if noExp != "" {
			prefix = "export "
		}

		line := fmt.Sprintf("%s%s ?= {{%s}}# %s\n", prefix, name, name, help)
		_, err = file.WriteString(line)
		errorMsg("Error writing to file: %v\n", err)

		fmt.Printf("%s%s ?= {{%s}}# %s >> %s\n", prefix, name, name, help, filename)
	}
}
