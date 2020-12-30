package app

import (
	"context"
	"go.uber.org/zap"
)



var _ Store = (*MemoryStore)(nil)

type  MemoryStore struct {
	logger *zap.Logger
	data map[string]*UserModel
}

func NewMemoryStore(logger *zap.Logger) *MemoryStore {
	data := make(map[string]*UserModel,0)
	return &MemoryStore{
		logger: logger,
		data:   data,
	}
}

func (m *MemoryStore)CreateProfile(ctx context.Context, model *UserModel) error{
	m.data[model.id.String()]= model
	return nil
}
func (m *MemoryStore ) GetProfile(ctx context.Context, id string) (*UserModel,error){
	return m.data[id],nil
}
