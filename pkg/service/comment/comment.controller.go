package comment

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"gitlab.com/vdat/mcsvc/chat/pkg/service/auth"
	"gitlab.com/vdat/mcsvc/chat/pkg/service/cors"
	"gitlab.com/vdat/mcsvc/chat/pkg/service/utils"
	"log"
	"net/http"
	"strconv"
)

func NewHandler(r *mux.Router) {
	r.HandleFunc("/api/v1/comment/{idArticle}", auth.AuthenMiddleJWT(GetCommentByArticleID)).Methods(http.MethodGet, http.MethodOptions)
	r.HandleFunc("/api/v1/comment/parent/{idParent}", auth.AuthenMiddleJWT(GetCommentByParentID)).Methods(http.MethodGet, http.MethodOptions)
	r.HandleFunc("/api/v1/comment", auth.AuthenMiddleJWT(CreateComment)).Methods(http.MethodPost, http.MethodOptions)
	r.HandleFunc("/api/v1/comment/rely", auth.AuthenMiddleJWT(CreateRely)).Methods(http.MethodPost, http.MethodOptions)
	r.HandleFunc("/api/v1/comment/{id}", auth.AuthenMiddleJWT(DeleteComment)).Methods(http.MethodDelete, http.MethodOptions)
	r.HandleFunc("/api/v1/comment/{id}", auth.AuthenMiddleJWT(UpdateCmt)).Methods(http.MethodPut, http.MethodOptions)

}

func GetCommentByArticleID(w http.ResponseWriter, r *http.Request) {
	cors.SetupResponse(&w, r)
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["idArticle"])
	list, err := GetCommentByArticle(int64(id))
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

func GetCommentByParentID(w http.ResponseWriter, r *http.Request) {
	cors.SetupResponse(&w, r)
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["idParent"])
	list, err := GetCommentByParentId(int64(id))
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

func CreateComment(w http.ResponseWriter, r *http.Request) {
	cors.SetupResponse(&w, r)
	user := auth.JWTparseOwner(r.Header.Get("Authorization"))
	fmt.Println(user)
	var payload PayLoad
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		log.Println(err)
		utils.ResponseErr(w, http.StatusBadRequest)
		return
	}
	payload.UserId = user
	dto, err := AddComment(payload)
	if err != nil {
		log.Println(err)
		utils.ResponseErr(w, http.StatusInternalServerError)
		return
	}
	utils.ResponseOk(w, dto)
}

func CreateRely(w http.ResponseWriter, r *http.Request) {
	cors.SetupResponse(&w, r)
	user := auth.JWTparseOwner(r.Header.Get("Authorization"))
	fmt.Println(user)
	var payload PayLoad
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		log.Println(err)
		utils.ResponseErr(w, http.StatusBadRequest)
		return
	}
	payload.UserId = user
	dto, err := AddRelyComment(payload)
	if err != nil {
		log.Println(err)
		utils.ResponseErr(w, http.StatusInternalServerError)
		return
	}
	utils.ResponseOk(w, dto)
}

func UpdateCmt(w http.ResponseWriter, r *http.Request) {
	cors.SetupResponse(&w, r)
	params := mux.Vars(r)
	user := auth.JWTparseOwner(r.Header.Get("Authorization"))
	id, _ := strconv.Atoi(params["id"])
	fmt.Println(user)
	var payload PayLoad
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		log.Println(err)
		utils.ResponseErr(w, http.StatusBadRequest)
		return
	}
	payload.UserId = user
	dto, err := UpdateComment(payload, int64(id))
	if err != nil {
		log.Println(err)
		utils.ResponseErr(w, http.StatusInternalServerError)
		return
	}
	utils.ResponseOk(w, dto)
}

func DeleteComment(w http.ResponseWriter, r *http.Request) {
	cors.SetupResponse(&w, r)
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	err = deleteComment(int64(id))
	if err != nil {
		log.Println(err)
		utils.ResponseErr(w, http.StatusInternalServerError)
		return
	}
	result := utils.ResponseBool{Result: true}
	utils.ResponseOk(w, result)
}
