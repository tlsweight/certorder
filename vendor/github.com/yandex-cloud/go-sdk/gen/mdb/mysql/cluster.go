// Code generated by sdkgen. DO NOT EDIT.

//nolint
package mysql

import (
	"context"

	"google.golang.org/grpc"

	mysql "github.com/yandex-cloud/go-genproto/yandex/cloud/mdb/mysql/v1"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/operation"
)

//revive:disable

// ClusterServiceClient is a mysql.ClusterServiceClient with
// lazy GRPC connection initialization.
type ClusterServiceClient struct {
	getConn func(ctx context.Context) (*grpc.ClientConn, error)
}

// AddHosts implements mysql.ClusterServiceClient
func (c *ClusterServiceClient) AddHosts(ctx context.Context, in *mysql.AddClusterHostsRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return mysql.NewClusterServiceClient(conn).AddHosts(ctx, in, opts...)
}

// Backup implements mysql.ClusterServiceClient
func (c *ClusterServiceClient) Backup(ctx context.Context, in *mysql.BackupClusterRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return mysql.NewClusterServiceClient(conn).Backup(ctx, in, opts...)
}

// Create implements mysql.ClusterServiceClient
func (c *ClusterServiceClient) Create(ctx context.Context, in *mysql.CreateClusterRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return mysql.NewClusterServiceClient(conn).Create(ctx, in, opts...)
}

// Delete implements mysql.ClusterServiceClient
func (c *ClusterServiceClient) Delete(ctx context.Context, in *mysql.DeleteClusterRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return mysql.NewClusterServiceClient(conn).Delete(ctx, in, opts...)
}

// DeleteHosts implements mysql.ClusterServiceClient
func (c *ClusterServiceClient) DeleteHosts(ctx context.Context, in *mysql.DeleteClusterHostsRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return mysql.NewClusterServiceClient(conn).DeleteHosts(ctx, in, opts...)
}

// Get implements mysql.ClusterServiceClient
func (c *ClusterServiceClient) Get(ctx context.Context, in *mysql.GetClusterRequest, opts ...grpc.CallOption) (*mysql.Cluster, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return mysql.NewClusterServiceClient(conn).Get(ctx, in, opts...)
}

// List implements mysql.ClusterServiceClient
func (c *ClusterServiceClient) List(ctx context.Context, in *mysql.ListClustersRequest, opts ...grpc.CallOption) (*mysql.ListClustersResponse, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return mysql.NewClusterServiceClient(conn).List(ctx, in, opts...)
}

type ClusterIterator struct {
	ctx  context.Context
	opts []grpc.CallOption

	err           error
	started       bool
	requestedSize int64
	pageSize      int64

	client  *ClusterServiceClient
	request *mysql.ListClustersRequest

	items []*mysql.Cluster
}

func (c *ClusterServiceClient) ClusterIterator(ctx context.Context, req *mysql.ListClustersRequest, opts ...grpc.CallOption) *ClusterIterator {
	var pageSize int64
	const defaultPageSize = 1000
	pageSize = req.PageSize
	if pageSize == 0 {
		pageSize = defaultPageSize
	}
	return &ClusterIterator{
		ctx:      ctx,
		opts:     opts,
		client:   c,
		request:  req,
		pageSize: pageSize,
	}
}

func (it *ClusterIterator) Next() bool {
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

	it.items = response.Clusters
	it.request.PageToken = response.NextPageToken
	return len(it.items) > 0
}

func (it *ClusterIterator) Take(size int64) ([]*mysql.Cluster, error) {
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

	var result []*mysql.Cluster

	for it.requestedSize > 0 && it.Next() {
		it.requestedSize--
		result = append(result, it.Value())
	}

	if it.err != nil {
		return nil, it.err
	}

	return result, nil
}

func (it *ClusterIterator) TakeAll() ([]*mysql.Cluster, error) {
	return it.Take(0)
}

func (it *ClusterIterator) Value() *mysql.Cluster {
	if len(it.items) == 0 {
		panic("calling Value on empty iterator")
	}
	return it.items[0]
}

func (it *ClusterIterator) Error() error {
	return it.err
}

// ListBackups implements mysql.ClusterServiceClient
func (c *ClusterServiceClient) ListBackups(ctx context.Context, in *mysql.ListClusterBackupsRequest, opts ...grpc.CallOption) (*mysql.ListClusterBackupsResponse, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return mysql.NewClusterServiceClient(conn).ListBackups(ctx, in, opts...)
}

type ClusterBackupsIterator struct {
	ctx  context.Context
	opts []grpc.CallOption

	err           error
	started       bool
	requestedSize int64
	pageSize      int64

	client  *ClusterServiceClient
	request *mysql.ListClusterBackupsRequest

	items []*mysql.Backup
}

func (c *ClusterServiceClient) ClusterBackupsIterator(ctx context.Context, req *mysql.ListClusterBackupsRequest, opts ...grpc.CallOption) *ClusterBackupsIterator {
	var pageSize int64
	const defaultPageSize = 1000
	pageSize = req.PageSize
	if pageSize == 0 {
		pageSize = defaultPageSize
	}
	return &ClusterBackupsIterator{
		ctx:      ctx,
		opts:     opts,
		client:   c,
		request:  req,
		pageSize: pageSize,
	}
}

