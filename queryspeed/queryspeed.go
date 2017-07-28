package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

type stat struct {
	url        string
	statusCode int
	speed      time.Duration
}

func main() {

	reader := bufio.NewReader(os.Stdin)
	b, err := ioutil.ReadAll(reader)
	if err != nil {
		panic(err)
	}

	urls := strings.Split(strings.TrimSpace(string(b)), " ")
	c := make(chan stat, len(urls))
	for _, url := range urls {
		go func(uri string) {
			statusCode, speed := Get(uri)
			c <- stat{url: uri, statusCode: statusCode, speed: speed}
		}(url)
	}

	select {
	case stat := <-c:
		fmt.Printf("\nGET %s. Received statusCode: %d. Took %d.\n", stat.url, stat.statusCode, stat.speed)
	case <-time.After(time.Second * 1500):
		fmt.Println("timeout.")
		return
	}
}

func Get(url string) (int, time.Duration) {
	start := time.Now()
	resp, err := http.Get(url)
	fmt.Printf("url: %#v", url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	fmt.Printf("\nGET %s. Received statusCode: %d. Took %d.\n", url, resp.StatusCode, time.Since(start))
	return resp.StatusCode, time.Since(start)
}
