package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListLocations(pageURL *string) (ResShallowLocations, error) {
	locations := ResShallowLocations{}
	url := baseURL + "/location-area/"
	if pageURL != nil {
		url = *pageURL
	}

	if val, exists := c.cache.Get(url); exists {
		err := json.Unmarshal(val, &locations)
		if err != nil {
			return locations, err
		}
		return locations, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return locations, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return locations, err
	}
	defer res.Body.Close()

	dat, err := io.ReadAll(res.Body)
	if err != nil {
		return locations, err
	}

	err = json.Unmarshal(dat, &locations)
	if err != nil {
		return locations, err
	}

	c.cache.Add(url, dat)
	return locations, nil
}
