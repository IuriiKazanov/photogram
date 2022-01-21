package usecase

import (
	"github.com/google/uuid"
	"photogram/app/constants"
	"photogram/app/entity"
)

type PostUseCase struct {
	UserRepository UserRepository
	PostRepository PostRepository
}

func NewPostUseCase(ur UserRepository, pr PostRepository) *PostUseCase {
	return &PostUseCase{
		UserRepository: ur,
		PostRepository: pr,
	}
}

func (pu *PostUseCase) Create(imageUrl, description string) (*entity.Post, error) {
	newPost := &entity.Post{
		Id:          entity.PostID(uuid.New().String()),
		Description: description,
		ImageUrl:    imageUrl,
		Accepts:     []entity.UserID{},
		Rejects:     []entity.UserID{},
		IsAccepted:  false,
	}

	err := pu.PostRepository.Store(newPost)
	if err != nil {
		return nil, err
	}

	return newPost, nil
}

func (pu *PostUseCase) Accept(userId entity.UserID, postID entity.PostID) error {
	post, err := pu.PostRepository.GetById(postID)
	if err != nil {
		return err
	}

	post.Accepts = append(post.Accepts, userId)
	if len(post.Accepts) >= constants.AcceptanceCriterion && !post.IsAccepted {
		post.IsAccepted = true
	}

	err = pu.PostRepository.Update(post)
	if err != nil {
		return err
	}

	return nil
}

func (pu *PostUseCase) Reject(userId entity.UserID, postID entity.PostID) error {
	post, err := pu.PostRepository.GetById(postID)
	if err != nil {
		return err
	}

	post.Rejects = append(post.Rejects, userId)

	// TODO: block posts or users with a lot rejects

	err = pu.PostRepository.Update(post)
	if err != nil {
		return err
	}

	return nil
}

func (pu *PostUseCase) AddToBookmark(userId entity.UserID, postID entity.PostID) error {
	user, err := pu.UserRepository.GetById(userId)
	if err == nil {
		return err
	}

	user.Bookmarks = append(user.Bookmarks, postID)

	err = pu.UserRepository.Update(user)
	if err != nil {
		return err
	}

	return nil
}
