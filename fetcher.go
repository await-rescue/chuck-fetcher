package main

import (
	"encoding/json"
	"errors"
	"log"
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
func (f *Fetcher) getRandomJoke() *Joke {
	response := Response{}

	timeout := time.Duration(5 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}
	res, err := client.Get(randomJokeURL)
	if err != nil {
		// Handle this in a better way/check timeouts work
		panic(err.Error())
	}

	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		panic(err.Error())
	}

	return &response.Value
}

func (f *Fetcher) run() {
	for {
		time.Sleep(1 * time.Second)
		if f.status != "running" {
			break
		}
		joke := f.getRandomJoke()
		f.cache.add(joke)
	}
}

// Gets jokes and stores them - we could assign an ID for filename so we can run multiple
func (f *Fetcher) start() error {
	if f.status != "running" {
		f.status = "running"
		go f.run()
		log.Println("Fetcher is running")
	} else {
		return errors.New("Fetcher is already running")
	}
	return nil
}

func (f *Fetcher) stop() error {
	if f.status != "stopped" {
		f.status = "stopped"
		f.cache.flush()
		log.Println("Fetcher stopped")
	} else {
		return errors.New("Fetcher is already stopped")
	}
	return nil
}

func (f *Fetcher) flushCacheTimer() {
	ticker := time.NewTicker(time.Minute)
	for range ticker.C {
		f.cache.flush()
	}
}

// NewFetcher returns a new Fetcher
func NewFetcher() *Fetcher {
	cache := NewCache("./cache/", "data.txt")
	fetcher := Fetcher{"stopped", cache}
	// Can we delay starting this until run?
	go fetcher.flushCacheTimer()
	return &fetcher
}
