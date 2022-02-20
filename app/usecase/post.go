package usecase

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"photogram/app/api"
	"photogram/app/entity"
)

type PostUseCase struct {
	Validator      *validator.Validate
	PostRepository PostRepository
}

func NewPostUseCase(pr PostRepository, v *validator.Validate) *PostUseCase {
	return &PostUseCase{
		Validator:      v,
		PostRepository: pr,
	}
}

func (pu *PostUseCase) Create(post *api.Post) (*entity.Post, error) {
	newPost := &entity.Post{
		Id:          entity.PostID(uuid.New().String()),
		UserID:      entity.UserID(post.UserID),
		Description: post.Description,
		ImageUrl:    post.ImageUrl,
	}

	err := pu.Validator.Struct(newPost)
	if err != nil {
		return nil, err
	}

	err = pu.PostRepository.Store(newPost)
	if err != nil {
		return nil, err
	}

	return newPost, nil
}
