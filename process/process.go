package process

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"

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
	byteValue := Ready()
	// We unmarshal our byteArray which contains our
	// jsonFile's content into 'movies' which we define in /structs
	var movies structs.Movies
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
	byteValue := Ready()

	var movies structs.Movies
	json.Unmarshal(byteValue, &movies)

	var movie structs.Movie

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

func Keywords(input string) {
	fmt.Println("Type keywords below")
	fmt.Print("\n>")
	var keywords string

	fmt.Scan(&keywords)

	fmt.Println(keywords)

	// ENCONTRAR COMO SEPARAR VARIAS PALABRAS
}
