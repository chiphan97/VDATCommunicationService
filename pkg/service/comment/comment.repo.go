package comment

type Repo interface {
	GetCommentById(id int64) (Comment, error)
	GetCommentByArticleID(id int64) ([]Comment, error)
	GetCommentByParentID(idParent int64) ([]Comment, error)
	InsertComment(comment Comment) (int64, error)
	InsertRelyComment(comment Comment) (int64, error)
	UpdateComment(comment Comment) (Comment, error)
	DeleteComment(id int64) error
}
