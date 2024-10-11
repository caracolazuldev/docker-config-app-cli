package main

import (
	// "bufio"
	// "io"
	"os"
)

type iFileSys interface {
	createOrAppendFile(filename string, content string) error
	OpenFile(filename string, flag int, perm os.FileMode) (*os.File, error)
	WriteString(s string) (n int, err error)
	Close() error
}

type FileSys struct{}

// File I/O
func (f *FileSys) createOrAppendFile(filename string, content string) error {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(content)
	return err
}

func (f *FileSys) OpenFile(filename string, flag int, perm os.FileMode) (*os.File, error) {
	return os.OpenFile(filename, flag, perm)
}

func (f *FileSys) WriteString(s string) (n int, err error) {
	return os.Stdout.WriteString(s)
}

func (f *FileSys) Close() error {
	return os.Stdout.Close()
}
