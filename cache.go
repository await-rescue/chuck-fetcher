package main

import (
	"fmt"
	"log"
	"os"
)

// Cache stores Jokes in memory - check capitalisation
type Cache struct {
	path     string
	filename string
	Jokes    map[int]string
}

func (c *Cache) add(joke *Joke) {
	c.Jokes[joke.ID] = joke.Joke
}

func (c *Cache) flush() {
	file, err := os.OpenFile(c.path+c.filename, os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Fatalf("Failed opening file: %s", err)
	}

	for _, v := range c.Jokes {
		_, err := file.WriteString(fmt.Sprintf("%s%s", v, "\n"))
		if err != nil {
			log.Fatalf("Failed writing to file: %s", err)
		}
	}

	c.Jokes = make(map[int]string)
}

// NewCache returns a Cache
func NewCache(path string, filename string) *Cache {
	cache := Cache{path, filename, make(map[int]string)}

	// Clear any existing cache file, ignore errors if it doesn't exist
	_ = os.RemoveAll(path)

	// TODO: make persistant in docker
	err := os.Mkdir(path, os.ModePerm)
	if err != nil {
		log.Fatalf("Failed to create dir: %s", err)
	}

	_, err = os.Create(path + filename)
	if err != nil {
		log.Fatalf("Failed to create file: %s", err)
	}

	return &cache
}
