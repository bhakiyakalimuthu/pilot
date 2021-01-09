package app

import (
	"context"

	"go.uber.org/zap"
)

var _ Store = (*MemoryStore)(nil)

type MemoryStore struct {
	logger *zap.Logger
	data   map[string]*UserModel
}

func (m *MemoryStore) CreateMedicalData(ctx context.Context, model *MedicalDataModel) error {
	panic("implement me")
}

func (m *MemoryStore) CreateDiagnosisProfile(ctx context.Context, model *DiagnosisProfileModel) error {
	panic("implement me")
}

func NewMemoryStore(logger *zap.Logger) *MemoryStore {
	data := make(map[string]*UserModel, 0)
	return &MemoryStore{
		logger: logger,
		data:   data,
	}
}

func (m *MemoryStore) CreateUser(ctx context.Context, model *UserModel) error {
	m.data[model.ID.String()] = model
	return nil
}
func (m *MemoryStore) GetUser(ctx context.Context, id string) (*UserModel, error) {
	return m.data[id], nil
}
