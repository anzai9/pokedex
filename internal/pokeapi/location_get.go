package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetLocation(name string) (ResShallowLocationAreas, error) {
	locationAreas := ResShallowLocationAreas{}
	url := baseURL + "/location-area/" + name

	if val, exists := c.cache.Get(url); exists {
		err := json.Unmarshal(val, &locationAreas)
		if err != nil {
			return locationAreas, err
		}
		return locationAreas, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return locationAreas, nil
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return locationAreas, nil
	}
	defer res.Body.Close()

	dat, err := io.ReadAll(res.Body)
	if err != nil {
		return locationAreas, nil
	}

	err = json.Unmarshal(dat, &locationAreas)
	if err != nil {
		return locationAreas, nil
	}

	c.cache.Add(url, dat)
	return locationAreas, nil
}
