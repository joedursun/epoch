package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	dateFormat = "2006-01-02 15:04:05"
)

func parseStdin() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		msString := strings.TrimSpace(scanner.Text())
		if len(msString) < 13 {
			fmt.Println("Invalid timestamp. Must be at least 13 digits long.")
			os.Exit(1)
		}

		ms, err := strconv.Atoi(msString)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		switch len(msString) {
		case 13:
			// milliseconds
			fmt.Println(time.Unix(0, int64(ms)*1_000_000).Format(dateFormat))
		case 16:
			// microseconds
			fmt.Println(time.Unix(0, int64(ms)*1_000).Format(dateFormat))
		case 19:
			// nanoseconds
			fmt.Println(time.Unix(0, int64(ms)).Format(dateFormat))
		default:
      fmt.Printf("Invalid timestamp: %d. Must be 13, 16, or 19 digits long.\n", ms)
		}
	}
}

func main() {
	var input string
	if len(os.Args) < 2 {
		_, err := fmt.Scanln(&input)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	} else if os.Args[1] == "-s" {
		parseStdin()
	} else {
		input = os.Args[1]
	}

	msString := strings.TrimSpace(input)
	if len(msString) < 13 {
    fmt.Printf("Invalid timestamp: %s. Must be at least 13 digits long.\n", msString)
		os.Exit(1)
	}

	ms, err := strconv.Atoi(msString)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	switch len(msString) {
	case 13:
		// milliseconds
		fmt.Println(time.Unix(0, int64(ms)*1_000_000).Format(dateFormat))
	case 16:
		// microseconds
		fmt.Println(time.Unix(0, int64(ms)*1_000).Format(dateFormat))
	case 19:
		// nanoseconds
		fmt.Println(time.Unix(0, int64(ms)).Format(dateFormat))
	}
}
