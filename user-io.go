package main

import (
	"fmt"
	"io"
	"os"
	"errors"
)

func OnError(tpl string, err error) {
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


