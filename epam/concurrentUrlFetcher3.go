package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type something3 struct {
	Url  string
	Body string
}

func ConcurrentURLFetcher3(ctx context.Context, urls []string) (m map[string]string) {
	m = make(map[string]string)
	l := len(urls)
	c := make(chan something3, l)
	wg := sync.WaitGroup{}
	wg.Add(l)

	for _, v := range urls {
		go func(ctx context.Context, url string) {
			defer wg.Done()
			body := hitUrl3(ctx, url)
			c <- something3{
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

// this execution is more randomized because the infinite loop for body is unexpected value when it is returning, sometime is not ready to and the map is empty
func hitUrl3(ctx context.Context, url string) (s string) {
	s += "body"
	for i := 1; ; i++ {
		select {
		case <-ctx.Done():
			return s
		default:
			time.Sleep(time.Millisecond * 1)
			s += fmt.Sprint(" ", i)
		}
	}

	return fmt.Sprintf("body from %s", url)
}

// execute it multiple times
func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel()

	inputs := []string{"https://example.com", "https://google.com", "https://openai.com"}
	fmt.Println(ConcurrentURLFetcher3(ctx, inputs))
}
