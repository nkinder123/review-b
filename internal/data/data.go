package data

import (
	"context"
	"github.com/go-kratos/kratos/contrib/registry/consul/v2"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/hashicorp/consul/api"
	v1 "review-b/api/review/v1"
	"review-b/internal/conf"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewDisCovery, NewGRPCClient, NewData, NewBusinessRepo)

// Data .
type Data struct {
	// TODO wrapped database client
	//s这个data中使用的是rpc的client
	rc  v1.ReviewClient
	log *log.Helper
}

// NewData .
func NewData(c *conf.Data, rc v1.ReviewClient, logger log.Logger) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{rc: rc, log: log.NewHelper(logger)}, cleanup, nil
}

func NewGRPCClient(d registry.Discovery) v1.ReviewClient {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("discovery:///review.service"),
		grpc.WithTimeout(3600*time.Second),
		grpc.WithDiscovery(d),
	)

	if err != nil {
		panic("rpc client connect has error")
	}
	client := v1.NewReviewClient(conn)
	return client
}
func NewDisCovery(c *conf.Register) registry.Discovery {
	config := api.DefaultConfig()
	config.Address = c.Address
	config.Scheme = c.Scheme
	client, err := api.NewClient(config)
	if err != nil {
		panic(err)
	}
	// new dis with consul client
	dis := consul.New(client)
	return dis
}
