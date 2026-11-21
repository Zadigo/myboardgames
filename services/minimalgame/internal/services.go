package internal

import (
	"context"
	"errors"
	"myboardgames/minimalgame/internal/logic"
	"myboardgames/minimalgame/internal/models"
	pb "myboardgames/minimalgame/minimalgameproto"
	"uuid"

	"github.com/redis/go-redis/v9"
)

type GameServices struct {
	pb.UnimplementedTableServiceServer
	gameRegistry models.GameRegistryInterface
	redisClient  *redis.Client
}

func (s *GameServices) localGetTable(id string) (models.TableInterface, error) {
	table, err := s.gameRegistry.GetTable(id)
	if err != nil {
		return nil, errors.Join(err, models.TableNotFound)
	}
	if table == nil {
		return nil, models.TableNotFound
	}
	return table, nil
}

func (s *GameServices) CreateTable(ctx context.Context, req *pb.CreateTableRequest) (*pb.TableResponse, error) {
	table := logic.NewTable()
	if err := s.gameRegistry.AddTable(table); err != nil {
		return nil, err
	}
	return &pb.TableResponse{
		Id: uuid.New().String(),
		Event: &pb.TableEvent{
			TableId: table.GetID(),
			Details: &pb.TableEventDetails{
				EventName: pb.EventName_TABLE_EVENT_CREATED,
			},
		},
	}, nil
}

func (s *GameServices) GetTable(ctx context.Context, req *pb.GetTableRequest) (*pb.TableResponse, error) {
	if table, err := s.localGetTable(req.GetId()); err != nil {
		return nil, err
	} else {
		return &pb.TableResponse{
			Id: uuid.New().String(),
			Event: &pb.TableEvent{
				TableId: table.GetID(),
				Details: &pb.TableEventDetails{},
			},
		}, nil
	}
}

func (s *GameServices) JoinTable(ctx context.Context, req *pb.JoinTableRequest) (*pb.TableResponse, error) {
	if table, err := s.localGetTable(req.GetTableId()); err != nil {
		return nil, err
	} else {
		player := s.gameRegistry.AddPlayer(req.GetClientId())
		table.AddPlayer(player)
		return &pb.TableResponse{
			Id: uuid.New().String(),
			Event: &pb.TableEvent{
				TableId: table.GetID(),
				Details: &pb.TableEventDetails{
					EventName: pb.EventName_TABLE_EVENT_ADD_PLAYER,
				},
			},
		}, nil
	}
}

func (s *GameServices) LeaveTable(ctx context.Context, req *pb.LeaveTableRequest) (*pb.TableResponse, error) {
	return &pb.TableResponse{}, nil
}

func (s *GameServices) StartGame(ctx context.Context, req *pb.StartGameRequest) (*pb.TableResponse, error) {
	return &pb.TableResponse{}, nil
}

func (s *GameServices) EndGame(ctx context.Context, req *pb.EndGameRequest) (*pb.TableResponse, error) {
	if table, err := s.localGetTable(req.GetTableId()); err != nil {
		return nil, err
	} else {
		s.gameRegistry.RemoveTable(table)
		return &pb.TableResponse{Id: table.GetID()}, nil
	}
}

func (s *GameServices) DistributeCards(ctx context.Context, req *pb.DistributeCardsRequest) (*pb.TableResponse, error) {
	return &pb.TableResponse{}, nil
}

func (s *GameServices) NextRound(ctx context.Context, req *pb.NextRoundRequest) (*pb.TableResponse, error) {
	return &pb.TableResponse{}, nil
}

func (s *GameServices) NextTurn(ctx context.Context, req *pb.NextTurnRequest) (*pb.TableResponse, error) {
	return &pb.TableResponse{}, nil
}

func (s *GameServices) ShuffleCards(ctx context.Context, req *pb.ShuffleCardsRequest) (*pb.TableResponse, error) {
	return &pb.TableResponse{}, nil
}

func (s *GameServices) BanPlayer(ctx context.Context, req *pb.BanPlayerRequest) (*pb.TableResponse, error) {
	return &pb.TableResponse{}, nil
}

func (s *GameServices) ExchangeCardsWithPlayer(ctx context.Context, req *pb.ExchangeCardsWithPlayerRequest) (*pb.TableResponse, error) {
	return &pb.TableResponse{}, nil
}

func (s *GameServices) ExchangeCardsWithDeck(ctx context.Context, req *pb.ExchangeCardsWithDeckRequest) (*pb.TableResponse, error) {
	return &pb.TableResponse{}, nil
}

func (s *GameServices) ExecuteInstruction(ctx context.Context, req *pb.ExecuteInstructionRequest) (*pb.TableResponse, error) {
	if table, err := s.localGetTable(req.GetTableId()); err != nil {
		return nil, err
	} else if table == nil {
		return nil, models.TableNotFound
	} else {
		if err := table.Execute(&models.Instruction{}); err != nil {
			return nil, err
		}
		return &pb.TableResponse{}, nil
	}
}

func (s *GameServices) Reset(ctx context.Context, req *pb.ResetRequest) (*pb.TableResponse, error) {
	return &pb.TableResponse{}, nil
}

func (s *GameServices) Ping(ctx context.Context, req *pb.PingRequest) (*pb.TableResponse, error) {
	return &pb.TableResponse{}, nil
}
