package app

import (
	"context"
	"go.uber.org/zap"
)



var _ Store = (*MemoryStore)(nil)

type  MemoryStore struct {
	logger *zap.Logger
	data map[string]interface{}
}

func NewMemoryStore(logger *zap.Logger, data map[string]interface{}) *MemoryStore {
	return &MemoryStore{
		logger: logger,
		data:   data,
	}
}

func (m *MemoryStore)CreateProfile(ctx context.Context)  {

}
