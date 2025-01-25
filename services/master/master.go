package main

import (
	"context"
	"math"
	"sync"

	"github.com/Mau-MD/SSEngine/libs/proto"
	"github.com/redis/go-redis/v9"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type MasterStatus int

const (
	MasterStatusAlive MasterStatus = iota
	MasterStatusStarting
	MasterStatusDead
)

type Master struct {
	proto.UnimplementedMasterServiceServer
	ctx                 context.Context
	redis               *redis.Client
	sessionToStreamNode sync.Map
	streams             []*StreamNode
	status              MasterStatus
}

type StreamNodeId string
type StreamNode struct {
	id       StreamNodeId
	load     int
	capacity int
	uri      string
}

func NewMaster(ctx context.Context, redis *redis.Client) *Master {
	return &Master{
		ctx:    ctx,
		redis:  redis,
		status: MasterStatusStarting,
	}
}

// sync with redis cache if the master died
func (m *Master) init() {
	sessionToStreamNode, err := m.redis.HGetAll(m.ctx, "session_to_stream_node").Result()
	if err != nil {
		return
	}

	for sessionId, streamNodeId := range sessionToStreamNode {
		m.sessionToStreamNode.Store(sessionId, streamNodeId)
	}
}

// Stream Node will call this method when a Listener is closed
func (m *Master) TerminateSession(ctx context.Context, req *proto.TerminateSessionRequest) (*proto.TerminateSessionResponse, error) {
	m.sessionToStreamNode.LoadAndDelete(req.SessionId)
	m.redis.HDel(m.ctx, "session_to_stream_node", req.SessionId)
	return &proto.TerminateSessionResponse{
		Success: true,
	}, nil
}

// For init requests
func (m *Master) AssignAndGetStreamNode(ctx context.Context, req *proto.AssignAndGetStreamNodeRequest) (*proto.AssignAndGetStreamNodeResponse, error) {
	streamNode := m.getStreamNodeWithLeastLoad()
	m.sessionToStreamNode.Store(req.SessionId, streamNode.id)
	m.redis.HSet(m.ctx, "session_to_stream_node", req.SessionId, streamNode.id)
	return &proto.AssignAndGetStreamNodeResponse{
		StreamNodeUri: streamNode.uri,
	}, nil
}

// for client requests
func (m *Master) GetStreamNode(ctx context.Context, req *proto.GetStreamNodeRequest) (*proto.GetStreamNodeResponse, error) {
	streamNodeId, ok := m.sessionToStreamNode.Load(req.SessionId)
	if !ok {
		return nil, status.Errorf(codes.NotFound, "stream node not found")
	}

	for _, streamNode := range m.streams {
		if streamNode.id == streamNodeId {
			return &proto.GetStreamNodeResponse{
				StreamNodeUri: streamNode.uri,
			}, nil
		}
	}
	return nil, status.Errorf(codes.NotFound, "stream node not found")
}

func (m *Master) getStreamNodeWithLeastLoad() *StreamNode {
	leastLoad := math.MaxInt
	leastLoadStreamNode := &StreamNode{}
	for _, streamNode := range m.streams {
		if streamNode.load < leastLoad {
			leastLoad = streamNode.load
			leastLoadStreamNode = streamNode
		}
	}
	return leastLoadStreamNode
}

func (m *Master) discoverStreamNodes() {
	// TODO: discover stream nodes from k8s
}

// Terminate Session from Stream Node
