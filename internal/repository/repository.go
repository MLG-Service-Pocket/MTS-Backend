package repository

import (
	"MTCTrueTechHack/internal/entities"
	"context"
)

type UserRepo interface {
	Create(ctx context.Context, user entities.User) (int, error)
	GetUser(ctx context.Context, userID int) (entities.User, error)
}

type EventsRepo interface {
	CreateAll(ctx context.Context, events []entities.Event) error
	GetAll(ctx context.Context) ([]entities.Event, error)
}
