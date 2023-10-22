// Code generated by sdkgen. DO NOT EDIT.

//nolint
package kafka

import (
	"context"

	"google.golang.org/grpc"

	kafka "github.com/yandex-cloud/go-genproto/yandex/cloud/mdb/kafka/v1"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/operation"
)

//revive:disable

// ConnectorServiceClient is a kafka.ConnectorServiceClient with
// lazy GRPC connection initialization.
type ConnectorServiceClient struct {
	getConn func(ctx context.Context) (*grpc.ClientConn, error)
}

// Create implements kafka.ConnectorServiceClient
func (c *ConnectorServiceClient) Create(ctx context.Context, in *kafka.CreateConnectorRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return kafka.NewConnectorServiceClient(conn).Create(ctx, in, opts...)
}

// Delete implements kafka.ConnectorServiceClient
func (c *ConnectorServiceClient) Delete(ctx context.Context, in *kafka.DeleteConnectorRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return kafka.NewConnectorServiceClient(conn).Delete(ctx, in, opts...)
}

// Get implements kafka.ConnectorServiceClient
func (c *ConnectorServiceClient) Get(ctx context.Context, in *kafka.GetConnectorRequest, opts ...grpc.CallOption) (*kafka.Connector, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return kafka.NewConnectorServiceClient(conn).Get(ctx, in, opts...)
}

// List implements kafka.ConnectorServiceClient
func (c *ConnectorServiceClient) List(ctx context.Context, in *kafka.ListConnectorsRequest, opts ...grpc.CallOption) (*kafka.ListConnectorsResponse, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return kafka.NewConnectorServiceClient(conn).List(ctx, in, opts...)
}

type ConnectorIterator struct {
	ctx  context.Context
	opts []grpc.CallOption

	err           error
	started       bool
	requestedSize int64
	pageSize      int64

	client  *ConnectorServiceClient
	request *kafka.ListConnectorsRequest

	items []*kafka.Connector
}

func (c *ConnectorServiceClient) ConnectorIterator(ctx context.Context, req *kafka.ListConnectorsRequest, opts ...grpc.CallOption) *ConnectorIterator {
	var pageSize int64
	const defaultPageSize = 1000
	pageSize = req.PageSize
	if pageSize == 0 {
		pageSize = defaultPageSize
	}
	return &ConnectorIterator{
		ctx:      ctx,
		opts:     opts,
		client:   c,
		request:  req,
		pageSize: pageSize,
	}
}

func (it *ConnectorIterator) Next() bool {
	if it.err != nil {
		return false
	}
	if len(it.items) > 1 {
		it.items[0] = nil
		it.items = it.items[1:]
		return true
	}
	it.items = nil // consume last item, if any

	if it.started && it.request.PageToken == "" {
		return false
	}
	it.started = true

	if it.requestedSize == 0 || it.requestedSize > it.pageSize {
		it.request.PageSize = it.pageSize
	} else {
		it.request.PageSize = it.requestedSize
	}

	response, err := it.client.List(it.ctx, it.request, it.opts...)
	it.err = err
	if err != nil {
		return false
	}

	it.items = response.Connectors
	it.request.PageToken = response.NextPageToken
	return len(it.items) > 0
}

func (it *ConnectorIterator) Take(size int64) ([]*kafka.Connector, error) {
	if it.err != nil {
		return nil, it.err
	}

	if size == 0 {
		size = 1 << 32 // something insanely large
	}
	it.requestedSize = size
	defer func() {
		// reset iterator for future calls.
		it.requestedSize = 0
	}()

	var result []*kafka.Connector

	for it.requestedSize > 0 && it.Next() {
		it.requestedSize--
		result = append(result, it.Value())
	}

	if it.err != nil {
		return nil, it.err
	}

	return result, nil
}

func (it *ConnectorIterator) TakeAll() ([]*kafka.Connector, error) {
	return it.Take(0)
}

func (it *ConnectorIterator) Value() *kafka.Connector {
	if len(it.items) == 0 {
		panic("calling Value on empty iterator")
	}
	return it.items[0]
}

func (it *ConnectorIterator) Error() error {
	return it.err
}

// Pause implements kafka.ConnectorServiceClient
func (c *ConnectorServiceClient) Pause(ctx context.Context, in *kafka.PauseConnectorRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return kafka.NewConnectorServiceClient(conn).Pause(ctx, in, opts...)
}

// Resume implements kafka.ConnectorServiceClient
func (c *ConnectorServiceClient) Resume(ctx context.Context, in *kafka.ResumeConnectorRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return kafka.NewConnectorServiceClient(conn).Resume(ctx, in, opts...)
}

// Update implements kafka.ConnectorServiceClient
func (c *ConnectorServiceClient) Update(ctx context.Context, in *kafka.UpdateConnectorRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return kafka.NewConnectorServiceClient(conn).Update(ctx, in, opts...)
}
