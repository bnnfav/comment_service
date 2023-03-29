package storage

import (
	"bay_store/comment_service/storage/postgres"
	"bay_store/comment_service/storage/repo"

	"github.com/jmoiron/sqlx"
)

type IStorage interface {
	Comment() repo.CommentStoreI
}

type storagePg struct {
	db           *sqlx.DB
	commentRepo repo.CommentStoreI
}

func NewStoragePg(db *sqlx.DB) *storagePg {
	return &storagePg{
		db:           db,
		commentRepo: postgres.NewCommentRepo(db),
	}
}

func (s storagePg) Comment() repo.CommentStoreI {
	return s.commentRepo
}
