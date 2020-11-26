package article

import (
	"encoding/json"
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
	r.HandleFunc("/api/v1/article", auth.AuthenMiddleJWT(handler.GetArticleByUserId)).Methods(http.MethodGet, http.MethodOptions)
	r.HandleFunc("/api/v1/article", auth.AuthenMiddleJWT(handler.StoreArticle)).Methods(http.MethodPost, http.MethodOptions)
	r.HandleFunc("/api/v1/article/{idArticle}", auth.AuthenMiddleJWT(handler.GetArticleById)).Methods(http.MethodGet, http.MethodOptions)
	r.HandleFunc("/api/v1/article/{idArticle}", auth.AuthenMiddleJWT(handler.UpdateArticle)).Methods(http.MethodPut, http.MethodOptions)
	r.HandleFunc("/api/v1/article/{idArticle}", auth.AuthenMiddleJWT(handler.DeleteArticle)).Methods(http.MethodDelete, http.MethodOptions)
	r.HandleFunc("/api/v1/search-article", auth.AuthenMiddleJWT(handler.GetArticleByTitle)).Methods(http.MethodGet, http.MethodOptions)
	r.HandleFunc("/api/v1/all-article", auth.AuthenMiddleJWT(handler.GetAllArticle)).Methods(http.MethodGet, http.MethodOptions)

}

// GetAllArticle godoc
// @Summary Get all article
// @Description Get all article
// @Tags article
// @Accept  json
// @Produce  json
// @Success 200 {array} Dto
// @Router /api/v1/all-article [get]
func (h *Handler) GetAllArticle(w http.ResponseWriter, r *http.Request) {
	cors.SetupResponse(&w, r)
	ctx := r.Context()

	list, err := h.service.Fetch(ctx)
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

// GetArticleByUserId godoc
// @Summary Get article by userid
// @Description Get the article corresponding to the input user
// @Tags article
// @Accept  json
// @Produce  json
// @Success 200 {array} Dto
// @Router /api/v1/article [get]
func (h *Handler) GetArticleByUserId(w http.ResponseWriter, r *http.Request) {
	cors.SetupResponse(&w, r)
	ctx := r.Context()

	user := auth.JWTparseOwner(r.Header.Get("Authorization"))
	list, err := h.service.GetByUserId(ctx, user)
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

// GetArticleById godoc
// @Summary Get article by idArticle
// @Description Get the article corresponding to the input idArticle
// @Tags article
// @Accept  json
// @Produce  json
// @Param idArticle path int true "ID of the article to be find"
// @Success 200 {object} Dto
// @Router /api/v1/article/{idArticle} [get]
func (h *Handler) GetArticleById(w http.ResponseWriter, r *http.Request) {
	cors.SetupResponse(&w, r)
	ctx := r.Context()
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["idArticle"])
	if err != nil {
		log.Println(err)
		utils.ResponseErr(w, http.StatusBadRequest)
		return
	}
	dto, err := h.service.GetByID(ctx, int64(id))
	if err != nil {
		log.Println(err)
		utils.ResponseErr(w, http.StatusInternalServerError)
		return
	}
	utils.ResponseOk(w, dto)
}

// GetArticleByTitle godoc
// @Summary Get article by title
// @Description Get the article corresponding to the input title
// @Tags article
// @Accept  json
// @Produce  json
// @Param title query string false "name search by title"
// @Success 200 {array} Dto
// @Router /api/v1/article/{idArticle} [get]
func (h *Handler) GetArticleByTitle(w http.ResponseWriter, r *http.Request) {
	cors.SetupResponse(&w, r)
	ctx := r.Context()
	params := r.URL.Query()["title"]
	if params == nil {
		utils.ResponseErr(w, http.StatusBadRequest)
		return
	}
	list, err := h.service.GetByTitle(ctx, params[0])
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

// StoreArticle godoc
// @Summary Insert article
// @Description Insert new article
// @Tags article
// @Accept  json
// @Produce  json
// @Param payload body Payload true "insert article"
// @Success 200 {object} Dto
// @Router /api/v1/article [post]
func (h *Handler) StoreArticle(w http.ResponseWriter, r *http.Request) {
	cors.SetupResponse(&w, r)
	ctx := r.Context()

	user := auth.JWTparseOwner(r.Header.Get("Authorization"))
	var p *Payload
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		log.Println(err)
		utils.ResponseErr(w, http.StatusBadRequest)
		return
	}
	p.CreatedBy = user
	p.UpdateBy = user

	dto, err := h.service.Store(ctx, p)
	if err != nil {
		log.Println(err)
		utils.ResponseErr(w, http.StatusInternalServerError)
		return
	}
	utils.ResponseOk(w, dto)
}

// UpdateArticle godoc
// @Summary Update article by idArticle
// @Description Update the article corresponding to the input idArticle
// @Tags article
// @Accept  json
// @Produce  json
// @Param idArticle path int true "ID of the article to be updated"
// @Param payload body Payload true "update article"
// @Success 200 {object} Dto
// @Router /api/v1/article/{idArticle} [put]
func (h *Handler) UpdateArticle(w http.ResponseWriter, r *http.Request) {
	cors.SetupResponse(&w, r)
	ctx := r.Context()
	params := mux.Vars(r)
	user := auth.JWTparseOwner(r.Header.Get("Authorization"))

	id, err := strconv.Atoi(params["idArticle"])
	if err != nil {
		log.Println(err)
		utils.ResponseErr(w, http.StatusBadRequest)
		return
	}

	var p *Payload
	err = json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		log.Println(err)
		utils.ResponseErr(w, http.StatusBadRequest)
		return
	}
	p.ID = int64(id)
	p.UpdateBy = user

	dto, err := h.service.Update(ctx, p)
	if err != nil {
		log.Println(err)
		utils.ResponseErr(w, http.StatusInternalServerError)
		return
	}
	utils.ResponseOk(w, dto)
}

// DeleteArticle godoc
// @Summary Delete article by idArticle
// @Description Delete the article corresponding to the input idArticle
// @Tags article
// @Accept  json
// @Produce  json
// @Param idArticle path int true "ID of the article to be updated"
// @Success 200 {object} utils.ResponseBool
// @Router /api/v1/article/{idArticle} [delete]
func (h *Handler) DeleteArticle(w http.ResponseWriter, r *http.Request) {
	cors.SetupResponse(&w, r)
	ctx := r.Context()
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["idArticle"])
	if err != nil {
		log.Println(err)
		utils.ResponseErr(w, http.StatusBadRequest)
		return
	}
	err = h.service.Delete(ctx, int64(id))
	if err != nil {
		log.Println(err)
		utils.ResponseErr(w, http.StatusInternalServerError)
		return
	}
	result := utils.ResponseBool{Result: true}
	utils.ResponseOk(w, result)
}
