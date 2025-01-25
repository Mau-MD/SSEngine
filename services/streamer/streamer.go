package main

import (
	"context"
	"log"

	"github.com/Mau-MD/SSEngine/libs/proto"
	"github.com/fatih/color"
	"github.com/google/uuid"
)

type Streamer struct {
	proto.UnimplementedStreamerServer
	id           string
	capacity     int
	uri          string
	masterClient proto.MasterServiceClient
	ctx          context.Context
}

func NewStreamer(ctx context.Context, masterClient proto.MasterServiceClient) *Streamer {
	streamer := &Streamer{
		id:           uuid.New().String(),
		capacity:     10,
		uri:          "localhost:50052",
		masterClient: masterClient,
		ctx:          ctx,
	}

	streamer.sendInitialPing()

	return streamer
}

func (s *Streamer) sendInitialPing() {
	color.Cyan("Sending initial ping to master")
	_, err := s.masterClient.ConnectToStreamNode(s.ctx, &proto.ConnectToStreamNodeRequest{
		StreamNode: &proto.StreamNode{
			Id:       s.id,
			Load:     1,
			Capacity: int32(s.capacity),
			Uri:      s.uri,
		},
	})
	if err != nil {
		log.Fatalf("failed to send initial ping: %v", err)
	}
}

func (s *Streamer) Heartbeat(ctx context.Context, req *proto.HeartbeatRequest) (*proto.HeartbeatResponse, error) {
	return &proto.HeartbeatResponse{Success: true}, nil
}
