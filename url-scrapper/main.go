package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"
)

var InputFilePathEmpty = errors.New("Input file path is empty")

const (
	Error   = "Error"
	Timeout = "Timeout"
)

type Result struct {
	Url      string
	Status   string
	Duration string
	Error    string
}

type Config struct {
	workers    int
	timeout    time.Duration
	inputFile  string
	outputFile string
}

func main() {
	config := getConfig()

	urls, err := getUrls(config.inputFile)
	if err != nil {
		log.Fatalf("Failed to get URLs: %v", err)
		return
	}

	results := checkUrls(urls, config)

	if err := saveResults(results, config); err != nil {
		log.Fatalf("Failed to save results: %v", err)
	}
}

func getConfig() *Config {
	return &Config{
		workers:    10,
		timeout:    5 * time.Second,
		inputFile:  "urls.txt",
		outputFile: "output.json",
	}
}

func getUrls(filepath string) ([]string, error) {
	if filepath != "" {
		content, err := os.ReadFile(filepath)
		if err != nil {
			return nil, err
		}
		var urls []string
		for _, line := range strings.Split(string(content), "\n") {
			line = strings.TrimSpace(line)
			if line != "" {
				urls = append(urls, line)
			}
		}
		return urls, nil
	}
	return nil, InputFilePathEmpty
}

func checkUrls(urls []string, config *Config) []Result {

	//why buffered urlChannel ? channel wouldn't require an immediate reciver!!
	urlCh := make(chan string, len(urls))

	for _, url := range urls {
		urlCh <- url
	}
	close(urlCh) //closing indicates - no more items to send

	resultsCh := make(chan Result)

	var wg sync.WaitGroup //this tells when all workers are done

	for i := 0; i < config.workers; i++ {
		wg.Add(1)
		go work(urlCh, resultsCh, config.timeout, &wg)
	}
	var results []Result
	var resultsLock sync.Mutex

	//why do we need another go routine for collecting results ???
	/*
		if we recieve results after wg.wait(), the workers might be blocked trying to send results to resultsCh with no one to read from it- DEADLOCK!
	*/
	go func() {
		for result := range resultsCh {
			//critical section :- updating shared variable
			resultsLock.Lock()
			results = append(results, result)
			resultsLock.Unlock()
		}

	}()

	wg.Wait()
	close(resultsCh)

	return results
}

func work(urlCh <-chan string, resultsCh chan<- Result, timeout time.Duration, wg *sync.WaitGroup) {
	defer wg.Done()
	client := &http.Client{}

	for url := range urlCh {
		result := pingUrl(client, url, timeout)
		resultsCh <- result
	}

}

func pingUrl(client *http.Client, url string, timeout time.Duration) Result {
	startTime := time.Now()

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return formatResponse(url, Error, 0*time.Second, err)
	}
	// Send request
	resp, err := client.Do(req)
	duration := time.Since(startTime)
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			return formatResponse(url, Timeout, duration, fmt.Errorf("request timed out"))
		}
		return formatResponse(url, Error, duration, err)
	}
	defer resp.Body.Close()

	// Return result
	return formatResponse(url, fmt.Sprintf("%d", resp.StatusCode), duration, nil)

}

func formatResponse(url, status string, duration time.Duration, err error) Result {
	var errStr string
	if err != nil {
		errStr = err.Error()
	}
	return Result{
		Url:      url,
		Status:   status,
		Duration: duration.String(),
		Error:    errStr,
	}
}

func saveResults(results []Result, config *Config) error {
	file, err := os.Create(config.outputFile)
	if err != nil {
		return err
	}
	defer file.Close()
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(results)
}
