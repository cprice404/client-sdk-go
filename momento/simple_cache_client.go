// Package momento represents API ScsClient interface accessors including control/data operations, errors, operation requests and responses for the SDK.
package momento

import (
	"context"
	"fmt"

	"github.com/momentohq/client-sdk-go/auth"
	"github.com/momentohq/client-sdk-go/config"
	"github.com/momentohq/client-sdk-go/internal/models"
	"github.com/momentohq/client-sdk-go/internal/momentoerrors"
	"github.com/momentohq/client-sdk-go/internal/services"
	"github.com/momentohq/client-sdk-go/internal/utility"
)

// ScsClient wraps lower level cache control and data operations.
type ScsClient interface {
	// CreateCache Create a new cache in your Momento account.
	CreateCache(ctx context.Context, request *CreateCacheRequest) error
	// DeleteCache Deletes a cache and all the items within your Momento account.
	DeleteCache(ctx context.Context, request *DeleteCacheRequest) error
	// ListCaches Lists all caches in your Momento account.
	ListCaches(ctx context.Context, request *ListCachesRequest) (*ListCachesResponse, error)

	// Set Stores an item in cache.
	Set(ctx context.Context, request *CacheSetRequest) error
	// Get Retrieve an item from the cache. Using cache key of type []bytes.
	Get(ctx context.Context, request *CacheGetRequest) (CacheGetResponse, error)
	// Delete an item from the cache.
	Delete(ctx context.Context, request *CacheDeleteRequest) error

	// Close Closes the client.
	Close()
}

// DefaultScsClient represents all information needed for momento client to enable cache control and data operations.
type DefaultScsClient struct {
	credentialProvider auth.CredentialProvider
	controlClient      *services.ScsControlClient
	dataClient         *services.ScsDataClient
	defaultTTLSeconds  uint32
}

type SimpleCacheClientProps struct {
	Configuration      config.Configuration
	CredentialProvider auth.CredentialProvider
	DefaultTTLSeconds  uint32
}

// NewSimpleCacheClient returns a new ScsClient with provided authToken, DefaultTTLSeconds, and opts arguments.
func NewSimpleCacheClient(props *SimpleCacheClientProps) (ScsClient, error) {
	if props.Configuration.GetClientSideTimeout() < 1 {
		return nil, momentoerrors.NewMomentoSvcErr(momentoerrors.InvalidArgumentError, "request timeout must not be 0", nil)
	}
	client := &DefaultScsClient{
		credentialProvider: props.CredentialProvider,
	}

	controlClient, err := services.NewScsControlClient(&models.ControlClientRequest{
		CredentialProvider: props.CredentialProvider,
		Configuration:      props.Configuration,
	})
	if err != nil {
		return nil, convertMomentoSvcErrorToCustomerError(momentoerrors.ConvertSvcErr(err))
	}

	dataClient, err := services.NewScsDataClient(&models.DataClientRequest{
		CredentialProvider: props.CredentialProvider,
		Configuration:      props.Configuration,
		DefaultTtlSeconds:  props.DefaultTTLSeconds,
	})
	if err != nil {
		return nil, convertMomentoSvcErrorToCustomerError(momentoerrors.ConvertSvcErr(err))
	}

	client.dataClient = dataClient
	client.controlClient = controlClient

	return client, nil
}

func (c *DefaultScsClient) CreateCache(ctx context.Context, request *CreateCacheRequest) error {
	err := c.controlClient.CreateCache(ctx, &models.CreateCacheRequest{
		CacheName: request.CacheName,
	})
	if err != nil {
		return convertMomentoSvcErrorToCustomerError(err)
	}
	return nil
}

func (c *DefaultScsClient) DeleteCache(ctx context.Context, request *DeleteCacheRequest) error {
	err := c.controlClient.DeleteCache(ctx, &models.DeleteCacheRequest{
		CacheName: request.CacheName,
	})
	if err != nil {
		return convertMomentoSvcErrorToCustomerError(err)
	}
	return nil
}

func (c *DefaultScsClient) ListCaches(ctx context.Context, request *ListCachesRequest) (*ListCachesResponse, error) {
	rsp, err := c.controlClient.ListCaches(ctx, &models.ListCachesRequest{
		NextToken: request.NextToken,
	})
	if err != nil {
		return nil, convertMomentoSvcErrorToCustomerError(err)
	}
	return &ListCachesResponse{
		nextToken: rsp.NextToken,
		caches:    convertCacheInfo(rsp.Caches),
	}, nil
}

func (c *DefaultScsClient) Set(ctx context.Context, request *CacheSetRequest) error {
	ttlToUse := c.defaultTTLSeconds
	if request.TTLSeconds._ttl != nil {
		ttlToUse = *request.TTLSeconds._ttl
	}
	err := c.dataClient.Set(ctx, &models.CacheSetRequest{
		CacheName:  request.CacheName,
		Key:        request.Key,
		Value:      request.Value,
		TtlSeconds: ttlToUse,
	})
	return convertMomentoSvcErrorToCustomerError(err)
}

func (c *DefaultScsClient) Get(ctx context.Context, request *CacheGetRequest) (CacheGetResponse, error) {
	rsp, err := c.dataClient.Get(ctx, &models.CacheGetRequest{
		CacheName: request.CacheName,
		Key:       request.Key,
	})
	if err != nil {
		return nil, convertMomentoSvcErrorToCustomerError(err)
	}
	return convertCacheGetResponse(rsp)
}

func (c *DefaultScsClient) Delete(ctx context.Context, request *CacheDeleteRequest) error {
	err := utility.IsKeyValid(request.Key)
	if err != nil {
		return convertMomentoSvcErrorToCustomerError(err)
	}
	err = c.dataClient.Delete(ctx, &models.CacheDeleteRequest{
		CacheName: request.CacheName,
		Key:       request.Key,
	})
	return convertMomentoSvcErrorToCustomerError(err)
}

func (c *DefaultScsClient) Close() {
	defer c.controlClient.Close()
	defer c.dataClient.Close()
}

func convertCacheGetResponse(r models.CacheGetResponse) (CacheGetResponse, MomentoError) {
	switch response := r.(type) {
	case *models.CacheGetMiss:
		return &CacheGetMiss{}, nil
	case *models.CacheGetHit:
		return &CacheGetHit{
			value: response.Value,
		}, nil
	default:
		return nil, momentoerrors.NewMomentoSvcErr(
			ClientSdkError,
			fmt.Sprintf("unexpected cache get status returned %+v", response),
			nil,
		)
	}
}

func convertMomentoSvcErrorToCustomerError(e momentoerrors.MomentoSvcErr) MomentoError {
	if e == nil {
		return nil
	}
	return NewMomentoError(e.Code(), e.Message(), e.OriginalErr())
}

func convertCacheInfo(i []models.CacheInfo) []CacheInfo {
	var convertedList []CacheInfo
	for _, c := range i {
		convertedList = append(convertedList, CacheInfo{
			name: c.Name,
		})
	}
	return convertedList
}
