package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Uncomment this block to pass the first stage
	for {
		fmt.Fprint(os.Stdout, "$ ")
		// Wait for user input
		buff, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			panic(err)
		}
		chunks := strings.Split(buff, " ")
		cmd := chunks[0]
		if cmd == string(EXIT) {
			os.Exit(0)
		} else if cmd == string(ECHO) {
			fmt.Fprint(os.Stdout, strings.Join(chunks[1:], " "))
		} else if cmd == string(TYPE) {
			testType := strings.TrimRight(chunks[1], "\n")
			if testType == string(ECHO) || testType == string(TYPE) || testType == string(EXIT) {
				fmt.Fprintf(os.Stdout, "%v is a shell builtin\n", testType)
			} else {
				fmt.Fprintf(os.Stdout, "%v: not found\n", testType)
			}
		} else {
			fmt.Fprintf(os.Stdout, "%v: command not found\n", strings.TrimRight(buff, "\n"))
		}
	}

}
