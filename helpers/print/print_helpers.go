package helpers

import (
	"fmt"
)

func Print_id(id uint16) {
	fmt.Printf("\n-Movie ID: %d\n", id)
}

func Print_title(title string) {
	fmt.Println("-Movie Title: " + title)
}

func Print_poster(poster string) {
	if len(poster) < 1 {
		fmt.Println("-(No poster available)")
	} else {
		fmt.Println("-Movie Poster: " + poster)
	}
}

func Print_description(description string) {
	fmt.Println("-Movie: " + description)
}

func Print_year(year uint16) {
	fmt.Printf("-Movie Year: %d\n", year)
}

func Print_actor(actor string) {
	fmt.Println(" ", actor)
}
