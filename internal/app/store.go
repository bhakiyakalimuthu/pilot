package app

import "context"

type Store interface {
	CreateProfile(ctx context.Context)
}


