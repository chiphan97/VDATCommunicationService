package comment

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"gitlab.com/vdat/mcsvc/chat/pkg/service/auth"
	"gitlab.com/vdat/mcsvc/chat/pkg/service/cors"
	"gitlab.com/vdat/mcsvc/chat/pkg/service/database"
	"gitlab.com/vdat/mcsvc/chat/pkg/service/utils"
	"log"
	"net/http"
	"strconv"
	"time"
)

type Handler struct {
	service Service
}

func NewHandler(r *mux.Router) {
	timeoutContext := time.Duration(2) * time.Second
	repo := NewRepoImpl(database.DB)
	service := NewServiceImpl(repo, timeoutContext)
	handler := &Handler{service: service}
	r.HandleFunc("/api/v1/comment/{idArticle}", auth.AuthenMiddleJWT(handler.GetCommentByArticleID)).Methods(http.MethodGet, http.MethodOptions)
	r.HandleFunc("/api/v1/comment/parent/{idParent}", auth.AuthenMiddleJWT(handler.GetCommentByParentID)).Methods(http.MethodGet, http.MethodOptions)
	r.HandleFunc("/api/v1/comment", auth.AuthenMiddleJWT(handler.CreateComment)).Methods(http.MethodPost, http.MethodOptions)
	r.HandleFunc("/api/v1/comment/rely", auth.AuthenMiddleJWT(handler.CreateRely)).Methods(http.MethodPost, http.MethodOptions)
	r.HandleFunc("/api/v1/comment/{id}", auth.AuthenMiddleJWT(handler.DeleteComment)).Methods(http.MethodDelete, http.MethodOptions)
	r.HandleFunc("/api/v1/comment/{id}", auth.AuthenMiddleJWT(handler.UpdateCmt)).Methods(http.MethodPut, http.MethodOptions)

}

// GetCommentByArticleId godoc
// @Summary Get comment by Article id
// @Description Get the comment corresponding to the input articleId
// @Tags comment
// @Accept  json
// @Produce  json
// @Param idArticle path int true "ID of the article to be find"
// @Success 200 {array} Dto
// @Router /api/v1/comment/{idArticle} [get]
func (h *Handler) GetCommentByArticleID(w http.ResponseWriter, r *http.Request) {
	cors.SetupResponse(&w, r)
	params := mux.Vars(r)
	ctx := r.Context()
	id, err := strconv.Atoi(params["idArticle"])
	list, err := h.service.GetCommentByArticle(ctx, int64(id))

	if err != nil {
		log.Println(err)
		utils.ResponseErr(w, http.StatusInternalServerError)
		return
	}
	if len(list) > 0 {
		utils.ResponseOk(w, list)
	} else {
		results := make([]Dto, 0)
		utils.ResponseOk(w, results)
	}
}

// GetCommentByParentId godoc
// @Summary Get comment by parent id
// @Description Get the comment corresponding to the input parentId
// @Tags comment
// @Accept  json
// @Produce  json
// @Param idParent path int true "ID of the parent comment to be find"
// @Success 200 {array} Dto
// @Router /api/v1/comment/parent/{idParent} [get]
func (h *Handler) GetCommentByParentID(w http.ResponseWriter, r *http.Request) {
	cors.SetupResponse(&w, r)
	params := mux.Vars(r)
	ctx := r.Context()

	id, err := strconv.Atoi(params["idParent"])
	list, err := h.service.GetCommentByParentId(ctx, int64(id))
	if err != nil {
		log.Println(err)
		utils.ResponseErr(w, http.StatusInternalServerError)
		return
	}
	if len(list) > 0 {
		utils.ResponseOk(w, list)
	} else {
		results := make([]Dto, 0)
		utils.ResponseOk(w, results)
	}

}

// addComment godoc
// @Summary Insert comment
// @Description Insert new comment
// @Tags comment
// @Accept  json
// @Produce  json
// @Param payload body PayLoad true "insert comment"
// @Success 200 {object} Dto
// @Router /api/v1/comment [post]
func (h *Handler) CreateComment(w http.ResponseWriter, r *http.Request) {
	cors.SetupResponse(&w, r)
	user := auth.JWTparseOwner(r.Header.Get("Authorization"))
	fmt.Println(user)
	var payload PayLoad
	err := json.NewDecoder(r.Body).Decode(&payload)
	ctx := r.Context()

	if err != nil {
		log.Println(err)
		utils.ResponseErr(w, http.StatusBadRequest)
		return
	}
	payload.UserId = user
	dto, err := h.service.AddComment(ctx, payload)
	if err != nil {
		log.Println(err)
		utils.ResponseErr(w, http.StatusInternalServerError)
		return
	}
	utils.ResponseOk(w, dto)
}

// addRelyComment godoc
// @Summary Insert comment rely
// @Description Insert new comment rely
// @Tags comment
// @Accept  json
// @Produce  json
// @Param payload body PayLoad true "insert rely comment"
// @Success 200 {object} Dto
// @Router /api/v1/comment/rely [post]
func (h *Handler) CreateRely(w http.ResponseWriter, r *http.Request) {
	cors.SetupResponse(&w, r)
	user := auth.JWTparseOwner(r.Header.Get("Authorization"))
	ctx := r.Context()

	var payload PayLoad
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		log.Println(err)
		utils.ResponseErr(w, http.StatusBadRequest)
		return
	}
	payload.UserId = user
	dto, err := h.service.AddRelyComment(ctx, payload)
	if err != nil {
		log.Println(err)
		utils.ResponseErr(w, http.StatusInternalServerError)
		return
	}
	utils.ResponseOk(w, dto)
}

// UpdateComment godoc
// @Summary Update Comment by idCmt
// @Description Update the Comment corresponding to the input id
// @Tags comment
// @Accept  json
// @Produce  json
// @Param id path int true "ID of the comment to be updated"
// @Param payload body PayLoad true "update comment"
// @Success 200 {object} Dto
// @Router /api/v1/comment/{id} [put]
func (h *Handler) UpdateCmt(w http.ResponseWriter, r *http.Request) {
	cors.SetupResponse(&w, r)
	params := mux.Vars(r)
	user := auth.JWTparseOwner(r.Header.Get("Authorization"))
	id, _ := strconv.Atoi(params["id"])
	ctx := r.Context()

	var payload PayLoad
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		log.Println(err)
		utils.ResponseErr(w, http.StatusBadRequest)
		return
	}
	payload.UserId = user
	dto, err := h.service.UpdateComment(ctx, payload, int64(id))
	if err != nil {
		log.Println(err)
		utils.ResponseErr(w, http.StatusInternalServerError)
		return
	}
	utils.ResponseOk(w, dto)
}

// DeleteComment godoc
// @Summary Delete comment by idCmt
// @Description Delete comment by id comment
// @Tags comment
// @Accept  json
// @Produce  json
// @Param id path int true "ID of the comment to be updated"
// @Success 200 {object} utils.ResponseBool
// @Router /api/v1/comment/{id} [delete]
func (h *Handler) DeleteComment(w http.ResponseWriter, r *http.Request) {
	cors.SetupResponse(&w, r)
	ctx := r.Context()
	fmt.Println(ctx)
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	err = h.service.deleteComment(ctx, int64(id))
	if err != nil {
		log.Println(err)
		utils.ResponseErr(w, http.StatusInternalServerError)
		return
	}
	result := utils.ResponseBool{Result: true}
	utils.ResponseOk(w, result)
}
