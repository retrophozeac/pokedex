package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListLocations(pageurl *string) (RespShallowLocations, error) {
	url := baseURL + "/location-area"
	if pageurl != nil {
		url = *pageurl
	}
	if val, ok := c.cache.Get(url); ok {
		locationResp := RespShallowLocations{}
		err := json.Unmarshal(val, &locationResp)
		if err != nil {
			return RespShallowLocations{}, err
		}
		// fmt.Printf("using cache")
		return locationResp, nil
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespShallowLocations{}, err
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespShallowLocations{}, nil
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespShallowLocations{}, err
	}
	locations := RespShallowLocations{}
	err = json.Unmarshal(dat, &locations)
	if err != nil {
		return RespShallowLocations{}, err
	}
	c.cache.Add(url, dat)
	return locations, nil
}
