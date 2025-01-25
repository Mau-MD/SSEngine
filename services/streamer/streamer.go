package main

import (
	"context"
)

type Streamer struct {
	ctx context.Context
}

func NewStreamer(ctx context.Context) *Streamer {
	return &Streamer{
		ctx: ctx,
	}
}
