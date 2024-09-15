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
	//fmt.Println("-----------------------running cache add")
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data[url] = cacheEntry{createdAt: time.Now(), val: data}
}

func (c Cache) Get(url string) ([]byte, bool) {
	//fmt.Println("-----------------------get first line")
	c.mu.Lock()
	defer c.mu.Unlock()
	//fmt.Println("-----------------------get after lock")
	data, exists := c.data[url]
	if exists{
		//fmt.Println("-----------------------get exists")
		return data.val, true
	}
	//fmt.Println("-----------------------get not exist")
	return []byte{}, false
}

func (c Cache) reapLoop(wait time.Duration){
	for {
		currentTime := time.Now()
		c.mu.Lock()
		for key, item := range c.data{
			if currentTime.Sub(item.createdAt) > (wait){
				//fmt.Println("-----------------------deleting map keys")
				//fmt.Println("-----------------------" + key)
				//fmt.Println("----------------------------------------")
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