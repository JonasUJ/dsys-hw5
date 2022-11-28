package main

import (
	"context"
	"main/pkg/grpc/replica"
	. "main/pkg/logging"
	"sync"
	"time"
)

type Replica struct {
	replica.UnimplementedReplicaServer
	bid  uint64
	done bool
	mu   sync.Mutex
}

func NewReplica() *Replica {
	return &Replica{
		mu: sync.Mutex{},
	}
}

// Bid implements replica.ReplicaServer
func (r *Replica) Bid(ctx context.Context, bid *replica.Amount) (*replica.BidResponse, error) {
	Logger.Printf("received bid request of %d", bid.Amount)

	r.mu.Lock()
	defer r.mu.Unlock()

	// Can't change bid when auction is over
	if r.done {
		Logger.Printf("auction is ended, no bid failed")
		return &replica.BidResponse{Status: replica.BidResponseStatus_Fail}, nil
	}

	// Check that they bid more than the current high
	resp := &replica.BidResponse{}
	if bid.Amount > r.bid {
		Logger.Printf("bid of %d success", bid.Amount)
		resp.Status = replica.BidResponseStatus_Success
		r.bid = bid.Amount
	} else {
		Logger.Printf("bid of %d failed", bid.Amount)
		resp.Status = replica.BidResponseStatus_Fail
	}

	return resp, nil
}

// Result implements replica.ReplicaServer
func (r *Replica) Result(context.Context, *replica.Void) (*replica.Amount, error) {
	Logger.Printf("result %d requested", r.bid)
	return &replica.Amount{Amount: r.bid}, nil
}

// End implements replica.ReplicaServer
func (r *Replica) End(ctx context.Context, amount *replica.Amount) (*replica.Amount, error) {
	Logger.Printf("end auction requested at %d", amount.Amount)

	r.mu.Lock()
	defer r.mu.Unlock()

	// If amount is 0, we get to say which value the auction ends at
	if amount.Amount == 0 {
		amount.Amount = r.bid
	}

	// Ends the auction at the given amount
	r.done = true
	r.bid = amount.Amount

	return &replica.Amount{Amount: amount.Amount}, nil
}

// Crash implements replica.ReplicaServer
func (*Replica) Crash(context.Context, *replica.Void) (*replica.Void, error) {
	time.AfterFunc(time.Second, func() { Logger.Fatal("replica crash requested") })
	return &replica.Void{}, nil
}
