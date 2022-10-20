package services

import (
	"encoding/json"
	"errors"
	"time"

	jwt "github.com/golang-jwt/jwt/v4"
)

type JwtWrapper struct {
	SecretKey      string
	Issuer         string
	ExpirationHour int64
}

type JwtClaims struct {
	Authorized bool
	Exp        uint
	Role_name  string
	User_id    uint
}

func (j *JwtWrapper) GenerateToken(userID uint, roleName string) (signedToken string, err error) {

	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user_id"] = userID
	claims["role_name"] = roleName
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(j.ExpirationHour)).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err = token.SignedString([]byte(j.SecretKey))

	if err != nil {
		return
	}
	return
}

func (j *JwtWrapper) ValidateToken(signedToken string) (data *JwtClaims, err error) {
	//Parse sign Token to jwt token (ถอดรหัสตัว Token)
	token, err := jwt.Parse(
		signedToken,
		func(token *jwt.Token) (interface{}, error) {
			return []byte(j.SecretKey), nil
		},
	)
	if err != nil {
		return
	}

	//set MapClaims ( จัดฟอร์มให้เป็น jwt.Mapclaims )
	claims, ok := token.Claims.(jwt.MapClaims)

	//check token valid ? and claims Success (ok) ?
	if ok && token.Valid {
		//change jwt.claims to json ( Pretty for use )
		newClaims, err := Pretty(claims)
		//change json to struct for use
		data = ClaimsData(newClaims)

		//check token has exp ?
		if data.Exp < uint(time.Now().Local().Unix()) {
			err = errors.New("JWT IS EXPIRED")
			return nil, err
		}

		if err != nil {
			return nil, err
		}

	}
	return data, nil
}

func Pretty(claims interface{}) (data []byte, err error) {
	data, err = json.MarshalIndent(claims, "", "")

	if err != nil {
		return nil, err
	}
	return
}

func ClaimsData(claims []byte) (data *JwtClaims) {
	json.Unmarshal(claims, &data)
	return
}
