package token

import (
	"errors"
	"log"
	"time"

	"root/config"
	pb "root/genprotos"

	"github.com/form3tech-oss/jwt-go"
)

type Tokens struct {
	RefreshToken string
}

var tokenKey = config.Load().TokenKey

func GenereteJWTToken(user *pb.LoginUserResponse) *Tokens {
	refreshToken := jwt.New(jwt.SigningMethodHS256)

	rftclaims := refreshToken.Claims.(jwt.MapClaims)
	rftclaims["user_id"] = user.Id
	rftclaims["username"] = user.UserName
	rftclaims["iat"] = time.Now().Unix()
	rftclaims["exp"] = time.Now().Add(24 * time.Hour).Unix()
	refresh, err := refreshToken.SignedString([]byte(tokenKey))
	if err != nil {
		log.Fatal("error while genereting refresh token : ", err)
	}

	return &Tokens{
		RefreshToken: refresh,
	}
}

func ExtractClaim(tokenStr string) (jwt.MapClaims, error) {
	var (
		token *jwt.Token
		err   error
	)

	keyFunc := func(token *jwt.Token) (interface{}, error) {
		return []byte(tokenKey), nil
	}
	token, err = jwt.Parse(tokenStr, keyFunc)
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !(ok && token.Valid) {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}
