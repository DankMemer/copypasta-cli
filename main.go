package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"time"
	"github.com/danielchatfield/go-chalk"
)

type ShitPost struct {
	Data struct {
		Children []struct {
			Post struct {
				Title string `json:"title"`
				Text string `json:"selftext"`
			} `json:"data"`
		} `json:"children"`
	} `json:"data"`
}

func main() {

	httpclient := &http.Client{}

	req, err := http.NewRequest("GET", "https://www.reddit.com/r/copypasta/top/.json?sort=top&t=week&showmedia=false&mediaonly=false&is_self=true&limit=100", nil)

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

	s, err := getData([]byte(body)) // turn the JSON response into a struct

	rand.Seed(time.Now().UnixNano()) // Seed the random number
	randomPost := rand.Intn(100 - 1) + 1 // TODO: make the first number (max) depend on how many "Children" structs are returned

	post := s.Data.Children[randomPost].Post
	fmt.Println(chalk.Yellow(post.Title), "\n",
	chalk.Blue(post.Text))
}

func getData(body []byte) (*ShitPost, error) {
	var s = new(ShitPost)
	err := json.Unmarshal(body, &s)
	if err != nil {
		log.Println("error:", err)
	}
	return s, err
}
