package main

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"

	"github.com/Mau-MD/SSEngine/libs/proto"
	"github.com/redis/go-redis/v9"
	"google.golang.org/grpc"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	defer grpcServer.Stop()

	redisClient := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	defer redisClient.Close()

	proto.RegisterMasterServiceServer(grpcServer, NewMaster(ctx, redisClient))

	go func() {

		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)
	<-signalChan

	cancel()
}
