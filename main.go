package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type ShitPost struct {
	Title string `json:"data.children[0].data.title"` // TODO: This doesn't work, probably need to change the struct
	Text string `json:"data.children[0].data.selftext"` // TODO: This doesn't work, probably need to change the struct
	Kind string `json:"kind"` // This works fine
}

func main() {

	httpclient := &http.Client{}

	req, err := http.NewRequest("GET", "https://reddit.com/r/copypasta/top/.json?sort=top&t=week&limit=1", nil)

	if err != nil {
		log.Println("error:", err)
		return
	}

	req.Header.Add("User-Agent", "Meme-cli")

	res, err := httpclient.Do(req)

	if err != nil {
		log.Println("error:", err)
		return
	}
	if res.StatusCode != 200 {
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