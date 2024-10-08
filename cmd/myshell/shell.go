package main

import (
	"errors"
	"os"
)

type builtin string

const (
	ECHO builtin = "echo"
	EXIT builtin = "exit"
	TYPE builtin = "type"
)

var builtinCommands = []builtin{ECHO, EXIT, TYPE}

func searchForCommand(path string, cmd string) (found bool) {
	pathMetadata, err := os.Stat(path)
	if errors.Is(err, os.ErrNotExist) {
		return
	}
	if err != nil {
		panic(err)
	}
	if pathMetadata.IsDir() {
		contents, err := os.ReadDir(path)
		if errors.Is(err, os.ErrNotExist) {
			return
		}
		if err != nil {
			panic(err)
		}
		for _, c := range contents {
			if c.Name() == cmd && !c.IsDir() {
				found = true
				break
			}
		}
		// Look inside directory
	} else if pathMetadata.Name() == cmd {
		found = true
	}
	return
}
