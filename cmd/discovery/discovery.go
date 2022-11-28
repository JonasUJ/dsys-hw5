package main

import (
	"context"
	"main/pkg/grpc/discovery"
	"sync"

	events "github.com/jonhoo/go-events"
	. "main/pkg/logging"
)

type Discovery struct {
	discovery.UnimplementedDiscoveryServer
	replicas   []string
	chReplicas chan string
	mu         sync.Mutex
	events     events.Events
}

func NewDiscovery() *Discovery {
	return &Discovery{
		replicas:   make([]string, 0),
		chReplicas: make(chan string, 4),
		events:     events.New(),
	}
}

// RegisterReplica implements discovery.DiscoveryServer
func (d *Discovery) RegisterReplica(ctx context.Context, node *discovery.Node) (*discovery.Void, error) {
	Logger.Printf("registering new replica '%s'", node.Host)

	d.mu.Lock()
	defer d.mu.Unlock()

	d.replicas = append(d.replicas, node.Host)
	d.events.Announce(events.Event{
		Tag:  "replica",
		Data: node.Host,
	})

	return &discovery.Void{}, nil
}

// Replicas implements discovery.DiscoveryServer
func (d *Discovery) Replicas(void *discovery.Void, stream discovery.Discovery_ReplicasServer) error {
	d.mu.Lock()
	for _, r := range d.replicas {
		stream.Send(&discovery.Node{Host: r})
	}
	d.mu.Unlock()

	ch := d.events.Listen("replica")
	for r := range ch {
		host := r.Data.(string)
		stream.Send(&discovery.Node{Host: host})
	}

	return nil
}
