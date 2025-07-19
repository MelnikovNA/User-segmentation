package grpcserver

import (
	"github.com/MelnikovNA/User-segmentation/proto/codegen/go/segmentation"
	"github.com/MelnikovNA/User-segmentation/segmentation-service/internal/service"
	"github.com/sirupsen/logrus"
)

type SegmentationServer struct {
	segmentation.UnimplementedSegmentServiceServer
	logger  *logrus.Logger
	service *service.Services
}

func newSegmentationServer(logger *logrus.Logger, service *service.Services) SegmentationServer {
	return SegmentationServer{
		logger:  logger,
		service: service}
}
