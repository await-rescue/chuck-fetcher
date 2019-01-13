package main

const randomJokeURL = "http://api.icndb.com/jokes/random"

// A Joke object
type Joke struct {
	ID   int    `json:"id"`
	Joke string `json:"joke"`
}

// A Response from the API
type Response struct {
	Type  string `json:"type"`
	Value Joke   `json:"value"`
}
