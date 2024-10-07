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
		if cmd == "exit" {
			os.Exit(0)
		} else if cmd == "echo" {
			fmt.Fprint(os.Stdout, strings.Join(chunks[1:], " "))
		} else {
			fmt.Fprintf(os.Stdout, "%v: command not found\n", strings.TrimRight(buff, "\n"))
		}
	}

}
