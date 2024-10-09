package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/urfave/cli/v2"
)

// File I/O
func createOrAppendFile(filename string, content string) error {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(content)
	return err
}

// User Input
func readUserInput(prompt string) (string, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)
	return reader.ReadString('\n')
}


func configFile() (string, error) {
	fmt.Println("Create a config file template, or add configs to one.")
	filename, err := readUserInput("File-name? ")
	if err != nil {
		if errors.Is(err, io.EOF) {
			fmt.Println("\nExiting. Goodbye!")
			return "", nil
		}
		return "", fmt.Errorf("could not read input: %v", err)
	}

	filename = strings.TrimSpace(filename)
	if filename == "" {
		return "", fmt.Errorf("filename is required")
	}

	// Remove .tpl suffix if present
	re := regexp.MustCompile(`\.tpl$`)
	filename = re.ReplaceAllString(filename, "")
	// if no suffix, add .conf
	if !strings.Contains(filename, ".") {
		filename += ".conf"
	}
	filename += ".tpl"

	return filename, nil
}

func addConfig(ctx *cli.Context) error {
	fmt.Println("Hello, I am the add-config tool.")
	fmt.Println("[Type Ctrl+D when done]\n")

	filename, err := configFile()
	if err != nil {
		return cli.Exit(err, 1)
	}

	for {
		name, err := readUserInput("\nConfig Name? ")
		if err != nil {
			return cli.Exit(fmt.Errorf("error reading config name: %v", err), 1)
		}
		name = strings.TrimSpace(name)
		if name == "" {
			return cli.Exit("Must provide a config name.", 1)
		}

		help, err := readUserInput("Help text? ")
		if err != nil {
			return cli.Exit(fmt.Errorf("error reading help text: %v", err), 1)
		}
		help = strings.TrimSpace(help)

		noExp, err := readUserInput("export? [yes] ")
		if err != nil {
			return cli.Exit(fmt.Errorf("error reading export option: %v", err), 1)
		}
		noExp = strings.TrimSpace(noExp)

		prefix := ""
		if noExp != "" {
			prefix = "export "
		}

		line := fmt.Sprintf("%s%s ?= {{%s}}# %s\n", prefix, name, name, help)
		if err := createOrAppendFile(filename, line); err != nil {
			return cli.Exit(fmt.Errorf("error writing to file: %v", err), 1)
		}

		fmt.Printf("%s%s ?= {{%s}}# %s >> %s\n", prefix, name, name, help, filename)
	}
}

func main() {
	app := &cli.App{
		Name:  "add-config",
		// Flags: []cli.BoolFlag{
		// 	Name: "-q"
		// 	Usage: "quiet mode",
		// },
		Action: addConfig,
	}

	if err := app.Run(os.Args); err != nil {
        log.Fatal(err)
    }
}
