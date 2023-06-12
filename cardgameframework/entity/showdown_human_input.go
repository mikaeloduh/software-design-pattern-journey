package entity

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type ShowdownHumanInput struct{}

func (i ShowdownHumanInput) InputString() string {
	reader := bufio.NewScanner(os.Stdin)

	for {
		reader.Scan()
		text := strings.TrimSpace(reader.Text())
		if text != "" {
			return text
		}
		fmt.Print("Invalid input. Please enter a non-empty string: ")
	}
}

func (i ShowdownHumanInput) InputNum(min int, max int) int {
	fmt.Print("(", min, " ~ ", max, "): ")

	reader := bufio.NewReader(os.Stdin)

	for {
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		input = strings.TrimSpace(input)
		num, err := strconv.Atoi(input)
		if err != nil {
			fmt.Print("Invalid input. Please enter a number: ")
			continue
		}

		if num < min || num > max {
			fmt.Print("Invalid input. Number must be between ", min, " and ", max, ": ")
			continue
		}

		return num
	}
}

func (i ShowdownHumanInput) InputBool() bool {
	fmt.Print("(y/n): ")

	scanner := bufio.NewScanner(os.Stdin)

	for {
		scanner.Scan()
		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "reading standard input:", err)
			os.Exit(1)
		}
		input := strings.ToLower(scanner.Text())
		if input == "y" {
			return true
		} else if input == "n" {
			return false
		} else {
			fmt.Print("Invalid input, please try again: ")
		}
	}
}
