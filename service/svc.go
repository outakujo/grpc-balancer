package service

import (
	"context"

	pb "grpc-balancer/pb"
)

type SvcService struct {
	tag string
	pb.UnimplementedSvcServer
}

func NewSvcService(tag string) *SvcService {
	return &SvcService{tag: tag}
}

func (s *SvcService) CreateSvc(ctx context.Context, req *pb.CreateSvcRequest) (*pb.CreateSvcReply, error) {
	return &pb.CreateSvcReply{}, nil
}
func (s *SvcService) UpdateSvc(ctx context.Context, req *pb.UpdateSvcRequest) (*pb.UpdateSvcReply, error) {
	return &pb.UpdateSvcReply{}, nil
}
func (s *SvcService) DeleteSvc(ctx context.Context, req *pb.DeleteSvcRequest) (*pb.DeleteSvcReply, error) {
	return &pb.DeleteSvcReply{}, nil
}
func (s *SvcService) GetSvc(ctx context.Context, req *pb.GetSvcRequest) (*pb.GetSvcReply, error) {
	return &pb.GetSvcReply{
		Result: s.tag,
	}, nil
}
func (s *SvcService) ListSvc(ctx context.Context, req *pb.ListSvcRequest) (*pb.ListSvcReply, error) {
	return &pb.ListSvcReply{}, nil
}
