package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// Fetcher gets data from an external API
type Fetcher struct {
	// Check capitalisation stuff
	status string
	cache  *Cache
}

// TODO: we could parameterise this
func (f *Fetcher) getRandomJoke() {
	response := Response{}

	// TODO: timeouts
	res, err := http.Get(randomJokeURL)
	if err != nil {
		// Handle this in a better way
		panic(err.Error())
	}

	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		panic(err.Error())
	}

	// Should return this so we can add to cache
	fmt.Println(response)
}

func (f *Fetcher) run() {
	for {
		time.Sleep(1 * time.Second)
		if f.status != "running" {
			break
		}
		f.getRandomJoke()
	}
}

// Gets jokes and stores them - we could assign an ID for filename so we can run multiple
func (f *Fetcher) start() {
	if f.status != "running" {
		f.status = "running"
		f.run()
	} else {
		// Could be an error
		fmt.Println("Processor is already running")
	}
	// TODO: cache flushing
}

func (f *Fetcher) stop() {
	f.status = "stopped"
	f.cache.flush()
}

// NewFetcher returns a new Fetcher
func NewFetcher() Fetcher {
	cache := NewCache()
	fetcher := Fetcher{"stopped", &cache}
	return fetcher
}
