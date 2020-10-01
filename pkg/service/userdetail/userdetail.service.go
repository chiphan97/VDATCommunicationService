package userdetail

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"gitlab.com/vdat/mcsvc/chat/pkg/service/auth"
	"gitlab.com/vdat/mcsvc/chat/pkg/service/database"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func AddUserDetailService(payload Payload) error {
	detail := payload.convertToModel()
	err := NewRepoImpl(database.DB).AddUserDetail(detail)
	if err != nil {
		return err
	}
	return nil
}
func UpdateUserDetailservice(payload Payload) error {
	detail := payload.convertToModel()
	err := NewRepoImpl(database.DB).UpdateUserDetail(detail)
	if err != nil {
		return err
	}
	return nil
}
func GetUserDetailByIDService(id string) (Dto, error) {
	var dto Dto
	detail, err := NewRepoImpl(database.DB).GetUserDetailById(id)
	if err != nil {
		return dto, err
	}
	dto = detail.ConvertToDto()
	return dto, nil
}
func GetListUserDetailService() ([]Dto, error) {
	dtos := make([]Dto, 0)
	userdetails, err := NewRepoImpl(database.DB).GetListUser()
	if err != nil {
		return dtos, err
	}
	for _, detail := range userdetails {
		dto := detail.ConvertToDto()
		dtos = append(dtos, dto)
	}
	return dtos, nil
}

//neu user chua co thi add thong tin tu token vao database user
func CheckUserDetailService(payload Payload) (Dto, error) {
	var dto Dto
	userdetail, err := NewRepoImpl(database.DB).GetUserDetailById(payload.ID)
	if err != nil {
		return dto, err
	}
	if userdetail == (UserDetail{}) {
		payload.Role = PATIENT
		err = AddUserDetailService(payload)
		if err != nil {
			return dto, err
		}
	} else {
		payload.Role = userdetail.Role
		err = UpdateUserDetailservice(payload)
		if err != nil {
			return dto, err
		}
	}

	userdetail, err = NewRepoImpl(database.DB).GetUserDetailById(payload.ID)
	if err != nil {
		return dto, err
	}
	dto = userdetail.ConvertToDto()
	return dto, nil
}

func JWTparseUser(tokenHeader string) (Payload, error) {
	var payload Payload
	splitted := strings.Split(tokenHeader, " ") // Bearer jwt_token

	block, _ := pem.Decode([]byte(auth.Jwtkey))
	var cert *x509.Certificate
	cert, _ = x509.ParseCertificate(block.Bytes)

	rsaPublicKey := cert.PublicKey.(*rsa.PublicKey)
	tokenPart := splitted[1]
	tk := &auth.UserClaims{}
	_, err := jwt.ParseWithClaims(tokenPart, tk, func(token *jwt.Token) (interface{}, error) {
		return rsaPublicKey, nil
	})
	if err != nil {
		return payload, err
	}
	payload = Payload{
		ID:       tk.Subject,
		FullName: tk.FullName,
		Username: tk.UserName,
		First:    tk.GivenName,
		Last:     tk.FamilyName,
	}
	return payload, nil
}

func getData(token string, keyword string, page string, pageSize string) []Dto {

	size, _ := strconv.Atoi(pageSize)
	pageInt, _ := strconv.Atoi(page)

	if size <= 0 {
		size = 10
	}
	if pageInt <= 0 {
		pageInt = 1
	}

	num := size * pageInt
	var expectNum int
	if pageInt > 1 {
		expectNum = size * (pageInt - 1)
	} else {
		expectNum = size * pageInt
	}

	fmt.Println(num)
	fmt.Println(expectNum)

	var (
		urlHost string = "https://vdat-mcsvc-kc-admin-api-auth-proxy.vdatlab.com/auth/admin/realms/vdatlab.com/users?search="
	)
	URL := fmt.Sprintf(urlHost+"%s"+"&max=%s"+"&first=%s", keyword, strconv.Itoa(num), strconv.Itoa(expectNum))
	fmt.Println(URL)

	var bearer = "Bearer " + token

	req, err := http.NewRequest("GET", URL, nil)
	req.Header.Add("Authorization", bearer)
	// Send req using http Client
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error on response.\n[ERRO] -", err)
	}
	body, _ := ioutil.ReadAll(resp.Body)
	var users []User
	json.Unmarshal([]byte(body), &users)
	//fmt.Print(users)
	var userDtos []Dto
	for i, _ := range users {
		fmt.Println(users[i].ID)
		detail, _ := NewRepoImpl(database.DB).GetUserDetailById(users[i].ID)
		if detail == (UserDetail{}) {
			fmt.Println("khong ton tai")
			users[i].Role = PATIENT
			payload := Payload{
				ID:   users[i].ID,
				Role: PATIENT,
			}
			err = AddUserDetailService(payload)
			if err != nil {
				fmt.Println(err)
			}
			Dto := users[i].ConvertUserToDto()
			userDtos = append(userDtos, Dto)

		} else {
			users[i].Role = detail.Role
			Dto := users[i].ConvertUserToDto()
			userDtos = append(userDtos, Dto)
		}
	}
	//fmt.Print(string(body))
	fmt.Println(len(userDtos))
	return userDtos
}

func GetListFromUserId(listUser []string) []Dto {
	var (
		urlHost string = "https://vdat-mcsvc-kc-admin-api-auth-proxy.vdatlab.com/auth/admin/realms/vdatlab.com/users/"
	)
	token := connect()
	var bearer = "Bearer " + token
	var userDtos []Dto
	for i, _ := range listUser {
		fmt.Println(listUser[i])
		req, err := http.NewRequest("GET", urlHost+listUser[i], nil)
		req.Header.Add("Authorization", bearer)
		// Send req using http Client
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			log.Println("Error on response.\n[ERRO] -", err)
		}
		body, _ := ioutil.ReadAll(resp.Body)
		var user User
		json.Unmarshal(body, &user)
		detail, _ := NewRepoImpl(database.DB).GetUserDetailById(listUser[i])
		if detail == (UserDetail{}) {
			user.Role = ""
		} else {
			user.Role = detail.Role
		}
		dto := user.ConvertUserToDto()
		userDtos = append(userDtos, dto)
	}
	//fmt.Println(userDtos)
	fmt.Println(len(userDtos))
	return userDtos
}
