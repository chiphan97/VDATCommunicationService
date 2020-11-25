package comment

type Repo interface {
	GetCommentByArticleID(id int64) ([]Comment, error)
	GetCommentByParentID(idParent int64) ([]Comment, error)
	InsertComment(comment Comment) (Comment, error)
	InsertRelyComment(comment Comment) (Comment, error)
	UpdateComment(comment Comment) (Comment, error)
	DeleteComment(id int64) error
}
