package grpcserver

import (
	"context"

	"github.com/MelnikovNA/User-segmentation/proto/codegen/go/common"
	"github.com/MelnikovNA/User-segmentation/proto/codegen/go/segmentation"
	"github.com/MelnikovNA/User-segmentation/segmentation-service/internal/domain"
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

func newResponse(err error) (*common.Response, error) {
	response := &common.Response{
		Result: err == nil,
	}

	if err != nil {
		response.Error = &common.Error{
			Error: err.Error(),
		}
	}

	return response, err
}

func (s SegmentationServer) DeletSegment(ctx context.Context, req *segmentation.DeleteSegmentRequest) (*common.Response, error) {
	s.logger.Printf("SegmentID: %v Deleted", req.Id)
	err := s.service.Segmentation.DeletSegment(ctx, req.Id)
	return newResponse(err)
}

func (s SegmentationServer) CreateSegment(ctx context.Context, req *segmentation.CreateSegmentRequest) (*segmentation.CreateSegmentResponse, error) {
	id, err := s.service.Segmentation.CreateSegment(ctx, &domain.Segmentation{
		Name: req.Name,
	})
	if err != nil {
		return nil, err
	}

	return &segmentation.CreateSegmentResponse{
		Id:      id,
		Success: true}, nil

}

func (s SegmentationServer) UpdateSegment(ctx context.Context, req *segmentation.UpdateSegmentRequest) (*common.Response, error) {
	err := s.service.Segmentation.UpdateSegment(ctx, &domain.Segmentation{
		ID:   req.Id,
		Name: req.SegmentNewName,
	})
	return newResponse(err)
}

func (s SegmentationServer) GetUserSegments(ctx context.Context, req *segmentation.GetUserSegmentsRequest) (*segmentation.GetUserSegmentsResponse, error) {
	segmentations, err := s.service.Segmentation.GetUserSegments(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	var result []*segmentation.SegmentObject

	for _, seg := range segmentations {
		result = append(result, &segmentation.SegmentObject{
			Id:   seg.ID,
			Name: seg.Name,
		})
	}

	return &segmentation.GetUserSegmentsResponse{Segments: result, Error: "ok"}, nil
}

func (s SegmentationServer) AssignRandomSegments(ctx context.Context, req *segmentation.AssignRandomSegmentsRequest) (*common.Response, error) {
	err := s.service.Segmentation.AssignRandomSegments(ctx, req.Id, req.Percentage)
	return newResponse(err)
}

func (s SegmentationServer) ListSegments(ctx context.Context, req *segmentation.ListSegmentsRequest) (*segmentation.ListSegmentsResponse, error) {
	segmentations, err := s.service.Segmentation.ListSegments(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	var result []*segmentation.SegmentObject

	for _, seg := range segmentations {
		result = append(result, &segmentation.SegmentObject{
			Id:   seg.ID,
			Name: seg.Name,
		})
	}
	return &segmentation.ListSegmentsResponse{Segments: result, Error: nil}, nil
}
