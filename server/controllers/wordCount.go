package controllers

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/gin-gonic/gin"
)

type WordCount struct {
	Word  string
	Count int
}

func WordCountHandler(c *gin.Context) {
	if c.Request.Method == "GET" {
		searchAllWords(c)
	} else {
		c.String(http.StatusMethodNotAllowed, "Method not allowed")
	}
}

func searchAllWords(c *gin.Context) {
	files, err := os.ReadDir("./uploads")
	if err != nil {
		c.String(http.StatusInternalServerError, "Error reading directory")
		return
	}

	wordCounts := make(map[string]int)
	var mu sync.Mutex
	var wg sync.WaitGroup

	for _, file := range files {
		if file.IsDir() {
			continue
		}
		wg.Add(1)
		go func(filename string) {
			defer wg.Done()
			path := filepath.Join(filepath.Join("./uploads", filename))

			file, err := os.Open(path)
			if err != nil {
				fmt.Printf("Failed to open file %s: %s\n", path, err)
				return
			}
			defer file.Close()

			scanner := bufio.NewScanner(file)
			scanner.Split(bufio.ScanWords)

			for scanner.Scan() {
				word := strings.ToLower(scanner.Text())
				mu.Lock()
				wordCounts[word]++
				mu.Unlock()
			}

			if err := scanner.Err(); err != nil {
				fmt.Printf("Error scanning file %s: %s\n", path, err)
			}
		}(file.Name())
	}

	wg.Wait()

	topWords := make([]WordCount, 0, len(wordCounts))
	for word, count := range wordCounts {
		topWords = append(topWords, WordCount{Word: word, Count: count})
	}

	totalWordCount := 0
	for _, wordCount := range topWords {
		totalWordCount += wordCount.Count
	}
	message := fmt.Sprintf("Total word count: %d\n", totalWordCount)
	c.String(http.StatusOK, message)
}
