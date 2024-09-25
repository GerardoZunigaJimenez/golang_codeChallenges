package epam

import (
	"context"
	"fmt"
	"sync"
)

type something struct {
	Url  string
	Body string
}

func ConcurrentURLFetcher(ctx context.Context, urls []string) (m map[string]string) {
	m = make(map[string]string)
	c := make(chan something)
	wg := sync.WaitGroup{}
	l := len(urls)
	wg.Add(l)

	for _, v := range urls {
		go func(ctx context.Context, url string) {
			defer wg.Done()
			body := hitUrl(ctx, url)
			c <- something{
				v,
				body,
			}
		}(ctx, v)
	}

	go func() {
		wg.Wait()
		close(c)
	}()

	for v := range c {
		m[v.Url] = v.Body
	}

	return m
}

func hitUrl(ctx context.Context, url string) string {
	return fmt.Sprintf("body from %s", url)
}

func main() {
	ctx := context.Background()

	inputs := []string{"https://example.com", "https://google.com", "https://openai.com"}
	fmt.Println(ConcurrentURLFetcher(ctx, inputs))
}
