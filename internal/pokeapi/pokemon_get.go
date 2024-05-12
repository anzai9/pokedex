package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetPokemon(name string) (Pokemon, error) {
	pokemon := Pokemon{}
	url := baseURL + "/pokemon/" + name

	if val, exists := c.cache.Get(url); exists {
		err := json.Unmarshal(val, &pokemon)
		if err != nil {
			return pokemon, err
		}
		return pokemon, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return pokemon, nil
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return pokemon, nil
	}
	defer res.Body.Close()

	dat, err := io.ReadAll(res.Body)
	if err != nil {
		return pokemon, nil
	}

	if err := json.Unmarshal(dat, &pokemon); err != nil {
		return pokemon, nil
	}
	c.cache.Add(url, dat)

	return pokemon, nil
}
