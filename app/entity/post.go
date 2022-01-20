package entity

type PostID string

type Post struct {
	Id          PostID
	Description string
	Image       *Image
	Accepts     []UserID
	Rejects     []UserID
	IsAccepted  bool
}

type Image struct {
	Base64string string
	Type string
	Name string
}