func (it *ClusterBackupsIterator) Next() bool {
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

	response, err := it.client.ListBackups(it.ctx, it.request, it.opts...)
	it.err = err
	if err != nil {
		return false
	}

	it.items = response.Backups
	it.request.PageToken = response.NextPageToken
	return len(it.items) > 0
}

func (it *ClusterBackupsIterator) Take(size int64) ([]*mysql.Backup, error) {
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

	var result []*mysql.Backup

	for it.requestedSize > 0 && it.Next() {
		it.requestedSize--
		result = append(result, it.Value())
	}

	if it.err != nil {
		return nil, it.err
	}

	return result, nil
}

func (it *ClusterBackupsIterator) TakeAll() ([]*mysql.Backup, error) {
	return it.Take(0)
}

func (it *ClusterBackupsIterator) Value() *mysql.Backup {
	if len(it.items) == 0 {
		panic("calling Value on empty iterator")
	}
	return it.items[0]
}

func (it *ClusterBackupsIterator) Error() error {
	return it.err
}

// ListHosts implements mysql.ClusterServiceClient
func (c *ClusterServiceClient) ListHosts(ctx context.Context, in *mysql.ListClusterHostsRequest, opts ...grpc.CallOption) (*mysql.ListClusterHostsResponse, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return mysql.NewClusterServiceClient(conn).ListHosts(ctx, in, opts...)
}

type ClusterHostsIterator struct {
	ctx  context.Context
	opts []grpc.CallOption

	err           error
	started       bool
	requestedSize int64
	pageSize      int64

	client  *ClusterServiceClient
	request *mysql.ListClusterHostsRequest

	items []*mysql.Host
}

func (c *ClusterServiceClient) ClusterHostsIterator(ctx context.Context, req *mysql.ListClusterHostsRequest, opts ...grpc.CallOption) *ClusterHostsIterator {
	var pageSize int64
	const defaultPageSize = 1000
	pageSize = req.PageSize
	if pageSize == 0 {
		pageSize = defaultPageSize
	}
	return &ClusterHostsIterator{
		ctx:      ctx,
		opts:     opts,
		client:   c,
		request:  req,
		pageSize: pageSize,
	}
}

func (it *ClusterHostsIterator) Next() bool {
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

	response, err := it.client.ListHosts(it.ctx, it.request, it.opts...)
	it.err = err
	if err != nil {
		return false
	}

	it.items = response.Hosts
	it.request.PageToken = response.NextPageToken
	return len(it.items) > 0
}

func (it *ClusterHostsIterator) Take(size int64) ([]*mysql.Host, error) {
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

	var result []*mysql.Host

	for it.requestedSize > 0 && it.Next() {
		it.requestedSize--
		result = append(result, it.Value())
	}

	if it.err != nil {
		return nil, it.err
	}

	return result, nil
}

func (it *ClusterHostsIterator) TakeAll() ([]*mysql.Host, error) {
	return it.Take(0)
}

func (it *ClusterHostsIterator) Value() *mysql.Host {
	if len(it.items) == 0 {
		panic("calling Value on empty iterator")
	}
	return it.items[0]
}

func (it *ClusterHostsIterator) Error() error {
	return it.err
}

// ListLogs implements mysql.ClusterServiceClient
func (c *ClusterServiceClient) ListLogs(ctx context.Context, in *mysql.ListClusterLogsRequest, opts ...grpc.CallOption) (*mysql.ListClusterLogsResponse, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return mysql.NewClusterServiceClient(conn).ListLogs(ctx, in, opts...)
}

type ClusterLogsIterator struct {
	ctx  context.Context
	opts []grpc.CallOption

	err           error
	started       bool
	requestedSize int64
	pageSize      int64

	client  *ClusterServiceClient
	request *mysql.ListClusterLogsRequest

	items []*mysql.LogRecord
}

func (c *ClusterServiceClient) ClusterLogsIterator(ctx context.Context, req *mysql.ListClusterLogsRequest, opts ...grpc.CallOption) *ClusterLogsIterator {
	var pageSize int64
	const defaultPageSize = 1000
	pageSize = req.PageSize
	if pageSize == 0 {
		pageSize = defaultPageSize
	}
	return &ClusterLogsIterator{
		ctx:      ctx,
		opts:     opts,
		client:   c,
		request:  req,
		pageSize: pageSize,
	}
}

func (it *ClusterLogsIterator) Next() bool {
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

	response, err := it.client.ListLogs(it.ctx, it.request, it.opts...)
	it.err = err
	if err != nil {
		return false
	}

	it.items = response.Logs
	it.request.PageToken = response.NextPageToken
	return len(it.items) > 0
}

func (it *ClusterLogsIterator) Take(size int64) ([]*mysql.LogRecord, error) {
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

	var result []*mysql.LogRecord

	for it.requestedSize > 0 && it.Next() {
		it.requestedSize--
		result = append(result, it.Value())
	}

	if it.err != nil {
		return nil, it.err
	}

	return result, nil
}

