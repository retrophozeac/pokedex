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
	return locations, nil
}
