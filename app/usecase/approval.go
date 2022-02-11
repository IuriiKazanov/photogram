package usecase

import (
	"errors"
	"photogram/app/entity"
)

type ApprovalUseCase struct {
	ApprovalRepository ApprovalRepository
	PostRepository     PostRepository
}

func (au *ApprovalUseCase) Approve(userID entity.UserID, postID entity.PostID) error {
	post, err := au.PostRepository.GetById(postID)
	if err != nil {
		return err
	}

	if post.UserID == userID {
		return errors.New("author can't accept own posts")
	}

	approval := &entity.Approval{PostID: postID, UserID: userID}
	err = au.ApprovalRepository.Store(approval)
	if err != nil {
		return err
	}
	return nil
}
