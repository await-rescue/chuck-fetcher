package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func getInput() string {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	// Remove newline
	text = strings.Replace(text, "\n", "", -1)
	return text
}

func main() {
	fetcher := NewFetcher()
	for {
		text := getInput()

		switch text {
		case "start":
			err := fetcher.start()
			if err != nil {
				fmt.Println(err)
			}
		case "stop":
			err := fetcher.stop()
			if err != nil {
				fmt.Println(err)
			}
		case "exit":
			err := fetcher.stop()
			if err != nil {
				fmt.Println(err)
			}
			os.Exit(0)
		default:
			fmt.Println(fmt.Sprintf("Invalid command: %s\n Available commands - [start, stop, exit]", text))
		}
	}
}
