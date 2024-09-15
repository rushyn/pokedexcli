package PokeAPI

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type exploer struct {
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	GameIndex int `json:"game_index"`
	ID        int `json:"id"`
	Location  struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Name  string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			EncounterDetails []struct {
				Chance          int `json:"chance"`
				ConditionValues []struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"condition_values"`
				MaxLevel int `json:"max_level"`
				Method   struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
				MinLevel int `json:"min_level"`
			} `json:"encounter_details"`
			MaxChance int `json:"max_chance"`
			Version   struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}

type region struct {
	Areas []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"areas"`
	GameIndices []struct {
		GameIndex  int `json:"game_index"`
		Generation struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"generation"`
	} `json:"game_indices"`
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	Region struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"region"`
}


func GetPokemon (arear string) error {
	var regionArear = region{}
	var url = "https://pokeapi.co/api/v2/location/" + arear + "/"
	var data = []byte{}
	var exists = false
	data, exists = cache.Get(url)
	
	if !exists{
		res, err := http.Get(url)

		if err != nil {
			log.Fatal(err)
		}

		data, err = io.ReadAll(res.Body)

		res.Body.Close()
		cache.Add(url, data)

		if res.StatusCode > 299 {
			fmt.Printf("Response failed with status code: %d \ncheck you spelling and try agine: \n", res.StatusCode)
			return err
		}

	}
	

	err := json.Unmarshal(data, &regionArear)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Exploring pastoria-city-area...")
	fmt.Println("Found Pokemon:")
	
	for _, item := range regionArear.Areas{
		var arearExplore = exploer{}
		var data = []byte{}
		var exists = false
		data, exists = cache.Get(item.URL)
	
		if !exists{
			res, err := http.Get(item.URL)
	
			if err != nil {
				log.Fatal(err)
			}
	
			data, err = io.ReadAll(res.Body)
	
			res.Body.Close()
			cache.Add(item.URL, data)
	
			if res.StatusCode > 299 {
				log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, data)
			}
			if err != nil {
				log.Fatal(err)
			}
		}

		err := json.Unmarshal(data, &arearExplore)
		if err != nil {
			fmt.Println(err)
		}

		for _, item := range arearExplore.PokemonEncounters{
			fmt.Println(" - " + item.Pokemon.Name)
		}
	}

	return nil
}

