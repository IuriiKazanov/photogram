package entity

type PostID string

type Post struct {
	Id          PostID
	UserID      UserID
	Description string
	ImageUrl    string
	//Accepts     []UserID
	//Rejects     []UserID
	IsAccepted bool
}
