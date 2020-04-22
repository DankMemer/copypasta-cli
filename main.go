package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type ShitPost struct {
	Title string `json:"data.children[0].data.title"`
	Text string `json:"data.children[0].data.selftext"`
	Message string `json:"message"`
}

func main() {
	res, err := http.Get("https://reddit.com/r/copypasta/top/.json?sort=top&t=week&limit=1")
	if err != nil {
		log.Println("error:", err)
		return
	}
	if res.StatusCode != 200 { // TODO: Why are we getting 429 after just a few reqs, not even close to the RL
		log.Fatalln("Non-OK Status:", res.Status)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println("error:", err)
		return
	}

	s, err := getData([]byte(body))
	log.Println(s)
}

func getData(body []byte) (*ShitPost, error) {
	var s = new(ShitPost)
	err := json.Unmarshal(body, &s)
	if err != nil {
		log.Println("error:", err)
	}
	return s, err
}