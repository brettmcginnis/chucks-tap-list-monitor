package site

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Site string

const (
	GreenWood Site = "GW"
)

func (s Site) List() []beerJson {
	var beers []beerJson

	err := json.Unmarshal([]byte(s.fetchTaps()), &beers)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
	}

	return beers
}

func (s Site) fetchTaps() string {
	url := fmt.Sprintf("https://taplists.web.app/data?menu=%s", s)

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error making the request:", err)
		return ""
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading the response:", err)
		return ""
	}

	return string(body)
}
