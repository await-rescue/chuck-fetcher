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

func (f *Fetcher) getRandomJoke() (*Joke, error) {
	timeout := time.Duration(5 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}
	response := JokeResponse{}

	res, err := client.Get(randomJokeURL)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	// TODO: check response code

	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		panic(err.Error())
	}

	return &response.Value, nil
}

func (f *Fetcher) run() {
	for f.status == "running" {
		joke, err := f.getRandomJoke()
		if err != nil {
			// retry
			continue
		}
		// TODO: could try again if we get an existing key
		f.cache.addJoke(joke)
		time.Sleep(3 * time.Second)
	}
}

func (f *Fetcher) start() error {
	if f.status != "running" {
		f.status = "running"
		// Could just put the function here in full to avoid run() being called
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
	cache := NewCache("./cache/", "joke_data.txt")
	fetcher := Fetcher{"stopped", cache}
	// TODO: Can we delay starting this until run (and not end up with multiple)?
	go fetcher.flushCacheTimer()
	return &fetcher
}