func (it *ClusterLogsIterator) TakeAll() ([]*mysql.LogRecord, error) {
	return it.Take(0)
}

func (it *ClusterLogsIterator) Value() *mysql.LogRecord {
	if len(it.items) == 0 {
		panic("calling Value on empty iterator")
	}
	return it.items[0]
}

func (it *ClusterLogsIterator) Error() error {
	return it.err
}

// ListOperations implements mysql.ClusterServiceClient
func (c *ClusterServiceClient) ListOperations(ctx context.Context, in *mysql.ListClusterOperationsRequest, opts ...grpc.CallOption) (*mysql.ListClusterOperationsResponse, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return mysql.NewClusterServiceClient(conn).ListOperations(ctx, in, opts...)
}

type ClusterOperationsIterator struct {
	ctx  context.Context
	opts []grpc.CallOption

	err           error
	started       bool
	requestedSize int64
	pageSize      int64

	client  *ClusterServiceClient
	request *mysql.ListClusterOperationsRequest

	items []*operation.Operation
}

func (c *ClusterServiceClient) ClusterOperationsIterator(ctx context.Context, req *mysql.ListClusterOperationsRequest, opts ...grpc.CallOption) *ClusterOperationsIterator {
	var pageSize int64
	const defaultPageSize = 1000
	pageSize = req.PageSize
	if pageSize == 0 {
		pageSize = defaultPageSize
	}
	return &ClusterOperationsIterator{
		ctx:      ctx,
		opts:     opts,
		client:   c,
		request:  req,
		pageSize: pageSize,
	}
}

func (it *ClusterOperationsIterator) Next() bool {
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

	response, err := it.client.ListOperations(it.ctx, it.request, it.opts...)
	it.err = err
	if err != nil {
		return false
	}

	it.items = response.Operations
	it.request.PageToken = response.NextPageToken
	return len(it.items) > 0
}

func (it *ClusterOperationsIterator) Take(size int64) ([]*operation.Operation, error) {
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

	var result []*operation.Operation

	for it.requestedSize > 0 && it.Next() {
		it.requestedSize--
		result = append(result, it.Value())
	}

	if it.err != nil {
		return nil, it.err
	}

	return result, nil
}

func (it *ClusterOperationsIterator) TakeAll() ([]*operation.Operation, error) {
	return it.Take(0)
}

func (it *ClusterOperationsIterator) Value() *operation.Operation {
	if len(it.items) == 0 {
		panic("calling Value on empty iterator")
	}
	return it.items[0]
}

func (it *ClusterOperationsIterator) Error() error {
	return it.err
}

// Move implements mysql.ClusterServiceClient
func (c *ClusterServiceClient) Move(ctx context.Context, in *mysql.MoveClusterRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return mysql.NewClusterServiceClient(conn).Move(ctx, in, opts...)
}

// RescheduleMaintenance implements mysql.ClusterServiceClient
func (c *ClusterServiceClient) RescheduleMaintenance(ctx context.Context, in *mysql.RescheduleMaintenanceRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return mysql.NewClusterServiceClient(conn).RescheduleMaintenance(ctx, in, opts...)
}

// Restore implements mysql.ClusterServiceClient
func (c *ClusterServiceClient) Restore(ctx context.Context, in *mysql.RestoreClusterRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return mysql.NewClusterServiceClient(conn).Restore(ctx, in, opts...)
}

// Start implements mysql.ClusterServiceClient
func (c *ClusterServiceClient) Start(ctx context.Context, in *mysql.StartClusterRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return mysql.NewClusterServiceClient(conn).Start(ctx, in, opts...)
}

// StartFailover implements mysql.ClusterServiceClient
func (c *ClusterServiceClient) StartFailover(ctx context.Context, in *mysql.StartClusterFailoverRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return mysql.NewClusterServiceClient(conn).StartFailover(ctx, in, opts...)
}

// Stop implements mysql.ClusterServiceClient
func (c *ClusterServiceClient) Stop(ctx context.Context, in *mysql.StopClusterRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return mysql.NewClusterServiceClient(conn).Stop(ctx, in, opts...)
}

// StreamLogs implements mysql.ClusterServiceClient
func (c *ClusterServiceClient) StreamLogs(ctx context.Context, in *mysql.StreamClusterLogsRequest, opts ...grpc.CallOption) (mysql.ClusterService_StreamLogsClient, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return mysql.NewClusterServiceClient(conn).StreamLogs(ctx, in, opts...)
}

// Update implements mysql.ClusterServiceClient
func (c *ClusterServiceClient) Update(ctx context.Context, in *mysql.UpdateClusterRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return mysql.NewClusterServiceClient(conn).Update(ctx, in, opts...)
}

// UpdateHosts implements mysql.ClusterServiceClient
func (c *ClusterServiceClient) UpdateHosts(ctx context.Context, in *mysql.UpdateClusterHostsRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return mysql.NewClusterServiceClient(conn).UpdateHosts(ctx, in, opts...)
}
