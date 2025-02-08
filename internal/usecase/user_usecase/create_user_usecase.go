package user_usecase

import (
	"context"
	"fullcycle-auction_go/internal/entity/user_entity"
	"fullcycle-auction_go/internal/internal_error"
)

type UserInputDTO struct {
	Name string `json:"name" binding:"required,min=1"`
}

type UserOutputDTO struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type UserUseCase struct {
	UserRepository user_entity.UserRepositoryInterface
}

func NewUserUseCase(
	userRepository user_entity.UserRepositoryInterface) UserUseCaseInterface {
	return &UserUseCase{
		UserRepository: userRepository,
	}
}

type UserUseCaseInterface interface {
	CreateUser(
		ctx context.Context,
		userInput UserInputDTO) (*UserOutputDTO, *internal_error.InternalError)
	FindUserById(
		ctx context.Context,
		id string) (*UserOutputDTO, *internal_error.InternalError)
}

func (u *UserUseCase) CreateUser(
	ctx context.Context,
	userInput UserInputDTO) (*UserOutputDTO, *internal_error.InternalError) {
	user, err := user_entity.CreateUser(userInput.Name)
	if err != nil {
		return nil, err
	}

	if err := u.UserRepository.CreateUser(ctx, user); err != nil {
		return nil, err
	}

	return &UserOutputDTO{
		Id:   user.Id,
		Name: user.Name,
	}, nil
}
