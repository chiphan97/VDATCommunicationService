package handler

import (
	_ "bytes"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"gitlab.com/vdat/mcsvc/chat/pkg/model"
	"gitlab.com/vdat/mcsvc/chat/pkg/service"
	"gitlab.com/vdat/mcsvc/chat/pkg/utils"
	"log"
	"net/http"
	"strings"
)

var userOnlineBroker *WsBroker

func UserOnlineHandler(w http.ResponseWriter, r *http.Request) {
	// authenticate
	var (
		claims model.UserClaims

		// https://accounts.vdatlab.com/auth/realms/vdatlab.com/protocol/openid-connect/certs
		publicKey = "MIICpTCCAY0CBgFrPLdvYjANBgkqhkiG9w0BAQsFADAWMRQwEgYDVQQDDAt2ZGF0bGFiLmNvbTAeFw0xOTA2MDkxNDQ4MDNaFw0yOTA2MDkxNDQ5NDNaMBYxFDASBgNVBAMMC3ZkYXRsYWIuY29tMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAleIkHoO6Q0GRQ4POIAKmN5Ev3zfAm8raTJQ1e/CbTXW4FQ0kDS9YPhLXPwcdnbxiL3rSGgz7+iWcq/Ix7yExuNbSyqDUjLUJSU6I9JvB1YP8GSaO8d996+TVCDC8E/VSID6wmfWbMNb5Ns6Y7YY/HAhj9zc73ObErvi0NV0BjeYAVOBqJKKgl9cHfyBshr+kpC/7nrbTRnAP7JQhKrQF6wBTKQiuJlEyYqvi1ugCRBYg2BZLPtTry+Kineb1DT8ynmxJjKMtr9hU0dsLPJpqW/4DWwNOarLOBP/K9WkfR2LUxbrm41goSTjJbz6s7f/Mvn/gDLjGjIsdlFP3Y7I2lwIDAQABMA0GCSqGSIb3DQEBCwUAA4IBAQBXU5Awwhv/cYJKCdSUzmtXpXty8KrdrHaDNa8potDXlEc2JrK3wHyFRwwfpBhkaicP0LllxRHUGUNsWFnggae1fudc75fysZ16NPH7VJlUuyV96K06K4v1aM5VCSWl5djky7rtyfi2W9iH2ddWZvCeSyFsSgCD4P5GjgYpsLy27g/cvdJJAdp/b7bweVDI1grlBtnInxLUPhJ4cnoNw3crh7twqKgG6F3GmZc2Hjl45LdlxBFfftDUYH66D1X0mdoipQCbg4JWlIxUZHVjJDIrSIlwnRMwjzCm7MUYv0ySmvsxgoNVI2NuFU6A/F7zlyVkDkmO4ilp4BueRtBKb7yR"
	)
	{
		var tok string
		queryToken := r.URL.Query()["token"]
		if len(queryToken) > 0 {
			tok = queryToken[0]
		}

		var accessToken string
		h := r.Header.Get("Authorization")
		a := strings.Split(h, " ")
		if len(a) == 2 {
			accessToken = a[1]
		} else if tok != "" {
			accessToken = tok
		} else {
			w.WriteHeader(400)
			return
		}

		token, err := jwt.ParseWithClaims(accessToken, &claims, func(token *jwt.Token) (interface{}, error) {
			return jwt.ParseRSAPublicKeyFromPEM([]byte("-----BEGIN CERTIFICATE-----\n" + publicKey + "\n-----END CERTIFICATE-----"))
		})

		if err != nil {
			fmt.Println("CANNOT parse token: ", err)
			w.WriteHeader(401)
			return
		}

		if token.Valid {
			fmt.Println("Token is VALID")
		} else if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				fmt.Println("That's not even a token")
			} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
				// Token is either expired or not active yet
				fmt.Println("Timing is everything")
			} else {
				fmt.Println("Couldn't handle this token:", err)
			}
		} else {
			fmt.Println("Couldn't handle this token:", err)
		}
	}

	if userOnlineBroker == nil {
		userOnlineBroker = &WsBroker{
			Inbound:    make(chan WsMessage),
			Outbound:   make(chan WsMessage),
			Register:   make(chan *WsClient),
			Unregister: make(chan *WsClient),
			Clients:    make(map[*WsClient]bool),
		}

		go userOnlineBroker.run()
	}

	conn, err := WsUpgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	user := model.UserOnline{
		HostName: r.URL.Path,
		SocketID: claims.Subject,
		UserID:   claims.Subject,
		Username: claims.UserName,
		First:    claims.GivenName,
		Last:     claims.FamilyName,
	}

	client := &WsClient{User: user, Broker: userOnlineBroker, Conn: conn, Send: make(chan []byte, 256)}
	client.Broker.Register <- client

	// Allow collection of memory referenced by the caller by doing all work in
	// new goroutines.
	go client.CheckUserOnlinePump(client.User.UserID)
	go client.WritePump()
	go client.ReadPump()

}

func RegisterUserOnline() {
	http.HandleFunc("/users", UserOnlineApi)
}

//API tìm kiếm người dùng filtter
func UserOnlineApi(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		fil := r.URL.Query()["keyword"]
		userOns, err := service.GetListUSerOnlineService(fil[0])
		if err != nil {
			utils.ResponseErr(w, http.StatusNotFound)
		}
		w.Write(utils.ResponseWithByte(userOns))
	default:
		utils.ResponseErr(w, http.StatusBadRequest)
	}
}
