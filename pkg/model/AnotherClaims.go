package model

import "github.com/dgrijalva/jwt-go"

type UserClaims struct {
	jwt.StandardClaims
	UserName   string `json:"name"`
	GivenName  string `json:"given_name"`
	FamilyName string `json:"family_name"`
}
