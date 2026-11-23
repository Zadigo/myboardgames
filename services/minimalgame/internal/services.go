package internal

import (
	"context"
	"errors"
	pb "myboardgames/gen/minimalgame/v1"
	"myboardgames/minimalgame/internal/logic"
	"myboardgames/minimalgame/internal/models"
	"time"
	"uuid"

	"github.com/redis/go-redis/v9"
)

type GameServices struct {
	pb.UnimplementedTableServiceServer
	gameRegistry models.GameRegistryInterface
	redisDb      *redis.Client
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

func (s *GameServices) localGetPlayersFunc(players []string, fn func(v string)) {
	for _, p := range players {
		fn(p)
	}
}

func (s *GameServices) getBaseTableResponse(table models.TableInterface, event pb.EventName) *pb.TableResponse {
	responseId := uuid.New().String()
	return &pb.TableResponse{
		Id:        responseId,
		UpdatedAt: time.Now().String(),
		CreatedAt: time.Now().String(),
		Event: &pb.TableEvent{
			Id:       responseId,
			TableId:  table.GetID(),
			Snapshot: &pb.TableSnapshot{},
			Details: &pb.TableEventDetails{
				EventName: event,
			},
		},
	}
}

func (s *GameServices) CreateTable(ctx context.Context, req *pb.CreateTableRequest) (*pb.TableResponse, error) {
	table := logic.NewTable(s.redisDb)
	if err := s.gameRegistry.AddTable(table); err != nil {
		return nil, err
	}
	return &pb.TableResponse{
		Id:  uuid.New().String(),
		Sub: table.GetSubscriptionKey(),
		Event: &pb.TableEvent{
			TableId: table.GetID(),
			Details: &pb.TableEventDetails{
				EventName: pb.EventName_EVENT_NAME_CREATED,
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
					EventName: pb.EventName_EVENT_NAME_ADD_PLAYER,
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
	if table, err := s.localGetTable(req.GetTableId()); err != nil {
		return nil, err
	} else {
		err := table.Execute(ctx, &models.Instruction{
			Action:  models.TableEventNextTurn,
			TableID: table.GetID(),
		})

		if err != nil {
			return nil, err
		}

		return &pb.TableResponse{
			Id: uuid.New().String(),
			Event: &pb.TableEvent{
				TableId: table.GetID(),
				Details: &pb.TableEventDetails{
					EventName: pb.EventName_EVENT_NAME_NEXT_TURN,
				},
				Snapshot: nil,
			},
		}, nil
	}
}

func (s *GameServices) NextTurn(ctx context.Context, req *pb.NextTurnRequest) (*pb.TableResponse, error) {
	if _, err := s.localGetTable(req.GetTableId()); err != nil {
		return nil, err
	} else {
		return &pb.TableResponse{}, nil
	}
}

func (s *GameServices) ShuffleCards(ctx context.Context, req *pb.ShuffleCardsRequest) (*pb.TableResponse, error) {
	if _, err := s.localGetTable(req.GetTableId()); err != nil {
		return nil, err
	} else {
		return &pb.TableResponse{}, nil
	}
}

func (s *GameServices) BanPlayer(ctx context.Context, req *pb.BanPlayerRequest) (*pb.TableResponse, error) {
	if table, err := s.localGetTable(req.GetTableId()); err != nil {
		return nil, err
	} else {
		if _, err := table.GetPlayer(req.GetClientId()); err != nil {
			return nil, err
		}
		return &pb.TableResponse{}, nil
	}
}

func (s *GameServices) ExchangeCardsWithPlayer(ctx context.Context, req *pb.ExchangeCardsWithPlayerRequest) (*pb.TableResponse, error) {
	if _, err := s.localGetTable(req.GetTableId()); err != nil {
		return nil, err
	} else {
		// lhp, err := table.GetPlayer(req.GetFrom())
		// if err != nil {
		// 	return nil, err
		// }

		// rhp, err := table.GetPlayer(req.GetTo())
		// if err != nil {
		// 	return nil, err
		// }

		// for _, c := range req.GetCards() {
		// 	table.ExchangeCard(c, lhp, rhp)
		// }
		return &pb.TableResponse{}, nil
	}
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
		instCtx, cancel := context.WithDeadline(ctx, time.Now().Add(30*time.Second))
		defer cancel()

		instruction := models.NewInstruction(instCtx, req.GetActionName(), req.GetFrom())
		if err := table.Execute(instCtx, instruction); err != nil {
			return nil, err
		}

		response := s.getBaseTableResponse(table, pb.EventName_EVENT_NAME_UNSPECIFIED)
		table.PublishSubscriptionMessage(response)
		return response, nil
	}
}

func (s *GameServices) ChangeDirection(ctx context.Context, req *pb.ChangeDirectionRequest) (*pb.TableResponse, error) {
	return &pb.TableResponse{}, nil
}

func (s *GameServices) PeekCards(ctx context.Context, req *pb.PeekCardsRequest) (*pb.TableResponse, error) {
	return &pb.TableResponse{}, nil
}

func (s *GameServices) Reset(ctx context.Context, req *pb.ResetRequest) (*pb.TableResponse, error) {
	return &pb.TableResponse{}, nil
}

func (s *GameServices) PingServer(ctx context.Context, req *pb.PingServerRequest) (*pb.PingServerResponse, error) {
	return &pb.PingServerResponse{Status: "ok"}, nil
}

func (s *GameServices) PingTable(ctx context.Context, req *pb.PingTableRequest) (*pb.TableResponse, error) {
	return &pb.TableResponse{}, nil
}
