package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"github.com/joho/godotenv"
	"log"	
	"strings"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

type Movie struct {
	Title       string `json:"title"`
	ReleaseDate string `json:"release_date"`
	Overview    string `json:"overview"`
	VoteAverage float64 `json:"vote_average"`
}

type MovieResponse struct {
	Results []Movie `json:"results"`
}

// The MovieDB API provides access to movie data, including currently playing movies, popular movies, top-rated movies, and upcoming movies.
// This program fetches movie data based on the specified type using command-line flags.	
 

func main() {
	// Define flag
	movieType := flag.String("type", "", "Type of movies to fetch: playing, popular, top, upcoming")
	flag.Parse()

	// Choose the right endpoint
	var endpoint string
	switch *movieType {
	case "playing":
		endpoint = "now_playing"
	case "popular":
		endpoint = "popular"
	case "top":
		endpoint = "top_rated"
	case "upcoming":
		endpoint = "upcoming"
	default:
		fmt.Println("Invalid type. Use: playing, popular, top, upcoming")
		return
	}

	token := os.Getenv("TMDB_API_TOKEN")
	token = strings.Trim(token, "\"") 

	if token == "" {
		fmt.Println("TMDB_API_TOKEN environment variable is not set.")	
		os.Exit(1)
	}

	// Build request
	url := fmt.Sprintf("https://api.themoviedb.org/3/movie/%s?language=en-US&page=1", endpoint)
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("accept", "application/json")
	req.Header.Add("Authorization", "Bearer "+token)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)	
		os.Exit(1)
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		os.Exit(1)
	}

	var data MovieResponse

	if err := json.Unmarshal(body, &data); err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		os.Exit(1)
	}

	for _, movie := range data.Results {
		fmt.Printf("Title: %s\nRelease Date: %s\nOverview: %s\nVote Average: %.1f\n\n",
			movie.Title, movie.ReleaseDate, movie.Overview, movie.VoteAverage)
	}

	fmt.Println("Movies fetched successfully.")
	os.Exit(0)
}
