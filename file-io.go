package main

import (
	// "bufio"
	// "io"
	"os"
)

type FileIO interface {
	createOrAppendFile(filename string, content string) error
	OpenFile(filename string, flag int, perm os.FileMode) (*os.File, error)
	WriteString(s string) (n int, err error)
	Close() error
}

type File struct {
	file *os.File
}

// File I/O
func (f *File) createOrAppendFile(filename string, content string) error {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(content)
	return err
}
