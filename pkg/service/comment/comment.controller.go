package comment

import (
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
	fmt.Println(id)
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
}

func CreateRely(w http.ResponseWriter, r *http.Request) {
	cors.SetupResponse(&w, r)
}

func DeleteComment(w http.ResponseWriter, r *http.Request) {
	cors.SetupResponse(&w, r)
}
