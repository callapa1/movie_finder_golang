package structs

import (
	"fmt"
)

// Movies struct which contains
// an array of movies
type Movies struct {
	Movies []Movie `json:"movies"`
}

// Movie struct which contains id, title,
// poster, description, year and actors
type Movie struct {
	Id          uint16  `json:"id"`
	Title       string  `json:"title"`
	Poster      string  `json:"poster"`
	Description string  `json:"description"`
	Year        uint16  `json:"year"`
	Actors      []Actor `json:"actors"`
}

// Actor struct which contains a name
type Actor struct {
	Name string `json:"name"`
}

func (m Movie) Print_movie() {
	m.Print_id()
	m.Print_title()
	m.Print_poster()
	m.Print_description()
	m.Print_year()

	if len(m.Actors) > 0 {
		fmt.Println("-Actors: ")
		for _, actor := range m.Actors {
			actor.Print_actor()
		}
	} else {
		fmt.Println("-(Actors info not found)")
	}
}

// Getters
func (m Movie) Print_id() {
	fmt.Printf("\n-Movie ID: %d\n", m.Id)
}

func (m Movie) Print_title() {
	fmt.Println("-Movie Title: " + m.Title)
}

func (m Movie) Print_poster() {
	if len(m.Poster) < 1 {
		fmt.Println("-(No poster available)")
	} else {
		fmt.Println("-Movie Poster: " + m.Poster)
	}
}

func (m Movie) Print_description() {
	if len(m.Description) < 1 {
		fmt.Println("-(No description available)")
	} else {
		fmt.Println("-Movie Description: " + m.Description)
	}
}

func (m Movie) Print_year() {
	fmt.Printf("-Movie Year: %d\n", m.Year)
}

func (a Actor) Print_actor() {
	fmt.Println("   ", a.Name)
}
