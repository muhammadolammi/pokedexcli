package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ExploreArea(area string) (AreaExploreStructResp, error) {
	const base = "https://pokeapi.co/api/v2/location-area/"
	url := base + area
	fmt.Printf("Exploring %v...\n", area)
	//check if url is in cache
	data, ok := c.cache.Get(url)
	//if the cache is available unmarhsal it and use it as the data, this will return early
	if ok {
		// i can write cache hit handler in future

		locationAreas := AreaExploreStructResp{}
		err := json.Unmarshal(data, &locationAreas)
		if err != nil {
			return AreaExploreStructResp{}, fmt.Errorf("error unmashalling data err : %v", err)
		}
		//return datas
		return locationAreas, nil

	}

	// if no cache we make a request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return AreaExploreStructResp{}, err
	}
	res, err := c.httpClient.Do(req)
	if err != nil {
		return AreaExploreStructResp{}, err
	}
	defer res.Body.Close()
	//check for status code error
	if res.StatusCode > 399 {
		return AreaExploreStructResp{}, fmt.Errorf("status code : %v", res.StatusCode)
	}
	// get the data (json string) from response
	data, err = io.ReadAll(res.Body)
	if err != nil {
		return AreaExploreStructResp{}, fmt.Errorf("error reading data err : %v", err)
	}
	//unmashal the data to  locationAreas LocationStructResp
	locationAreas := AreaExploreStructResp{}
	err = json.Unmarshal(data, &locationAreas)
	if err != nil {
		return AreaExploreStructResp{}, fmt.Errorf("error unmashalling data err : %v", err)
	}

	//we cache the data we get
	c.cache.Add(url, data)
	//return datas

	return locationAreas, nil
}
