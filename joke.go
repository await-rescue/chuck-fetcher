package main

const randomJokeURL = "http://api.icndb.com/jokes/random"

// A Joke object
type Joke struct {
	ID   int    `json:"id"`
	Joke string `json:"joke"`
}

// A JokeResponse from the API for Joke objects
type JokeResponse struct {
	Type  string `json:"type"`
	Value Joke   `json:"value"`
}
