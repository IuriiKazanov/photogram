package usecase

import (
	"photogram/app/entity"
)

type UserRepository interface {
	Store(user *entity.User) error
	GetById(id entity.UserID) (*entity.User, error)
	Update(user *entity.User) error
}

type PostRepository interface {
	Store(post *entity.Post) error
	GetByID(postId entity.PostID) (*entity.Post, error)
	GetApproved() ([]*entity.Post, error)
	Update(post *entity.Post) error
	Delete(postId entity.PostID) error
}

type ApprovalRepository interface {
	Store(approval *entity.Approval) error
}

type BookmarkRepository interface {
	Store(approval *entity.Bookmark) error
}

type RejectionRepository interface {
	Store(approval *entity.Rejection) error
}
