package main

import (
	"fmt"
	"math/rand"
)

// Cache stores Jokes in memory
type Cache struct {
	ID    int
	Jokes []*Joke
}

// Decide if we want pointers or structs
func (c *Cache) add(joke *Joke) {
	c.Jokes = append(c.Jokes, joke)
}

func (c *Cache) flush() {
	// TODO: write to file
	fmt.Println("Flushing cache")
	c.Jokes = make([]*Joke, 0)
}

// NewCache returns a Cache
func NewCache() Cache {
	id := rand.Intn(100000)
	cache := Cache{id, make([]*Joke, 0)}
	return cache
}
