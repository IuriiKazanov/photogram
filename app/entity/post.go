package entity

type PostID string

type Post struct {
	Id          PostID `validate:"required"`
	UserID      UserID `validate:"required"`
	Description string `validate:"required"`
	ImageUrl    string `validate:"required"`
}
