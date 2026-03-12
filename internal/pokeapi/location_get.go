package pokeapi

import(
	"encoding/json"
	"io"
	"net/http"
	"fmt"
)

func (c *Client) GetLocationDetails(locationName string) (LocationDetails, error) {
	url := baseURL + "/location-area/" + locationName

	if val, ok := c.cache.Get(url); ok {
		location := LocationDetails{}
		err := json.Unmarshal(val, &location)
		if err != nil {
			return LocationDetails{}, err
		}
		return location, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationDetails{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationDetails{}, err
	}

	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationDetails{}, fmt.Errorf("error reading Location-Area data: %v", err)
	}

	location := LocationDetails{}
	if err := json.Unmarshal(data, &location); err != nil {
		return LocationDetails{}, fmt.Errorf("Error unmarshaling Location-Area data: %s", err)
	}

	c.cache.Add(url, data)
	return location, nil
}