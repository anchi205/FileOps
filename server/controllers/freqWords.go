package controllers

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"sync"

	"github.com/gin-gonic/gin"
)

func freqWordsHandler(c *gin.Context) {
	sortOrder := c.Query("sortOrder")
	limit := c.Query("limit")
	searchTopWords(c, sortOrder, limit)
}

func searchTopWords(c *gin.Context, sortOrder string, limit string) {
	files, err := os.ReadDir("./uploads")
	if err != nil {
		c.String(http.StatusInternalServerError, "Error reading directory")
		return
	}

	sortDescending := true
	maxWords := 10

	if sortOrder == "asc" {
		sortDescending = false
	}
	if limit != "" {
		value, err := strconv.Atoi(limit)
		if err == nil && value > 0 {
			maxWords = value
		}
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

	sort.Slice(topWords, func(i, j int) bool {
		if sortDescending {
			return topWords[i].Count > topWords[j].Count
		}
		return topWords[i].Count < topWords[j].Count
	})

	if len(topWords) > maxWords {
		topWords = topWords[:maxWords]
	}

	var message strings.Builder
	for _, wordCount := range topWords {
		message.WriteString(fmt.Sprintf("%s: %d\n", wordCount.Word, wordCount.Count))
	}

	c.String(http.StatusOK, message.String())
}
