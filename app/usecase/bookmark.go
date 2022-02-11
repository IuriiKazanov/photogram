package usecase

import "photogram/app/entity"

type BookmarkUseCase struct {
	BookmarkRepository BookmarkRepository
}

func (bu *BookmarkUseCase) AddToBookmark(userId entity.UserID, postID entity.PostID) error {
	newBookmark := &entity.Bookmark{UserID: userId, PostID: postID}
	err := bu.BookmarkRepository.Store(newBookmark)
	if err != nil {
		return err
	}

	return nil
}
