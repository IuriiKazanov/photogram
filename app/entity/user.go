package entity

type UserID string

type User struct {
	Id        UserID
	Username  string
	Rating    int
	Bookmarks []PostID
}
