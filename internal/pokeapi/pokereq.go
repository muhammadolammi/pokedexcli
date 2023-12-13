package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) MakeRequest(pageUrl *string) (LocationStructResp, error) {
	url := baseUrl
	if pageUrl != nil {
		url = *pageUrl
	}
	//check if url is in cache
	data, ok := c.cache.Get(url)
	//if the cache is available unmarhsal it and use it as the data, this will return early
	if ok {
		fmt.Println("Cache Hit!")
		locationAreas := LocationStructResp{}
		err := json.Unmarshal(data, &locationAreas)
		if err != nil {
			return LocationStructResp{}, fmt.Errorf("error unmashalling data err : %v", err)
		}
		//return datas
		return locationAreas, nil

	}
	fmt.Println("Cache miss!")

	// if no cache we make a request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationStructResp{}, err
	}
	res, err := c.httpClient.Do(req)
	if err != nil {
		return LocationStructResp{}, err
	}
	defer res.Body.Close()
	//check for status code error
	if res.StatusCode > 399 {
		return LocationStructResp{}, fmt.Errorf("status code : %v", res.StatusCode)
	}
	// get the data (json string) from response
	data, err = io.ReadAll(res.Body)
	if err != nil {
		return LocationStructResp{}, fmt.Errorf("error reading data err : %v", err)
	}
	//unmashal the data to  locationAreas LocationStructResp
	locationAreas := LocationStructResp{}
	err = json.Unmarshal(data, &locationAreas)
	if err != nil {
		return LocationStructResp{}, fmt.Errorf("error unmashalling data err : %v", err)
	}

	//we cache the data we get
	c.cache.Add(url, data)
	//return datas

	return locationAreas, nil
}
