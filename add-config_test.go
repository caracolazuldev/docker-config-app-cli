package main

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
)

//go:generate mockgen -destination=mock_io.go -package=main github.com/urfave/cli/v2 App,Context

func TestAddConfig(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockApp := NewMockApp(ctrl)
	mockCtx := NewMockContext(ctrl)

	// Mock the createOrAppendFile function
	mockCreateOrAppendFile := func(filename, content string) error {
		if filename == "test.conf.tpl" && content == "export TEST ?= {{TEST}}# This is a test config\n" {
			return nil
		}
		return errors.New("unexpected file or content")
	}
	createOrAppendFile = mockCreateOrAppendFile

	// Mock the readUserInput function
	mockReadUserInput := func(prompt string) (string, error) {
		switch prompt {
		case "File-name? ":
			return "test.conf", nil
		case "\nConfig Name? ":
			return "TEST", nil
		case "Help text? ":
			return "This is a test config", nil
		case "export? [yes] ":
			return "", nil
		default:
			return "", errors.New("unexpected prompt")
		}
	}
	readUserInput = mockReadUserInput

	// Call the addConfig function
	err := addConfig(mockCtx)
	if err != nil {
		t.Errorf("addConfig returned an error: %v", err)
	}
}
