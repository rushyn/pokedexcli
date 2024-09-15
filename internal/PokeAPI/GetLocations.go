package PokeAPI

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/rushyn/pokedexcli/internal/pokecache"
)

type Locations struct {
	Count    int    	`json:"count"`
	Next     string 	`json:"next"`
	Previous string    `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

var LocationsMap = Locations{
	Count:    0,
	Next:     "https://pokeapi.co/api/v2/location?offset=0&limit=20",
	Previous: "",
	Results:  []struct{Name string "json:\"name\""; URL string "json:\"url\""}{},
}

var cleanSicle = 120 * time.Second
var cache = pokecache.NewCache(cleanSicle)

func GetLocations (action string) error {

	var url string
	if action == "next"{
		if LocationsMap.Next == ""{
			fmt.Println("!!! no next page !!!")
			return nil
		}
		url = LocationsMap.Next
	}
	if action == "back"{
		if LocationsMap.Previous == ""{
			fmt.Println("!!! no back page !!!")
			return nil
		}
		url = LocationsMap.Previous
	}


	var data = []byte{}
	var exists = false
	data, exists = cache.Get(url)
	
	if !exists{
		//fmt.Println("-----------------------running not exits stuff")
		res, err := http.Get(url)

		if err != nil {
			log.Fatal(err)
		}

		data, err = io.ReadAll(res.Body)

		res.Body.Close()
		cache.Add(url, data)

		if res.StatusCode > 299 {
			log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, data)
		}
		if err != nil {
			log.Fatal(err)
		}
	}
	

	err := json.Unmarshal(data, &LocationsMap)
	if err != nil {
		fmt.Println(err)
	}
	
	for _, item := range LocationsMap.Results{
		fmt.Println(item.Name)
	}

	return nil
}

