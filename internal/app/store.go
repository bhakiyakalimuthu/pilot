package app

import "context"

type Store interface {
	CreateProfile(ctx context.Context, model *UserModel) error
	GetProfile(ctx context.Context, id string) (*UserModel,error)
}


