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
	GetById(postId entity.PostID) (*entity.Post, error)
	Update(post *entity.Post) error
	Delete(postId entity.PostID) error
}
