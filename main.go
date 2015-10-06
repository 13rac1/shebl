package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
)

func main() {
	// Get the filename to work with
	scriptPath := ""
	if len(os.Args) > 1 {
		scriptPath = os.Args[1]
	} else {
		log.Fatal("Filename required")
		os.Exit(1)
	}

	file, err := os.Open(scriptPath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Create a bash command and connect the pipes
	cmd := exec.Command("/bin/bash")
	stdin, err := cmd.StdinPipe()
	stdout, err := cmd.StdoutPipe()
	stderr, err := cmd.StderrPipe()
	// Mind the pipes.
	go printPipe(stdout)
	go printPipe(stderr)

	err = cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Press 'Enter' to alternatively show the commands and run them.")

	userin := bufio.NewReader(os.Stdin)

	start := true
	s := bufio.NewScanner(file)
	for s.Scan() {
		if len(s.Text()) != 0 {
			if start {
				start = false
			} else {
				// Only require enter after the second line in the script
				userin.ReadBytes('\n')
			}
			// Print a pretend prompt
			fmt.Print("$ " + s.Text())
			// Have the user press enter to run it.
			userin.ReadBytes('\n')

			// Send the line to bash
			stdin.Write(s.Bytes())
			stdin.Write([]byte("\n"))
		}
	}
	if s.Err() != nil {
		log.Fatal(s.Err())
	}

	stdin.Close()

	err = cmd.Wait()
}

// Given an io input source outputs until EOF.
func printPipe(pipe io.ReadCloser) {
	s := bufio.NewScanner(pipe)

	for s.Scan() {
		fmt.Println(s.Text())
	}
}
