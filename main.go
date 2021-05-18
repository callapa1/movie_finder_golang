package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	help "github.com/callapa1/movies_api/helpers"
	"github.com/callapa1/movies_api/structs"
)

func main() {

	// Open our jsonFile
	jsonFile, err := os.Open("movies.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("(Successfully opened movies.json)")
		fmt.Println()
	}
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()
	// read our opened xmlFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	fmt.Println("## Welcome to The Movie API ##")
	fmt.Println("-Type:")
	fmt.Println("  ALL -> to fetch all movies")
	fmt.Println("  [ID] -> any number to fetch movie by ID")
	fmt.Println("  [keywords] -> any words included in a title")
	var answer string
	fmt.Scanln(&answer)

	// Movies array initialized
	var movies structs.Movies

	// We unmarshal our byteArray which contains our
	// jsonFile's content into 'movies' which we defined above
	json.Unmarshal(byteValue, &movies)

	// we iterate through every movie within our movies array and
	// print out the movie ID, title, poster, description, year, actors
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
