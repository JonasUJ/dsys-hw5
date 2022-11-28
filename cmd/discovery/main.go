package main

import (
	"main/pkg/grpc/discovery"
	"main/pkg/netutil"

	"google.golang.org/grpc"
)

func main() {
	d := NewDiscovery()
	netutil.Listen("", "50050", func(s *grpc.Server) {
		discovery.RegisterDiscoveryServer(s, d)
	})
}
