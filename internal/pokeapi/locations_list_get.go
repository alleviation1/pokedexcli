package pokeapi

import(
	"encoding/json"
	"io"
	"net/http"
	"fmt"
)

// Get 20 Locations
func (c *Client) GetLocations(pageURL *string) (Locations, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	// attempt to find in cache
	if cacheResult, ok := c.cache.Get(url); ok {
		locations := Locations{}
		err := json.Unmarshal(cacheResult, &locations)
		if err != nil {
			return Locations{}, fmt.Errorf("Error unmarshaling cached Location-Area data: %v", err)
		}

		return locations, nil
	}
	
	// default http call if data is not cached
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Locations{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Locations{}, err
	}

	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return Locations{}, fmt.Errorf("Error reading Location-Area data: %v", err)
	}

	locations := Locations{}
	if err := json.Unmarshal(data, &locations); err != nil {
		return Locations{}, fmt.Errorf("Error unmarshaling Location-Area data: %s", err)
	}

	c.cache.Add(url, data)
	return locations, nil
}