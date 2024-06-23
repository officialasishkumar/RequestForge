package main

import (
	"RequestForge/httpclient"
	"RequestForge/logger"
	"RequestForge/models"
	"RequestForge/requests"
	"flag"
	"fmt"
	"sync"
	"time"
)

func main() {
	startTime := time.Now()

	jsonFilePath := flag.String("f", "", "Path to the JSON file containing requests")
	concurrent := flag.Bool("concurrent", false, "Enable concurrent processing of requests")
	verbose := flag.Bool("v", false, "Enable verbose output")
	flag.Parse()

	logger.Init()

	if *jsonFilePath == "" {
		fmt.Println("Please provide the path to the JSON file using -f flag.")
		return
	}

	reqs, err := requests.LoadRequests(*jsonFilePath)
	if err != nil {
		panic(err)
	}

	client := httpclient.NewClient(10*time.Second, *verbose)

	if *concurrent {
		processRequestsConcurrently(reqs, client)
	} else {
		processRequestsSequentially(reqs, client)
	}

	elapsedTime := time.Since(startTime)
	fmt.Printf("Executed in %.3f seconds\n", elapsedTime.Seconds())
}

func processRequestsSequentially(requests []models.Request, client *httpclient.Client) {
	for _, reqData := range requests {
		client.SendRequest(reqData)
	}
}

func processRequestsConcurrently(requests []models.Request, client *httpclient.Client) {
	var wg sync.WaitGroup
	for _, reqData := range requests {
		wg.Add(1)
		go func(reqData models.Request) {
			defer wg.Done()
			client.SendRequest(reqData)
		}(reqData)
	}
	wg.Wait()
}
