package entity

type PostID string

type Post struct {
	Id          PostID
	Description string
	ImageUrl    string
	Accepts     []UserID
	Rejects     []UserID
	IsAccepted  bool
}
