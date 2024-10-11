package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
)

type iUser interface {
	OnError(tpl string, err error)
	ReadLn(prompt string) (string, error)
	PrintLn(a ...interface{}) (int, error)
}

type Terminal struct{}

func (u *Terminal) OnError(tpl string, err error) {
	if tpl == "" {
		tpl = "Error reading input: %v\n"
	}

	if err != nil {
		if errors.Is(err, io.EOF) {
			fmt.Println("\nExiting. Goodbye!")
			return
		}
		fmt.Fprintf(os.Stderr, tpl, err)
		return
	}
}

// User Input
func (u *Terminal) ReadLn(prompt string) (string, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)
	return reader.ReadString('\n')
}

func (u *Terminal) PrintLn(a ...interface{}) (int, error) {
	return fmt.Println(a...)
}
