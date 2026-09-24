package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocations(pageURL *string) (RespShallowLocations, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	var data []byte

	d, found := c.cache.Get(url)
	if found != true {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return RespShallowLocations{}, err
		}

		resp, err := c.httpClient.Do(req)
		if err != nil {
			return RespShallowLocations{}, err
		}

		defer resp.Body.Close()

		if resp.StatusCode > 299 || resp.StatusCode < 200 {
			return RespShallowLocations{}, fmt.Errorf("Failed to get location area: HTTP status %d", resp.StatusCode)
		}

		data, err = io.ReadAll(resp.Body)
		if err != nil {
			return RespShallowLocations{}, err
		}

		c.cache.Add(url, data)

	} else {
		data = d
	}

	locationsResp := RespShallowLocations{}
	err := json.Unmarshal(data, &locationsResp)
	if err != nil {
		return RespShallowLocations{}, err
	}

	return locationsResp, nil
}
