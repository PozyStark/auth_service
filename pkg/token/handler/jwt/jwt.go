package jwt

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
)

const (
	BASE_TOKEN    = "BASE"
	ACCESS_TOKEN  = "ACCESS"
	REFRESH_TOKEN = "REFRESH"
)

var (
	SIGN_METHOD_HS256 = jwt.SigningMethodHS256
)

var (
	ErrInvalidIssuer = errors.New("token has invalid issuer")
	ErrExpiredToken  = errors.New("token has expired")
	ErrCastToken     = errors.New("token cast error")
)

type ITokenPayload interface {
	Valid() error
	GetPayload() jwt.Claims
}

type JWTTokenHandler struct {
	Secret     string
	SignMethod jwt.SigningMethod
}

func NewTokenCreator(
	secret string,
	signMethod jwt.SigningMethod,
) *JWTTokenHandler {
	return &JWTTokenHandler{
		Secret:     secret,
		SignMethod: signMethod,
	}
}

type UserInfo struct {
	Permissions []string
	Roles       []string
	Groups      []string
}

type BaseTokenPayload struct {
	Issuer    string
	TokenId   string
	Jti       string
	IssuedAt  int64
	ExpiredAt int64
}

type RefreshTokenPayload struct {
	BaseTokenPayload
	TokenType string
	UserId    string
}

type AccessTokenPayload struct {
	BaseTokenPayload
	TokenType string
	UserId    string
	SuperUser bool
	UserInfo
}

func NewBaseTokenPayload(
	issuer string, tokenId string, jti string, expiredTime time.Duration,
) *BaseTokenPayload {
	payload := BaseTokenPayload{
		Issuer:    issuer,
		TokenId:   tokenId,
		Jti:       jti,
		IssuedAt:  time.Now().Unix(),
		ExpiredAt: time.Now().Add(expiredTime).Unix(),
	}
	return &payload
}

func NewRefreshTokenPayload(
	issuer string, tokenId string, jti string, expiredTime time.Duration, userId string,
) *RefreshTokenPayload {
	btp := BaseTokenPayload{
		Issuer:    issuer,
		TokenId:   tokenId,
		Jti:       jti,
		IssuedAt:  time.Now().Unix(),
		ExpiredAt: time.Now().Add(expiredTime).Unix(),
	}
	rtp := RefreshTokenPayload{
		BaseTokenPayload: btp,
		TokenType:        REFRESH_TOKEN,
		UserId:           userId,
	}
	return &rtp
}

func NewAccessTokenPayload(
	issuer string,
	tokenId string,
	jti string,
	expiredTime time.Duration,
	userId string,
	superuser bool,
	permissions []string,
	roles []string,
	groups []string,
) *AccessTokenPayload {
	btp := BaseTokenPayload{
		Issuer:    issuer,
		TokenId:   tokenId,
		Jti:       jti,
		IssuedAt:  time.Now().Unix(),
		ExpiredAt: time.Now().Add(expiredTime).Unix(),
	}
	atp := AccessTokenPayload{
		BaseTokenPayload: btp,
		TokenType:        ACCESS_TOKEN,
		UserId:           userId,
		SuperUser:        superuser,
		UserInfo: UserInfo{
			Permissions: permissions,
			Roles:       roles,
			Groups:      groups,
		},
	}
	return &atp
}

func (btp *BaseTokenPayload) GetPayload() jwt.Claims {
	return btp
}

func (atp *AccessTokenPayload) GetPayload() jwt.Claims {
	return atp
}

func (rtp *RefreshTokenPayload) GetPayload() jwt.Claims {
	return rtp
}

func (payload *BaseTokenPayload) Valid() error {
	expiredTime := time.Unix(payload.ExpiredAt, 0)
	if time.Now().After(expiredTime) {
		return ErrExpiredToken
	}
	return nil
}

func (th *JWTTokenHandler) CreateToken(tokenPayload ITokenPayload) (string, error) {
	token := jwt.NewWithClaims(th.SignMethod, tokenPayload.GetPayload())
	return token.SignedString([]byte(th.Secret))
}

func (th *JWTTokenHandler) ParseToken(tokenStirng string, tokenPayload ITokenPayload) (*jwt.Token, error) {
	keyFunc := func(t *jwt.Token) (any, error) {
		return []byte(th.Secret), nil
	}
	return jwt.ParseWithClaims(tokenStirng, tokenPayload, keyFunc)
}

func (th *JWTTokenHandler) ParseBaseToken(tokenString string) (*BaseTokenPayload, error) {
	parsedToken, err := th.ParseToken(tokenString, &BaseTokenPayload{})
	if err != nil {
		return nil, err
	}
	btp, ok := parsedToken.Claims.(*BaseTokenPayload)
	if !ok {
		return nil, ErrCastToken
	}
	return btp, nil
}

func (th *JWTTokenHandler) ParseAccessToken(tokenString string) (*AccessTokenPayload, error) {
	parsedToken, err := th.ParseToken(tokenString, &AccessTokenPayload{})
	if err != nil {
		return nil, err
	}
	btp, ok := parsedToken.Claims.(*AccessTokenPayload)
	if !ok {
		return nil, ErrCastToken
	}
	return btp, nil
}

func (th *JWTTokenHandler) ParseRefreshToken(tokenString string) (*RefreshTokenPayload, error) {
	parsedToken, err := th.ParseToken(tokenString, &RefreshTokenPayload{})
	if err != nil {
		return nil, err
	}
	btp, ok := parsedToken.Claims.(*RefreshTokenPayload)
	if !ok {
		return nil, ErrCastToken
	}
	return btp, nil
}
