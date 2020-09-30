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
}

//API tìm kiếm người dùng filtter
func GetUserApi(w http.ResponseWriter, r *http.Request) {
	cors.SetupResponse(&w, r)
	fil := r.URL.Query()["keyword"]
	token := connect()
	w.Write(utils.ResponseWithByte(getData(token, fil[0])))

	//a:= []string{"b9018379-8394-4205-9104-2d85d69943db","b767e36c-e4a9-4d8c-886c-181427ec4e2c"}
	//getListFromUserId(a)
}
func CheckUserDetailApi(w http.ResponseWriter, r *http.Request) {
	cors.SetupResponse(&w, r)

	payload, err := JWTparseUser(r.Header.Get("Authorization"))
	if err != nil {
		utils.ResponseErr(w, http.StatusBadRequest)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	payload.Role = PATIENT

	dto, err := CheckUserDetailService(payload)
	if err != nil {
		utils.ResponseErr(w, http.StatusInternalServerError)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	uo := useronline.Payload{
		HostName: r.URL.RawPath,
		SocketID: payload.ID,
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
