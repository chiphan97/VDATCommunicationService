package comment

import "gitlab.com/vdat/mcsvc/chat/pkg/service/database"

func GetCommentByArticle(idArticle int64) (results []Dto, err error) {
	list, err := NewRepoImpl(database.DB).GetCommentByArticleID(idArticle)
	if err != nil {
		return nil, err
	}
	for _, cmt := range list {
		results = append(results, cmt.convertToDto())
	}
	return results, nil
}

func GetCommentByParentId(parentId int64) (results []Dto, err error) {
	list, err := NewRepoImpl(database.DB).GetCommentByParentID(parentId)
	if err != nil {
		return nil, err
	}
	for _, cmt := range list {
		results = append(results, cmt.convertToDto())
	}
	return results, nil
}

func AddComment(payload PayLoad) (Dto, error) {
	var comment Comment
	cmt := payload.convertToModel()
	lastId, err := NewRepoImpl(database.DB).InsertComment(cmt)
	if err != nil {
		return comment.convertToDto(), err
	}

	newCmt, err := NewRepoImpl(database.DB).GetCommentById(lastId)

	return newCmt.convertToDto(), nil
}

func AddRelyComment(payload PayLoad) (Dto, error) {
	var comment Comment
	cmt := payload.convertToModel()
	lastId, err := NewRepoImpl(database.DB).InsertRelyComment(cmt)
	if err != nil {
		return comment.convertToDto(), err
	}

	newCmt, err := NewRepoImpl(database.DB).GetCommentById(lastId)

	return newCmt.convertToDto(), nil
}

func deleteComment(idCmt int64) (err error) {
	err = NewRepoImpl(database.DB).DeleteComment(idCmt)
	return err
}
