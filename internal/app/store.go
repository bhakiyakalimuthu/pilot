package app

import "context"

type Store interface {
	CreateUser(ctx context.Context, model *UserModel) error
	GetUser(ctx context.Context, id string) (*UserModel,error)
	CreateMedicalData(ctx context.Context, model *MedicalDataModel) error
}


