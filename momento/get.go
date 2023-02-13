package momento

import (
	"context"

	client_sdk_go "github.com/momentohq/client-sdk-go/internal/protos"
)

///// GetResponse ///////

type GetResponse interface {
	isGetResponse()
}

// Miss response to a cache Get api request.
type GetMiss struct{}

func (GetMiss) isGetResponse() {}

// Hit response to a cache Get api request.
type GetHit struct {
	value []byte
}

func (GetHit) isGetResponse() {}

// ValueString Returns value stored in cache as string if there was Hit. Returns an empty string otherwise.
func (resp GetHit) ValueString() string {
	return string(resp.value)
}

// ValueByte Returns value stored in cache as bytes if there was Hit. Returns nil otherwise.
func (resp GetHit) ValueByte() []byte {
	return resp.value
}

///// GetRequest ///////

type GetRequest struct {
	// Name of the cache to get the item from
	CacheName string
	// string or byte key to be used to store item
	Key Bytes

	grpcRequest  *client_sdk_go.XGetRequest
	grpcResponse *client_sdk_go.XGetResponse
	response     GetResponse
}

func (r GetRequest) cacheName() string { return r.CacheName }

func (r GetRequest) key() Bytes { return r.Key }

func (r GetRequest) requestName() string { return "Get" }

func (r *GetRequest) initGrpcRequest(scsDataClient) error {
	var err error

	var key []byte
	if key, err = prepareKey(r); err != nil {
		return err
	}

	r.grpcRequest = &client_sdk_go.XGetRequest{
		CacheKey: key,
	}

	return nil
}

func (r *GetRequest) makeGrpcRequest(client scsDataClient, metadata context.Context) (grpcResponse, error) {
	resp, err := client.grpcClient.Get(metadata, r.grpcRequest)
	if err != nil {
		return nil, err
	}

	r.grpcResponse = resp

	return resp, nil
}

func (r *GetRequest) interpretGrpcResponse() error {
	resp := r.grpcResponse

	if resp.Result == client_sdk_go.ECacheResult_Hit {
		r.response = GetHit{value: resp.CacheBody}
		return nil
	} else if resp.Result == client_sdk_go.ECacheResult_Miss {
		r.response = GetMiss{}
		return nil
	} else {
		return errUnexpectedGrpcResponse
	}
}