package app

import "go.uber.org/zap"

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

