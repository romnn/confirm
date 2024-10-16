package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func confirm(message string, tries int) bool {
	r := bufio.NewReader(os.Stdin)
	for tries > 0 {
		tries--
		fmt.Printf("%s [y/n]: ", message)

		res, err := r.ReadString('\n')
		if err != nil {
			panic(err)
		}
		// empty input (i.e. "\n")
		if len(res) < 2 {
			continue
		}
		return strings.ToLower(strings.TrimSpace(res))[0] == 'y'
	}
	return false
}

func main() {
	message := "confirm"
	if len(os.Args) > 1 {
		message = os.Args[1]
	}
	if confirm(message, 1) {
		os.Exit(0)
	} else {
		os.Exit(1)
	}
}
