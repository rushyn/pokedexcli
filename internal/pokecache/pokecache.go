package pokecache

import (
	"sync"
	"time"
)


type cacheEntry struct{
	createdAt 	time.Time
	val			[]byte
}

type Cache struct{
	data	map[string]cacheEntry
	mu		*sync.Mutex
}

func (c Cache) Add(url string, data []byte){
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data[url] = cacheEntry{createdAt: time.Now(), val: data}
}

func (c Cache) Get(url string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	data, exists := c.data[url]
	if exists{
		return data.val, true
	}
	return []byte{}, false
}

func (c Cache) reapLoop(wait time.Duration){
	for {
		currentTime := time.Now()
		c.mu.Lock()
		for key, item := range c.data{
			if currentTime.Sub(item.createdAt) > (wait){
				delete(c.data, key)
			}
		}
		c.mu.Unlock()
		time.Sleep(wait / 30)
	}
}

func NewCache(wait time.Duration) Cache {
	var c =  Cache{
		data: 	map[string]cacheEntry{},
		mu:		&sync.Mutex{},
	}
	go c.reapLoop(wait)
	return c
}