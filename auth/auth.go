package auth

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"cek/exception"
	"cek/helper"
	"cek/model/domain"

	c "cek/configuration"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/twinj/uuid"
)

type CreateAuthFunc func(userID uint, tokenDetails *TokenDetails)

type AccessDetails struct {
	UserID   uint
	Role     string
	Level    string
	Nip      string
	Name     string
	MainID   uint
	MainRole string
}

type NewAccessDetails struct {
	Nip  string
	Name string
	Role string
	ID   uint
}

type TokenDetails struct {
	AccessToken         string
	RefreshToken        string
	AccessUUID          string
	RefreshUUID         string
	UserAgent           string
	RemoteAddress       string
	AtExpired           int64
	RefreshTokenExpired int64
}

func Auth(next func(c *gin.Context, auth *AccessDetails), roles []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Check JWT Token
		tokenAuth, err := ExtractTokenMetadata(ExtractToken(c.Request))
		if err != nil {
			helper.PanicIfError(exception.ErrUnauthorized)
		}

		// Check Permission User
		// if !helper.Contains(roles, tokenAuth.Role) {
		// 	helper.PanicIfError(exception.ErrPermissionDenied)
		// }

		next(c, tokenAuth)
	}
}

func CreateToken(user *domain.User, level string, userAgent, remoteAddress *string, createAuthFunc CreateAuthFunc) (*TokenDetails, error) {
	td := &TokenDetails{
		AtExpired:           time.Now().Add(time.Hour * 24 * 3).Unix(),
		AccessUUID:          uuid.NewV4().String(),
		RefreshTokenExpired: time.Now().Add(time.Hour * 24 * 7).Unix(),
		RefreshUUID:         uuid.NewV4().String(),
		UserAgent:           *userAgent,
		RemoteAddress:       *remoteAddress,
	}

	var err error

	configuration, err := c.LoadConfig()
	if err != nil {
		return nil, err
	}

	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["access_uuid"] = td.AccessUUID
	atClaims["id"] = user.ID
	atClaims["role"] = user.Role
	atClaims["name"] = user.Name
	atClaims["nip"] = user.Nip
	atClaims["level"] = level
	atClaims["exp"] = td.AtExpired
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	td.AccessToken, err = at.SignedString([]byte(configuration.AccessSecret))
	if err != nil {
		return nil, err
	}

	rtClaims := jwt.MapClaims{}
	rtClaims["refresh_uuid"] = td.RefreshUUID
	rtClaims["id"] = user.ID
	rtClaims["exp"] = td.RefreshTokenExpired
	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)
	td.RefreshToken, err = rt.SignedString([]byte(configuration.RefreshSecret))
	if err != nil {
		return nil, err
	}

	createAuthFunc(user.ID, td)

	return td, nil
}

func ExtractTokenMetadata(stringToken string) (*AccessDetails, error) {
	token, err := VerifyToken(stringToken)
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		userID, err := strconv.ParseUint(fmt.Sprintf("%.f", claims["id"]), 10, 64)
		if err != nil {
			return nil, err
		}
		return &AccessDetails{
			Nip:    claims["nip"].(string),
			UserID: uint(userID),
			Name:   claims["name"].(string),
			Role:   claims["role"].(string),
			Level:  claims["level"].(string),
		}, nil
	}
	return nil, err
}

func VerifyToken(tokenString string) (*jwt.Token, error) {
	configuration, err := c.LoadConfig()
	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(configuration.AccessSecret), nil
	})
	return token, nil
}

func ExtractToken(r *http.Request) string {
	bearToken := r.Header.Get("Authorization")
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}
