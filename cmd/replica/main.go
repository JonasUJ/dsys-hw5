package main

import (
	"context"
	"main/pkg/grpc/discovery"
	"main/pkg/grpc/replica"
	"main/pkg/netutil"

	"google.golang.org/grpc"
)

func main() {
	r := NewReplica()

	// Register for discovery
	client := discovery.NewDiscoveryClient(netutil.Dial("discovery", "50050"))
	client.RegisterReplica(context.Background(), &discovery.Node{Host: netutil.Ip()})

	netutil.Listen("", "50050", func(s *grpc.Server) {
		replica.RegisterReplicaServer(s, r)
	})
}
