package main

import (
	"context"
	"errors"
	"main/pkg/grpc/frontend"
	"main/pkg/grpc/replica"
	"main/pkg/netutil"
	"math/rand"
	"sync"
	"time"

	. "main/pkg/logging"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Frontend struct {
	frontend.UnimplementedFrontendServer
	replicas []replica.ReplicaClient
	mu       sync.Mutex
}

func NewFrontend() *Frontend {
	return &Frontend{
		replicas: make([]replica.ReplicaClient, 0),
		mu:       sync.Mutex{},
	}
}

// Bid implements frontend.FrontendServer
func (f *Frontend) Bid(ctx context.Context, amount *frontend.Amount) (*frontend.BidResponse, error) {
	Logger.Printf("received bid %d", amount.Amount)

	f.mu.Lock()
	defer f.mu.Unlock()

	// Forward to replicas sequentially
	for i, r := range f.replicas {
		ctx, cancel := context.WithTimeout(ctx, time.Millisecond*200)
		defer cancel()

		resp, err := r.Bid(ctx, &replica.Amount{Amount: amount.Amount})

		// Short-circuit as soon as one of the replicas fails us
		if err != nil {
			if status.Code(err) != codes.DeadlineExceeded {
				Logger.Printf("replica %d returned an error in response to bid %d: %v", i, amount.Amount, err)
				return &frontend.BidResponse{Status: frontend.BidResponseStatus_Exception}, nil
			}
		} else if resp.Status != replica.BidResponseStatus_Success {
			Logger.Printf("replica %d failed our request to bid %d", i, amount.Amount)
			return &frontend.BidResponse{Status: frontend.BidResponseStatus_Fail}, nil
		}
	}

	Logger.Printf("bid %d was successful", amount.Amount)
	return &frontend.BidResponse{Status: frontend.BidResponseStatus_Success}, nil
}

// Result implements frontend.FrontendServer
func (f *Frontend) Result(ctx context.Context, void *frontend.Void) (*frontend.Amount, error) {
	Logger.Printf("result requested")

	f.mu.Lock()
	defer f.mu.Unlock()

	for i, r := range f.replicas {
		ctx, cancel := context.WithTimeout(ctx, time.Millisecond*200)
		defer cancel()

		resp, err := r.Result(ctx, &replica.Void{})
		if err == nil {
			Logger.Printf("replica %d responded with result %d", i, resp.Amount)
			return &frontend.Amount{Amount: resp.Amount}, nil
		}
	}

	Logger.Printf("no replica responded to result request")
	return nil, errors.New("could not reach any replicas")
}

// Crash implements frontend.FrontendServer
func (*Frontend) Crash(context.Context, *frontend.Void) (*frontend.Void, error) {
	time.AfterFunc(time.Second, func() { Logger.Fatal("frontend crash requested") })
	return &frontend.Void{}, nil
}

// CrashReplica implements frontend.FrontendServer
func (f *Frontend) CrashReplica(ctx context.Context, void *frontend.Void) (*frontend.Void, error) {
	Logger.Printf("replica crash requested")

	// Shallow copy before shuffling
	replicas := make([]replica.ReplicaClient, len(f.replicas))
	copy(replicas, f.replicas)

	// Shuffle replicas
	rand.Shuffle(len(replicas), func(i, j int) {
		replicas[i], replicas[j] = replicas[j], replicas[i]
	})

	// Loop through until one crashes
	for _, r := range replicas {
		ctx, cancel := context.WithTimeout(ctx, time.Millisecond*200)
		defer cancel()

		_, err := r.Crash(ctx, &replica.Void{})
		if err == nil {
			return void, nil
		}
	}

	Logger.Printf("no replicas to crash")
	return nil, errors.New("no replicas")
}

func (f *Frontend) AddReplica(host string) {
	Logger.Printf("adding replica '%s'", host)

	f.mu.Lock()
	defer f.mu.Unlock()

	client := replica.NewReplicaClient(netutil.Dial(host, "50050"))

	f.replicas = append(f.replicas, client)
}
