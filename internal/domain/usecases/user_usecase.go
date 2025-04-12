package usecases

import (
	"context"

	"github.com/marcopaulosilva/poc_devin/internal/domain/entities"
)

type UserUseCase interface {
	GetUser(ctx context.Context, id string) (*entities.User, error)
	GetUsers(ctx context.Context) ([]entities.User, error)
}

type UserUseCaseImpl struct {
	userRepo UserRepository
}

type UserRepository interface {
	GetUser(ctx context.Context, id string) (*entities.User, error)
	GetUsers(ctx context.Context) ([]entities.User, error)
}

func NewUserUseCase(userRepo UserRepository) UserUseCase {
	return &UserUseCaseImpl{
		userRepo: userRepo,
	}
}

func (uc *UserUseCaseImpl) GetUser(ctx context.Context, id string) (*entities.User, error) {
	return uc.userRepo.GetUser(ctx, id)
}

func (uc *UserUseCaseImpl) GetUsers(ctx context.Context) ([]entities.User, error) {
	return uc.userRepo.GetUsers(ctx)
}
