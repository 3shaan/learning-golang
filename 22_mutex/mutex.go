package main

import (
	"fmt"
	"sync"
)

type Post struct {
	view int
	mu   sync.Mutex
}

func (p *Post) inc(wg *sync.WaitGroup) {
	defer func() {
		p.mu.Unlock()
		wg.Done()
	}()
	p.mu.Lock()
	p.view++
}

func main() {
	post := Post{
		view: 0,
	}

	wg := sync.WaitGroup{}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go post.inc(&wg)
	}

	wg.Wait()

	fmt.Println("Post view", post.view)

}
