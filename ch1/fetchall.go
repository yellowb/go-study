package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	urls := []string{"http://www.baidu.com", "http://www.163.com", "http://www.zhihu.com"}
	start := time.Now()

	ch := make(chan string)
	for _, url := range urls {
		go fetch(url, ch)
	}
	for range urls {
		fmt.Println(<-ch)
	}

	seconds := time.Since(start).Seconds()
	fmt.Printf("All urls finished. Time: %.2fs \n", seconds)
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	seconds := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %dbytes %s", seconds, nbytes, url)
}
