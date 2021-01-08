package app

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"github.com/lithammer/shortuuid"
	"time"
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

func (s *Service) CreateUser(ctx context.Context, user *User)  error{
	// s.logger.Info("user input",zap.Any("user",user))
	// fmt.Println(user)

	id := uuid.New()
	fmt.Println(id)
	shortID := shortuuid.New()
	fmt.Println(shortID)
	s.logger.Info("user id's",zap.String("id",id.String()),zap.String("shortID",shortID))
	model := &UserModel{
		ID:          id,
		ShortID:     shortID,
		Name:        user.Name,
		PhoneNumber: user.PhoneNumber,
		EmailID:     user.EmailID,
		Address:     user.Address,
		Age:         user.Age,
		DOB:         user.DOB,
		Sex:         user.Sex,
		Height:      user.Height,
		Weight:      user.Weight,
	}
	if err:=s.store.CreateUser(ctx,model);err!=nil {
		s.logger.Error("failed to create profile",zap.Error(err))
		return err
	}
	return nil
}

func (s *Service) GetUser(ctx context.Context, userID string)  (*UserModel, error){
	return s.store.GetUser(ctx,userID)
}

func (s *Service) CreateMedicalData(ctx context.Context, data *MedicalData)  error{
	m:= MedicalDataModel{
		ShortID:  data.PatientID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Year:      data.DemographicFactor.Year,
		Sex:       data.DemographicFactor.Sex,
	}
	if err:= s.store.CreateMedicalData(ctx,&m);err!=nil {
		s.logger.Error("failed to create medical data",zap.Error(err))
		return err
	}
	return nil
}