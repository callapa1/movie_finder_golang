package process

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	h_input "github.com/callapa1/movies_api/helpers/input"
	h_print "github.com/callapa1/movies_api/helpers/print"
	structs "github.com/callapa1/movies_api/structs"
)

func Ready() []byte {
	// Open jsonFile, handle error and read
	jsonFile, err := os.Open("movies.json")

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("(Successfully opened movies.json)")
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	return byteValue
}

func Fetch_all() {
	var movies structs.Movies
	byteValue := Ready()
	// We unmarshal our byteArray which contains our
	// jsonFile's content into 'movies' which we define in /structs
	json.Unmarshal(byteValue, &movies)

	// We iterate through every movie within our movies
	// array and print out their details
	for i := 0; i < len(movies.Movies); i++ {

		h_print.Print_id(movies.Movies[i].Id)
		h_print.Print_title(movies.Movies[i].Title)
		h_print.Print_poster(movies.Movies[i].Poster)
		h_print.Print_description(movies.Movies[i].Description)
		h_print.Print_year(movies.Movies[i].Year)

		if len(movies.Movies[i].Actors) > 0 {
			fmt.Println("-Actors: ")
			for _, actor := range movies.Movies[i].Actors {
				h_print.Print_actor(actor.Name)
			}
		} else {
			fmt.Println("-(Actors info not found)")
		}
	}
}

func Fetch_by_id(id string) {
	var movies structs.Movies
	var movie structs.Movie
	byteValue := Ready()
	json.Unmarshal(byteValue, &movies)

	for i := 0; i < len(movies.Movies); i++ {
		if strconv.Itoa(int(movies.Movies[i].Id)) == id {
			movie = movies.Movies[i]
			break
		}
	}

	if movie.Title == "" {
		fmt.Println("No movie matches the ID!")
	} else {
		h_print.Print_id(movie.Id)
		h_print.Print_title(movie.Title)
		h_print.Print_poster(movie.Poster)
		h_print.Print_description(movie.Description)
		h_print.Print_year(movie.Year)

		if len(movie.Actors) > 0 {
			fmt.Println("-Actors: ")
			for _, actor := range movie.Actors {
				h_print.Print_actor(actor.Name)
			}
		} else {
			fmt.Println("-(Actors info not found)")
		}
	}

}

func Keywords() {
	var movies structs.Movies
	var movie structs.Movie
	byteValue := Ready()
	json.Unmarshal(byteValue, &movies)

	fmt.Println()
	fmt.Println("Type one word per line (5 max)")
	fmt.Println()

	input := h_input.Keywords_loop()

	for _, v := range input {
		for i := 0; i < len(movies.Movies); i++ {
			if strings.Contains(movies.Movies[i].Title, v) {
				movie = movies.Movies[i]
				break
			}
		}

		if len(movie.Title) > 0 {
			break
		}
	}

	h_print.Print_title(movie.Title)
}
