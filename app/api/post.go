package api

import "photogram/app/entity"

type Post struct {
	UserID      entity.UserID
	Description string
	ImageUrl    string
}
