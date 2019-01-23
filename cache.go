package main

import (
	"fmt"
	"log"
	"os"
)

// Cache stores data in memory - check capitalisation
type Cache struct {
	path     string
	filename string
	data     map[int]string
}

func (c *Cache) addJoke(joke *Joke) {
	c.data[joke.ID] = joke.Joke
}

// Flush the cache and dump to a file
func (c *Cache) flush() {
	file, err := os.OpenFile(c.path+c.filename, os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Fatalf("Failed opening file: %s", err)
	}
	defer file.Close()

	for _, value := range c.data {
		_, err := file.WriteString(fmt.Sprintf("%s%s", value, "\n"))
		if err != nil {
			log.Fatalf("Failed writing to file: %s", err)
		}
	}

	c.data = make(map[int]string)
	log.Println("Cache flushed")
}

// NewCache returns a Cache and creates files for cache flushing
func NewCache(path string, filename string) *Cache {
	cache := Cache{path: path, filename: filename, data: make(map[int]string)}

	// Clear any existing cache file, ignore errors if it doesn't exist
	_ = os.RemoveAll(path)

	err := os.Mkdir(path, os.ModePerm)
	if err != nil {
		log.Printf("Failed to create dir: %s", err)
	}

	_, err = os.Create(path + filename)
	if err != nil {
		log.Printf("Failed to create file: %s", err)
	}

	return &cache
}
