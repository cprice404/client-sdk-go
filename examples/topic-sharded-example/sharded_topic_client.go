package main

import (
	"context"
	"fmt"
	"github.com/momentohq/client-sdk-go/auth"
	"github.com/momentohq/client-sdk-go/config"
	"github.com/momentohq/client-sdk-go/config/logger"
	"github.com/momentohq/client-sdk-go/momento"
	"github.com/momentohq/client-sdk-go/responses"
	"math/rand"
)

type shardedTopicClient struct {
	topicClient       momento.TopicClient
	numShardsPerTopic int
	log               logger.MomentoLogger
}

func NewShardedTopicClient(topicsConfiguration config.TopicsConfiguration, credentialProvider auth.CredentialProvider, numShardsPerTopic int) (momento.TopicClient, error) {
	topicClient, err := momento.NewTopicClient(topicsConfiguration, credentialProvider)
	if err != nil {
		return nil, err
	}
	return shardedTopicClient{
		topicClient:       topicClient,
		numShardsPerTopic: numShardsPerTopic,
		log:               topicsConfiguration.GetLoggerFactory().GetLogger("sharded-topic-client"),
	}, nil
}

func (s shardedTopicClient) Subscribe(ctx context.Context, request *momento.TopicSubscribeRequest) (momento.TopicSubscription, error) {
	return s.topicClient.Subscribe(ctx, request)
}

func (s shardedTopicClient) Publish(ctx context.Context, request *momento.TopicPublishRequest) (responses.TopicPublishResponse, error) {
	topicNamePrefix := request.TopicName
	shardToPublishTo := rand.Intn(s.numShardsPerTopic)
	topicName := fmt.Sprintf("%s-%d", topicNamePrefix, shardToPublishTo)
	shardRequest := momento.TopicPublishRequest{
		CacheName: cacheName,
		TopicName: topicName,
		Value:     request.Value,
	}
	return s.topicClient.Publish(ctx, &shardRequest)
}

func (s shardedTopicClient) Close() {
	s.topicClient.Close()
}
