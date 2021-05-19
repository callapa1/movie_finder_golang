package process

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	h_input "github.com/callapa1/movies_api/helpers/input"
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

	movies_array := movies.Movies
	fmt.Println(movies_array)
	// We iterate through every movie within our movies
	// array and print out their details
	for i := 0; i < len(movies_array); i++ {
		movies_array[i].Print_movie()
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
		movie.Print_movie()
	}

}

func Keywords() {
	var movies structs.Movies
	var movie structs.Movie
	byteValue := Ready()
	json.Unmarshal(byteValue, &movies)
	movies_array := movies.Movies

	fmt.Println()
	fmt.Println("Type one word per line (5 max)")
	fmt.Println()

	input := h_input.Keywords_loop()

	for _, v := range input {
		for i := 0; i < len(movies_array); i++ {
			if strings.Contains(movies_array[i].Title, v) {
				movie = movies_array[i]
				break
			}
		}

		if len(movie.Title) > 0 {
			break
		}
	}

	movie.Print_movie()
}
