package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"sync/atomic"

	"services/counter"

	"google.golang.org/grpc"
)

type counterServer struct {
	count uint32
}

func (s *counterServer) UpdateCount(context.Context, *counter.Empty) (*counter.Response, error) {
	var newCount = atomic.AddUint32(&s.count, 1)

	return &counter.Response{
		Count: newCount,
	}, nil
}

var (
	mainServer counter.CounterServiceServer = new(counterServer)
)

func main() {
	go func() {
		lis, err := net.Listen("tcp", ":50551")
		if err != nil {
			log.Fatal(err)
		}
		defer lis.Close()

		grpcServer := grpc.NewServer()

		counter.RegisterCounterServiceServer(grpcServer, mainServer)

		if err := grpcServer.Serve(lis); err != nil {
			log.Fatal(err)
		}
	}()

	http.Handle("/", http.FileServer(http.Dir("./public")))

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
