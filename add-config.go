package main

import (
	"errors"
	"fmt"
	"io"
	"regexp"
	"sort"
	"strings"

	"github.com/urfave/cli/v2"
)

func configFile() (string, error) {
	var user = services.Term

	user.PrintLn("Create a config file template, or add configs to one.\n")

	filename, err := user.ReadLn("File-name? ")
	if err != nil {
		if errors.Is(err, io.EOF) {
			return "", err
		}
		panic(fmt.Sprintf("Could not read input: %v", err))
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

func promptForConfig() (string, string, string, error) {
	var user = services.Term

	// overly complex, but I'm just learning to use maps and slices

	type prompt struct {
		Order  int
		Name   string
		Prompt string
	}
	prompts := []prompt{
		{0, "name", "Config Name?"},
		{1, "help", "Help text?"},
		{2, "noExp", "export? [no]"},
	}

	sort.Slice(prompts, func(i, j int) bool {
		return prompts[i].Order < prompts[j].Order
	})

	results := make(map[string]string)

	for _, prompt := range prompts {
		input, err := user.ReadLn(prompt.Prompt + " ")
		if err != nil {
			if !errors.Is(err, io.EOF) {
				err = fmt.Errorf("error reading %s: %v", prompt.Prompt, err)
			}
			return "", "", "", err
		}
		input = strings.TrimSpace(input)
		if prompt.Name == "name" && input == "" {
			return "", "", "", errors.New("must provide a config name")
		}
		results[prompt.Name] = input
	}

	if results["noExp"] != "" &&
		!strings.HasPrefix(strings.ToLower(results["noExp"]), "n") {
		results["noExp"] = "export "
	}

	return results["noExp"], results["name"], results["help"], nil
}

/**
 * Main()
 */
func addConfig(ctx *cli.Context) error {

	var term = services.Term

	term.PrintLn("Hello, I am the add-config tool.")
	term.PrintLn("[Type Ctrl+D when done]\n")

	filename, err := configFile()
	if err != nil {
		if errors.Is(err, io.EOF) {
			return cli.Exit("\nExiting. Goodbye!", 0)
		}
		return cli.Exit(err, 1)
	}

	for {
		prefix, name, help, err := promptForConfig()

		if err != nil {
			return cli.Exit(err, 1)
		}

		line := fmt.Sprintf("%s%s ?= {{%s}}# %s\n", prefix, name, name, help)

		if err := services.Files.createOrAppendFile(filename, line); err != nil {
			return cli.Exit(fmt.Errorf("error writing to file: %v", err), 1)
		}

		fmt.Printf("%s%s ?= {{%s}}# %s >> %s\n", prefix, name, name, help, filename)
	}
}
