package comment

import "database/sql"

type RepoImpl struct {
	Db *sql.DB
}

func NewRepoImpl(db *sql.DB) Repo {
	return &RepoImpl{Db: db}
}

func (cmt RepoImpl) GetCommentByArticleID(id int64) ([]Comment, error) {
	panic("implement me")
}

func (cmt RepoImpl) GetCommentByParentID(idParent int64) ([]Comment, error) {
	panic("implement me")
}

func (cmt RepoImpl) InsertComment(comment Comment) (Comment, error) {
	panic("implement me")
}

func (cmt RepoImpl) InsertRelyComment(comment Comment) (Comment, error) {
	panic("implement me")
}

func (cmt RepoImpl) UpdateComment(comment Comment) (Comment, error) {
	panic("implement me")
}

func (cmt RepoImpl) DeleteComment(id int64) error {
	panic("implement me")
}
