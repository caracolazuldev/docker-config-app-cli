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

func configFile(stdin *bufio.Reader) (string, error) {

	fmt.Println("Create a config file template, or add configs to one.")
	fmt.Print("File-name? ")
	filename, err := stdin.ReadString('\n')
	if err != nil {
		if errors.Is(err,io.EOF){
			fmt.Println("\nExiting. Goodbye!")
			return "", nil
		}
		fmt.Errorf("could not read input: %v\n", err)
		return "", err
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

	reader := bufio.NewReader(os.Stdin)
	filename, err := configFile(reader);

	if err != nil {
        return cli.Exit(err, 1)
    }

	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	OnError("Error opening file: %v\n", err)
	defer file.Close()

	for {
		fmt.Print("\nConfig Name? ")
		name, err := reader.ReadString('\n')
		OnError("Error reading config name: %v\n", err)
		name = strings.TrimSpace(name)
		if name == "" {
			return cli.Exit("Must provide a config name.", 1)
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

	return nil
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
