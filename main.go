package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	httpclient := &http.Client{}
	req, err := http.NewRequest("GET", "https://www.reddit.com/r/copypasta/top/.json?sort=top&t=week&limit=1", nil)

	if err != nil {
		log.Println("error:", err)
		return
	}

	res, err := httpclient.Do(req)

	if err != nil {
		log.Println("error:", err)
		return
	}
	fmt.Println(res.Body)
	defer res.Body.Close()
}
