package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type something2 struct {
	Url  string
	Body string
}

func ConcurrentURLFetcher2(ctx context.Context, urls []string) (m map[string]string) {
	m = make(map[string]string)
	l := len(urls)
	c := make(chan something2, l)
	wg := sync.WaitGroup{}
	wg.Add(l)

	for _, v := range urls {
		go func(ctx context.Context, url string) {
			defer wg.Done()
			body := hitUrl2(ctx, url)
			c <- something2{
				v,
				body,
			}
		}(ctx, v)
	}

	wg.Wait()
	for {
		select {
		case s := <-c:
			m[s.Url] = s.Body
		case <-ctx.Done():
			fmt.Println("goodbye, the content of the channel is ", len(c))
			return m
		}
	}

	return m
}

func hitUrl2(ctx context.Context, url string) (s string) {
	return "body"
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel()

	inputs := []string{"https://example.com", "https://google.com", "https://openai.com"}
	fmt.Println(ConcurrentURLFetcher2(ctx, inputs))
}
