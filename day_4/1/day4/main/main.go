//////////////////////////////////////////////////////////////////////
//
// Given is a producer-consumer szenario, where a producer reads in
// tweets from a mockstream and a consumer is processing the
// data. Your task is to change the code so that the producer as well
// as the consumer can run concurrently
//

package main

import (
	"fmt"
	"sync"
	"time"
)

func producer(stream Stream, ch chan<- *Tweet) {
	// defer()
	for {
		tweet, err := stream.Next()
		if err == ErrEOF {
			close(ch)
			return
		}

		ch <- tweet
	}
}

func consumer(tweets <-chan *Tweet, wg *sync.WaitGroup) {
	defer wg.Done()
	for t := range tweets {
		if t.IsTalkingAboutGo() {
			fmt.Println(t.Username, "\ttweets about golang")
		} else {
			fmt.Println(t.Username, "\tdoes not tweet about golang")
		}
	}
}

func main() {
	start := time.Now()
	stream := GetMockStream()

	ch := make(chan *Tweet, 1)

	var wg sync.WaitGroup
	// Consumer
	wg.Add(1)
	go consumer(ch, &wg)

	// Producer
	producer(stream, ch)

	wg.Wait()
	fmt.Printf("Process took %s\n", time.Since(start))
}
