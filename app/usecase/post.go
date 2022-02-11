package usecase

import (
	"github.com/google/uuid"
	"photogram/app/entity"
)

type PostUseCase struct {
	PostRepository PostRepository
}

func NewPostUseCase(pr PostRepository) *PostUseCase {
	return &PostUseCase{
		PostRepository: pr,
	}
}

func (pu *PostUseCase) Create(imageUrl, description string, userID entity.UserID) (*entity.Post, error) {
	newPost := &entity.Post{
		Id:          entity.PostID(uuid.New().String()),
		UserID:      userID,
		Description: description,
		ImageUrl:    imageUrl,
		IsAccepted:  false,
	}

	err := pu.PostRepository.Store(newPost)
	if err != nil {
		return nil, err
	}

	return newPost, nil
}
