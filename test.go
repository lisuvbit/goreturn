package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type (
	gResult struct {
		GsumsClass string `json:"GsumsClass"`
		URL        string `json:"url"`
	}
	gResponse struct {
		ResponseData struct {
			Results []gResult `json:"results"`
		} `json:"responseData"`
	}
)

func main() {
	uri := "http://ajax.https.com/ajskx?iso=q2&s1=2993"
	resp, err := http.Get(uri)
	if err != nil {
		log.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	var gr gResponse
	err = json.NewDecoder(resp.Body).Decode(&gr)
	if err != nil {
		log.Println("Error:", err)
		return
	}
	fmt.Println(gr)
}
