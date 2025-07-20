package gateway

import (
	"context"
	"net/http"
)

type Configs struct {
	Host        string
	HttpPort    string
	GrpcClients GrpcClients
	Cors        bool
}

type GrpcClients struct {
	SegmentService string
}

type Gateway struct {
	handler http.Handler
	host    string
	port    string
}

func NewGateway(ctx context.Context, cfg *Configs, lg *logrus)
