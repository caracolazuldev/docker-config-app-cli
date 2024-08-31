package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Add config to a file. Type Ctrl+D to quit.")

	fmt.Print("\nFile-name? ")
	filename, err := reader.ReadString('\n')
	OnError("Error reading input: %v\n", err)

	filename = strings.TrimSpace(filename)

	if filename == "" {
		return
	}

	// Remove .conf or .tpl suffix if present
	re := regexp.MustCompile(`\.conf$|\.tpl$`)
	filename = re.ReplaceAllString(filename, "")
	filename += ".tpl"

	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	OnError("Error opening file: %v\n", err)
	defer file.Close()

	for {
		fmt.Print("\nConfig Name? ")
		name, err := reader.ReadString('\n')
		OnError("Error reading config name: %v\n", err)
		name = strings.TrimSpace(name)

		if name == "" {
			return
		}

		fmt.Print("Help text? ")
		help, err := reader.ReadString('\n')
		OnError("Error reading help text: %v\n", err)
		help = strings.TrimSpace(help)

		fmt.Print("export? [yes] ")
		noExp, err := reader.ReadString('\n')
		OnError("Error reading export option: %v\n", err)
		noExp = strings.TrimSpace(noExp)

		prefix := ""
		if noExp != "" {
			prefix = "export "
		}

		line := fmt.Sprintf("%s%s ?= {{%s}}# %s\n", prefix, name, name, help)
		_, err = file.WriteString(line)
		OnError("Error writing to file: %v\n", err)

		fmt.Printf("%s%s ?= {{%s}}# %s >> %s\n", prefix, name, name, help, filename)
	}
}
