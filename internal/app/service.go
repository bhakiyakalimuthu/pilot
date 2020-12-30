package app

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

type Service struct {
	logger *zap.Logger
	store Store
}

func NewService(logger *zap.Logger, store Store) *Service {
	return &Service{
		logger: logger,
		store: store,
	}
}

func (s *Service) CreateProfile(ctx context.Context, user *User)  error{
	// s.logger.Info("user input",zap.Any("user",user))
	// fmt.Println(user)

	id := uuid.New()
	fmt.Println(id)
	model := &UserModel{
		id:          id,
		Name:        user.Name,
		PhoneNumber: user.PhoneNumber,
		Address:     user.Address,
		Age:         user.Age,
		DOB:         user.DOB,
		Sex:         user.Sex,
		Height:      user.Height,
		Weight:      user.Weight,
	}
	s.store.CreateProfile(ctx,model)
	return nil
}

func (s *Service) GetProfile(ctx context.Context, userID string)  (*UserModel, error){
	return s.store.GetProfile(ctx,userID)
}