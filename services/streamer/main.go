package main

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"

	"github.com/Mau-MD/SSEngine/libs/proto"
	"github.com/fatih/color"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	lis, err := net.Listen("tcp", "0.0.0.0:50052")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	defer grpcServer.Stop()

	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect to master: %v", err)
	}
	defer conn.Close()

	masterClient := proto.NewMasterServiceClient(conn)
	streamer := NewStreamer(ctx, masterClient)

	proto.RegisterStreamerServer(grpcServer, streamer)

	go func() {
		color.Green("Streamer server initialized and listening on :50052")
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)
	<-signalChan

	cancel()
}
