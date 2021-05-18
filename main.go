package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"

	help "github.com/callapa1/movies_api/helpers"
	"github.com/callapa1/movies_api/structs"
)

func main() {
	// UI
	fmt.Println("## Welcome to The Movie API ##")

	input := input_loop()

	if input == "/ALL" {
		fetch_all()
	} else if is_number(input) {
		fmt.Println("string")
	}

}

func input_loop() string {
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

func ready() []byte {
	// Open jsonFile, handle error and read
	jsonFile, err := os.Open("movies.json")

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("(Successfully opened movies.json)")
		fmt.Println()
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	return byteValue
}

func fetch_all() {
	byteValue := ready()
	// We unmarshal our byteArray which contains our
	// jsonFile's content into 'movies' which we define in /structs
	var movies structs.Movies
	json.Unmarshal(byteValue, &movies)

	// We iterate through every movie within our movies
	// array and print out their details
	for i := 0; i < len(movies.Movies); i++ {

		help.Print_id(movies.Movies[i].Id)
		help.Print_title(movies.Movies[i].Title)
		help.Print_poster(movies.Movies[i].Poster)
		help.Print_description(movies.Movies[i].Description)
		help.Print_year(movies.Movies[i].Year)

		if len(movies.Movies[i].Actors) > 0 {
			fmt.Println("-Actors: ")
			for _, actor := range movies.Movies[i].Actors {
				help.Print_actor(actor.Name)
			}
		} else {
			fmt.Println("-(Actors info not found)")
		}
	}
}

func is_number(value string) bool {
	if _, err := strconv.Atoi(value); err == nil {
		return true
	} else {
		return false
	}
}
