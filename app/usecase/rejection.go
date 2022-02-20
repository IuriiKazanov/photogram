package usecase

import (
	"errors"
	"photogram/app/entity"
)

type RejectionUseCase struct {
	RejectionRepository RejectionRepository
	PostRepository      PostRepository
}

func (ru *RejectionUseCase) Reject(userID entity.UserID, postID entity.PostID) error {
	post, err := ru.PostRepository.GetByID(postID)
	if err != nil {
		return err
	}

	if post.UserID == userID {
		return errors.New("author can't reject own posts")
	}

	rejection := &entity.Rejection{PostID: postID, UserID: userID}
	err = ru.RejectionRepository.Store(rejection)
	if err != nil {
		return err
	}

	return nil
}
