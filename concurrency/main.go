package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
)

func main() {
	filePath := "sample.txt"
	wordCount := countWordsConcurrent(filePath)
	for word, count := range wordCount {
		fmt.Printf("%s: %d\n", word, count)
	}
}

func countWordsConcurrent(filePath string) map[string]int {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return nil
	}
	defer file.Close()

	wordCount := make(map[string]int)
	var mu sync.Mutex
	var wg sync.WaitGroup
	wordChannel := make(chan string)

	go func() {
		scanner := bufio.NewScanner(file)
		scanner.Split(bufio.ScanWords)
		for scanner.Scan() {
			wordChannel <- scanner.Text()
		}
		close(wordChannel)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for word := range wordChannel {
			mu.Lock()
			wordCount[strings.ToLower(word)]++
			mu.Unlock()
		}
	}()

	wg.Wait()
	return wordCount
}
