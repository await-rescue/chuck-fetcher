package main

import (
	"fmt"
	"math/rand"
)

// Cache stores Jokes in memory
type Cache struct {
	ID    int
	Jokes map[int]string
}

func (c *Cache) add(joke *Joke) {
	c.Jokes[joke.ID] = joke.Joke
}

func (c *Cache) flush() {
	// TODO: write to file
	fmt.Println("Flushing cache")
	fmt.Println(c.Jokes)
	c.Jokes = make(map[int]string)
}

// NewCache returns a Cache
func NewCache() *Cache {
	id := rand.Intn(100000)
	cache := Cache{id, make(map[int]string)}
	return &cache
}
