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
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationStructResp{}, fmt.Errorf("error reading data err : %v", err)
	}
	//unmashal the data to  locationAreas LocationStructResp
	locationAreas := LocationStructResp{}
	err = json.Unmarshal(data, &locationAreas)
	if err != nil {
		return LocationStructResp{}, fmt.Errorf("error unmashalling data err : %v", err)
	}
	//return datas
	return locationAreas, nil
}
