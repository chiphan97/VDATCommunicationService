package userdetail

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"github.com/dgrijalva/jwt-go"
	"gitlab.com/vdat/mcsvc/chat/pkg/service/auth"
	"gitlab.com/vdat/mcsvc/chat/pkg/service/database"
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
		err = AddUserDetailService(payload)
		if err != nil {
			return dto, err
		}
	} else {
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
