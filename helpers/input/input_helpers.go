package helpers

import (
	"fmt"
	"strconv"
)

func Input_loop() string {
	var input string

	for {
		fmt.Println("-Type:")
		fmt.Println("  /ALL -> to fetch all movies")
		fmt.Println("  [number] -> to fetch movie by ID")
		fmt.Println("  words -> to type keywords included in titles")
		fmt.Println("")

		fmt.Print(">")
		fmt.Scanln(&input)
		if len(input) > 0 {
			break
		}
	}
	return input
}

func Is_number(value string) bool {
	if _, err := strconv.Atoi(value); err == nil {
		return true
	} else {
		return false
	}
}
