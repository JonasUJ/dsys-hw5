package main

import (
	"context"
	"main/pkg/grpc/discovery"
	"main/pkg/grpc/frontend"
	"main/pkg/grpc/replica"
	"main/pkg/netutil"
	"math/rand"
	"time"

	. "main/pkg/logging"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func main() {
	f := NewFrontend()

	// Discover replicas in background
	go func() {
		client := discovery.NewDiscoveryClient(netutil.Dial("discovery", "50050"))

		stream, err := client.Replicas(context.Background(), &discovery.Void{})
		if err != nil {
			Logger.Fatal(err)
		}

		for {
			node, err := stream.Recv()
			if err != nil {
				Logger.Printf("discovery replicas err: %v", err)
				return
			}

			f.AddReplica(node.Host)
		}
	}()

	// End the auction after 30-60 seconds
	time.AfterFunc(time.Second*time.Duration(rand.Intn(30)+30), func() {
		Logger.Printf("ending auction")

		amt := &replica.Amount{}
		for _, r := range f.replicas {
			ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*200)
			defer cancel()

			newAmt, err := r.End(ctx, amt)
			if err != nil {
				if status.Code(err) != codes.DeadlineExceeded {
					Logger.Printf("%v", err)
				}
			} else {
				amt = newAmt
			}
		}
	})

	netutil.Listen("", "50050", func(s *grpc.Server) {
		frontend.RegisterFrontendServer(s, f)
	})
}
