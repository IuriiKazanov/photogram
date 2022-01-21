package usecase

import (
	"github.com/google/uuid"
	"photogram/app/entity"
)

type UserUseCase struct {
	UserRepository UserRepository
}

func NewUserUseCase(ur UserRepository) *UserUseCase {
	return &UserUseCase{
		UserRepository: ur,
	}
}

func (uu *UserUseCase) Create(username string) (*entity.User, error) {
	newUser := &entity.User{
		Id:        entity.UserID(uuid.New().String()),
		Username:  username,
		Rating:    0,
		Bookmarks: []entity.PostID{},
	}

	err := uu.UserRepository.Store(newUser)
	if err != nil {
		return nil, err
	}

	return newUser, err
}
