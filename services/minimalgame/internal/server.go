package internal

import (
	"context"
	"myboardgames/minimalgame/internal/models"
	pb "myboardgames/minimalgame/minimalgameproto"

	"github.com/redis/go-redis/v9"
)

type GameServices struct {
	pb.UnimplementedTableServiceServer
	gameRegistry models.GameRegistryInterface
	redisClient  *redis.Client
}

func (s *GameServices) CreateTable(ctx context.Context, req *pb.CreateTableRequest) (*pb.TableResponse, error) {
	return &pb.TableResponse{Id: "1"}, nil
}

func (s *GameServices) GetTable(ctx context.Context, req *pb.GetTableRequest) (*pb.TableResponse, error) {
	return &pb.TableResponse{Id: "1"}, nil
}
