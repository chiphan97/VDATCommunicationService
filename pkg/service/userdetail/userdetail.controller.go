package userdetail

import (
	"bytes"
	"encoding/json"
	"github.com/gorilla/mux"
	"gitlab.com/vdat/mcsvc/chat/pkg/service/auth"
	"gitlab.com/vdat/mcsvc/chat/pkg/service/cors"
	"gitlab.com/vdat/mcsvc/chat/pkg/service/useronline"
	"gitlab.com/vdat/mcsvc/chat/pkg/service/utils"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

func RegisterUserApi(r *mux.Router) {
	r.HandleFunc("/api/v1/user", GetUserApi).Methods(http.MethodGet, http.MethodOptions)
	r.HandleFunc("/api/v1/user/info", auth.AuthenMiddleJWT(CheckUserDetailApi)).Methods(http.MethodGet, http.MethodOptions)
	r.HandleFunc("/api/v1/user/online", auth.AuthenMiddleJWT(UserLogOutApi)).Methods(http.MethodDelete, http.MethodOptions)
}

//API tìm kiếm người dùng filtter

// find user by keyword godoc
// @Summary find user by keyword
// @Description find user by keyword
// @Tags user
// @Accept  json
// @Produce  json
// @Param keyword query string false "name search by keyword"
// @Param page query int false "page"
// @Param pageSize query int false "pageSize"
// @Success 200 {array} Dto
// @Router /api/v1/user [get]
func GetUserApi(w http.ResponseWriter, r *http.Request) {
	cors.SetupResponse(&w, r)
	fil := r.URL.Query()["keyword"]
	page := r.URL.Query()["page"]
	pageSize := r.URL.Query()["pageSize"]

	if page[0] == "" {
		page[0] = "1"
	}
	if pageSize[0] == "" {
		pageSize[0] = "10"
	}
	listUser := getData(fil[0], page[0], pageSize[0])
	if len(listUser) == 0 {
		json.NewEncoder(w).Encode(listUser)
	} else {
		w.Write(utils.ResponseWithByte(listUser))
	}
	//a:= []string{"b9018379-8394-4205-9104-2d85d69943db","b767e36c-e4a9-4d8c-886c-181427ec4e2c"}
	//getListFromUserId(a)
}

// checkUser godoc
// @Summary check user api
// @Description check user api
// @Tags user
// @Accept  json
// @Produce  json
// @Success 200 {object} Dto
// @Router /api/v1/user/info [get]
func CheckUserDetailApi(w http.ResponseWriter, r *http.Request) {
	cors.SetupResponse(&w, r)

	payload, err := JWTparseUser(r.Header.Get("Authorization"))
	if err != nil {
		utils.ResponseErr(w, http.StatusBadRequest)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	dto, err := CheckUserDetailService(payload)
	if err != nil {
		utils.ResponseErr(w, http.StatusInternalServerError)
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatalln(err)
		return
	}

	//dto.HostName = utils.GetLocalIP()
	dto.HostName = payload.ID
	dto.SocketID = utils.ArraySocketId[0]
	utils.ArraySocketId = utils.DeleteItemInArray(utils.ArraySocketId)
	utils.WriteLines(utils.ArraySocketId, "socketid.data")

	uo := useronline.Payload{
		HostName: dto.HostName,
		SocketID: dto.SocketID,
		UserID:   payload.ID,
	}
	err = useronline.AddUserOnlineService(uo)
	if err != nil {
		utils.ResponseErr(w, http.StatusInternalServerError)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(utils.ResponseWithByte(dto))

	// check user he thong neu login chua ton tai thong tin trong he thong thi ghi vao database

}

// user logout godoc
// @Summary user logout
// @Description user logout api
// @Tags user
// @Accept  json
// @Produce  json
// @Param hostName query string false "hostName"
// @Param socketId query string false "socketId"
// @Success 200 {object} boolean
// @Router /api/v1/user/online [delete]
func UserLogOutApi(w http.ResponseWriter, r *http.Request) {
	cors.SetupResponse(&w, r)
	hostname := r.URL.Query()["hostName"][0]
	socketID := r.URL.Query()["socketId"][0]
	err := useronline.DeleteUserOnlineService(socketID, hostname)
	if err != nil {
		utils.ResponseErr(w, http.StatusInternalServerError)
		return
	}
	utils.ArraySocketId = utils.RestoreItemArray(utils.ArraySocketId, socketID)
	utils.WriteLines(utils.ArraySocketId, "socketid.data")
	w.Write(utils.ResponseWithByte(true))
}

func connect() string {
	const (
		clientSecret string = "7161982e-cabe-44d3-ade1-324698d2f5d8"
		clientId     string = "chat.services.vdatlab.com"
		urlHost      string = "https://accounts.vdatlab.com/auth/realms/vdatlab.com/protocol/openid-connect/token"
	)

	client := &http.Client{}
	data := url.Values{}
	data.Set("client_id", clientId)
	data.Add("client_secret", clientSecret)
	data.Add("grant_type", "client_credentials")

	req, err := http.NewRequest("POST", urlHost, bytes.NewBufferString(data.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")

	if err != nil {
		log.Println(err)
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
	}

	f, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}
	resp.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	var token Token
	json.Unmarshal(f, &token)
	//fmt.Print(token.AccessToken)
	//fmt.Println(string(f))

	return token.AccessToken
}
