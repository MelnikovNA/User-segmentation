package gateway

import (
	"context"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/sirupsen/logrus"
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

func NewGateway(ctx context.Context, cfg *Configs, lg *logrus.Logger) (*Gateway, error) {
	mux := runtime.NewServeMux(
		runtime.WithErrorHandler(apierrors)
	)
}