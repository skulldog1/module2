package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func joinChannels(chs ...<-chan int) chan int {
	mergedCh := make(chan int)

	go func() {
		wg := &sync.WaitGroup{}

		wg.Add(len(chs))

		for _, ch := range chs {
			go func(ch <-chan int, wg *sync.WaitGroup) {
				defer wg.Done()
				for id := range ch {
					mergedCh <- id
				}
			}(ch, wg)
		}

		wg.Wait()
		close(mergedCh)
	}()

	return mergedCh
}

func generateData() chan int {
	out := make(chan int, 1000)

	go func() {
		defer close(out)
		for {
			select {
			case _, ok := <-out:
				if !ok {
					return
				}
			case out <- rand.Intn(100):
			}
		}
	}()

	return out
}

func main() {
	rand.Seed(time.Now().UnixNano())
	a := make(chan int)
	b := make(chan int)
	c := make(chan int)

	out := generateData()

	go func() {
		for num := range out {
			a <- num
		}
		close(a)
	}()

	go func() {
		for num := range out {
			b <- num
		}
		close(b)
	}()

	go func() {
		for num := range out {
			c <- num
		}
		close(c)
	}()

	mainChan := joinChannels(a, b, c)

	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()

	done := make(chan struct{})

	go func() {
		for num := range mainChan {
			fmt.Println(num)
		}
		close(done)
	}()

	for {
		select {
		case <-done:
			fmt.Println("all channels are closed")
			return
		case <-ticker.C:
			fmt.Println("timeout, closing all channels")
			close(a)
			close(b)
			close(c)
			return
		}
	}
}
