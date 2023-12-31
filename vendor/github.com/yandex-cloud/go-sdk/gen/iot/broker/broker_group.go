// Code generated by sdkgen. DO NOT EDIT.

package broker

import (
	"context"

	"google.golang.org/grpc"
)

// Broker provides access to "broker" component of Yandex.Cloud
type Broker struct {
	getConn func(ctx context.Context) (*grpc.ClientConn, error)
}

// NewBroker creates instance of Broker
func NewBroker(g func(ctx context.Context) (*grpc.ClientConn, error)) *Broker {
	return &Broker{g}
}

// Broker gets BrokerService client
func (b *Broker) Broker() *BrokerServiceClient {
	return &BrokerServiceClient{getConn: b.getConn}
}
