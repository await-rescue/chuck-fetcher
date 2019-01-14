package main

import (
	"encoding/json"
	"errors"
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

	// return this so we can add to cache
	// check whether it's safe to return a pointer
	fmt.Println(response)

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
	} else {
		return errors.New("Fetcher is already running")
	}

	return nil
}

func (f *Fetcher) stop() {
	f.status = "stopped"
	fmt.Println(f.cache.Jokes[0])
	f.cache.flush()
}

func (f *Fetcher) flushCacheTimer() {
	ticker := time.NewTicker(time.Minute)
	for range ticker.C {
		f.cache.flush()
	}
}

// NewFetcher returns a new Fetcher
func NewFetcher() *Fetcher {
	cache := NewCache()
	fetcher := Fetcher{"stopped", cache}
	// Can we delay starting this until run?
	go fetcher.flushCacheTimer()
	return &fetcher
}
