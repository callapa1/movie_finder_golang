package main

import (
	"fmt"

	h_input "github.com/callapa1/movies_api/helpers/input"
	process "github.com/callapa1/movies_api/process"
)

func main() {
	// UI
	fmt.Println("## Welcome to The Movie API ##")
	input := h_input.Input_loop()

	if input == "/ALL" {
		process.Fetch_all()
	} else if h_input.Is_number(input) {
		process.Fetch_by_id()
	}
}
