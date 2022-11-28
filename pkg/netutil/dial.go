package netutil

import (
	"net"

	"google.golang.org/grpc"
	. "main/pkg/logging"
)

func Dial(host, port string) *grpc.ClientConn {
	addr := net.JoinHostPort(host, port)
	Logger.Printf("dialing %s", addr)
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		Logger.Fatal("fail to dial")
	}

	return conn
}

func Listen(host, port string, register func(*grpc.Server)) {
	addr := net.JoinHostPort(host, port)
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		Logger.Fatal("fail to listen")
	}

	grpcServer := grpc.NewServer()
	register(grpcServer)
	Logger.Printf("serving on %s", addr)
	if err := grpcServer.Serve(listener); err != nil {
		Logger.Fatal("stopped serving")
	}
}
