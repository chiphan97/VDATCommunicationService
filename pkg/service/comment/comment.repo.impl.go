package comment

import "database/sql"

type RepoImpl struct {
	Db *sql.DB
}

func NewRepoImpl(db *sql.DB) Repo {
	return &RepoImpl{Db: db}
}

func (cmt RepoImpl) GetCommentById(id int64) (Comment, error) {
	cmts := make([]Comment, 0)
	query := `SELECT * FROM Comment WHERE id_cmt = $1 `
	rows, err := cmt.Db.Query(query, id)
	if err != nil {
		return cmts[0], err
	}
	for rows.Next() {
		cmt := Comment{}
		err := rows.Scan(&cmt.ID,
			&cmt.IdArticle,
			&cmt.Content,
			&cmt.Type,
			&cmt.ParentID,
			&cmt.Num,
			&cmt.Version,
			&cmt.CreatedBy,
			&cmt.UpdateBy,
			&cmt.CreatedAt,
			&cmt.UpdateAt,
		)
		if err != nil {
			return cmts[0], err
		}
		cmts = append(cmts, cmt)
	}
	defer rows.Close()
	return cmts[0], nil
}

func (cmt RepoImpl) GetCommentByArticleID(id int64) ([]Comment, error) {
	cmts := make([]Comment, 0)
	query := `SELECT * FROM Comment WHERE id_article = $1 and parentId = -1`
	rows, err := cmt.Db.Query(query, id)
	if err != nil {
		return cmts, err
	}
	for rows.Next() {
		cmt := Comment{}
		err := rows.Scan(&cmt.ID,
			&cmt.IdArticle,
			&cmt.Content,
			&cmt.Type,
			&cmt.ParentID,
			&cmt.Num,
			&cmt.Version,
			&cmt.CreatedBy,
			&cmt.UpdateBy,
			&cmt.CreatedAt,
			&cmt.UpdateAt,
		)
		if err != nil {
			return cmts, err
		}
		cmts = append(cmts, cmt)
	}
	defer rows.Close()
	return cmts, nil
}

func (cmt RepoImpl) GetCommentByParentID(idParent int64) ([]Comment, error) {
	cmts := make([]Comment, 0)
	query := `SELECT * FROM Comment WHERE parentId = $1 `
	rows, err := cmt.Db.Query(query, idParent)
	if err != nil {
		return cmts, err
	}
	for rows.Next() {
		cmt := Comment{}
		err := rows.Scan(&cmt.ID,
			&cmt.IdArticle,
			&cmt.Content,
			&cmt.Type,
			&cmt.ParentID,
			&cmt.Num,
			&cmt.Version,
			&cmt.CreatedBy,
			&cmt.UpdateBy,
			&cmt.CreatedAt,
			&cmt.UpdateAt,
		)
		if err != nil {
			return cmts, err
		}
		cmts = append(cmts, cmt)
	}
	defer rows.Close()
	return cmts, nil
}

func (cmt RepoImpl) InsertComment(comment Comment) (lastId int64, err error) {
	statement := `INSERT INTO Comment (id_article,content,type,create_by,update_by) VALUES ($1,$2,$3,$4,$5) RETURNING id_cmt`
	err = cmt.Db.QueryRow(statement,
		comment.IdArticle,
		comment.Content,
		comment.Type,
		comment.CreatedBy,
		comment.UpdateBy).Scan(&lastId)

	if err != nil {
		return
	}

	return
}

func (cmt RepoImpl) InsertRelyComment(comment Comment) (lastId int64, err error) {
	statement := `INSERT INTO Comment (id_article,content,parentId,type,create_by,update_by) VALUES ($1,$2,$3,$4,$5,$6) RETURNING id_cmt`
	err = cmt.Db.QueryRow(statement,
		comment.IdArticle,
		comment.Content,
		comment.ParentID,
		comment.Type,
		comment.CreatedBy,
		comment.UpdateBy).Scan(&lastId)
	if err != nil {
		return
	}
	return
}

func (cmt RepoImpl) UpdateComment(comment Comment, id int64) error {
	statement := `Update Comment set content= $1, type = $2 , version = version+1 where id_cmt = $3`
	_, err := cmt.Db.Exec(statement, comment.Content, comment.Type, id)
	if err != nil {
		return err
	}
	return nil

}

func (cmt RepoImpl) DeleteComment(id int64) error {
	statement := `DELETE FROM Comment WHERE id_cmt=$1`
	cmt.Db.Exec(statement, id)
	return nil
}
