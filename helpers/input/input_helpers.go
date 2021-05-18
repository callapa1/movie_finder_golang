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
		fmt.Println("  [keywords] -> any words included in a title")

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
