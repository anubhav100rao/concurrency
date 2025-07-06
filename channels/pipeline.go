package channels

import (
	"fmt"
	"strings"
)

// stage 1, splitting each word
func split(texts <-chan string) <-chan []string {
	out := make(chan []string)
	go func() {
		defer close(out)
		for text := range texts {
			out <- strings.Fields(text)
		}
	}()
	return out
}

// stage 2, capitalising each word
func uppercase(in <-chan []string) <-chan []string {
	out := make(chan []string)
	go func() {
		defer close(out)
		for words := range in {
			for i, w := range words {
				words[i] = strings.ToUpper(w)
			}
			out <- words
		}
	}()
	return out
}

func RunPipleline() {
	texts := make(chan string)
	go func() {
		defer close(texts)
		texts <- "hello world"
		texts <- "go channels rock"
		texts <- "smart work"
		texts <- "go concurrency building blocks"
		texts <- "1. goroutings"
		texts <- "2. sync package"
		texts <- "3. channels"
		texts <- "4. select block"
	}()

	for words := range uppercase(split(texts)) {
		fmt.Println(words)
	}
}
