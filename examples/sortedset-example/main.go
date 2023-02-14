package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/momentohq/client-sdk-go/auth"
	"github.com/momentohq/client-sdk-go/config"
	"github.com/momentohq/client-sdk-go/momento"
)

const (
	cacheName = "test-cache"
	setName   = "my-set"
)

func main() {
	// Create Momento client
	client := getClient()
	ctx := context.Background()

	// Create cache
	setupCache(client, ctx)

	// Put score for each element to set
	// Using counter, element N has score N
	for i := 1; i < 11; i++ {
		_, err := client.SortedSetPut(ctx, &momento.SortedSetPutRequest{
			CacheName: cacheName,
			SetName:   setName,
			Elements: []*momento.SortedSetScoreRequestElement{{
				Name:  momento.StringBytes{Text: fmt.Sprintf("element-%d", i)},
				Score: float64(i),
			}},
		})
		if err != nil {
			panic(err)
		}
	}

	// Fetch sorted set
	fmt.Println("\n\nFetching all elements from sorted set:")
	fmt.Println("--------------")
	fetchResp, err := client.SortedSetFetch(ctx, &momento.SortedSetFetchRequest{
		CacheName: cacheName,
		SetName:   setName,
	})
	if err != nil {
		panic(err)
	}

	displayElements(setName, fetchResp)

	// Fetch top 5 elements in descending order (high -> low)
	fmt.Println("\n\nFetching Top 5 elements from sorted set:")
	fmt.Println("--------------")
	top5Rsp, err := client.SortedSetFetch(ctx, &momento.SortedSetFetchRequest{
		CacheName:       cacheName,
		SetName:         setName,
		NumberOfResults: momento.FetchLimitedElements{Limit: 5},
		Order:           momento.DESCENDING,
	})
	if err != nil {
		panic(err)
	}

	displayElements(setName, top5Rsp)
}

func getClient() momento.ScsClient {
	credProvider, err := auth.NewEnvMomentoTokenProvider("MOMENTO_AUTH_TOKEN")
	if err != nil {
		panic(err)
	}
	client, err := momento.NewSimpleCacheClient(&momento.SimpleCacheClientProps{
		Configuration:      config.LatestLaptopConfig(),
		CredentialProvider: credProvider,
		DefaultTTL:         60 * time.Second,
	})
	if err != nil {
		panic(err)
	}
	return *client
}

func setupCache(client momento.ScsClient, ctx context.Context) {
	err := client.CreateCache(ctx, &momento.CreateCacheRequest{
		CacheName: "test-cache",
	})
	if err != nil {
		var momentoErr momento.MomentoError
		if errors.As(err, &momentoErr) {
			if momentoErr.Code() != momento.AlreadyExistsError {
				panic(err)
			}
		}
	}
}

func displayElements(setName string, resp momento.SortedSetFetchResponse) {
	switch r := resp.(type) {
	case momento.SortedSetFetchHit:
		for _, e := range r.Elements {
			fmt.Printf("setName: %s, elementName: %s, score: %f\n", setName, e.Name, e.Score)
		}
		fmt.Println("")
	case momento.SortedSetFetchMiss:
		fmt.Println("we regret to inform you there is no such set")
		os.Exit(1)
	}
}