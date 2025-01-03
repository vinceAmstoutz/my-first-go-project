package main

// Exercise available here: https://go.dev/tour/concurrency/10
// Instructions:
// 1. Fetch the URLs concurrently using goroutines to improve performance.
// 2. Ensure that no URL is fetched more than once by keeping track of the URLs that have already been processed.
// The goal is to practice concurrency in Go by using goroutines to fetch URLs in parallel without duplication.

import (
	"fmt"
	"sync"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher, visited map[string]bool, mu *sync.Mutex, wg *sync.WaitGroup) {
	defer wg.Done()

	if depth <= 0 || visited[url] {
		return
	}

	mu.Lock()
	visited[url] = true
	mu.Unlock()

	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("found: %s %q\n", url, body)

	for _, u := range urls {
		wg.Add(1)
		Crawl(u, depth-1, fetcher, visited, mu, wg)
	}
}

func main() {
	var (
		mu      sync.Mutex
		wg      sync.WaitGroup
		visited = make(map[string]bool)
	)

	fetcher := createFetcher()

	wg.Add(1)
	Crawl("https://golang.org/", 4, fetcher, visited, &mu, &wg)
	wg.Wait()
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}

	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
func createFetcher() Fetcher {
	return fakeFetcher{
		"https://golang.org/": &fakeResult{
			"The Go Programming Language",
			[]string{
				"https://golang.org/pkg/",
				"https://golang.org/cmd/",
			},
		},
		"https://golang.org/pkg/": &fakeResult{
			"Packages",
			[]string{
				"https://golang.org/",
				"https://golang.org/cmd/",
				"https://golang.org/pkg/fmt/",
				"https://golang.org/pkg/os/",
			},
		},
		"https://golang.org/pkg/fmt/": &fakeResult{
			"Package fmt",
			[]string{
				"https://golang.org/",
				"https://golang.org/pkg/",
			},
		},
		"https://golang.org/pkg/os/": &fakeResult{
			"Package os",
			[]string{
				"https://golang.org/",
				"https://golang.org/pkg/",
			},
		},
	}
}
