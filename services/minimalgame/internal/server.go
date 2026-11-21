package internal

import (
	"context"
	pb "myboardgames/minimalgame/minimalgameproto"
)

type Server struct {
	pb.UnimplementedTableServiceServer
}

func (s *Server) CreateTable(ctx context.Context, req *pb.CreateTableRequest) (*pb.TableResponse, error) {
	return &pb.TableResponse{Id: "1"}, nil
}

func (s *Server) GetTable(ctx context.Context, req *pb.GetTableRequest) (*pb.TableResponse, error) {
	return &pb.TableResponse{Id: "1"}, nil
}
