package tokenHelper

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/spf13/viper"
)

type Tokens struct {
	AccessToken  string
	RefreshToken string
}

type AccessTokenDataAdmin struct {
	AdminId uuid.UUID
}

type AccessTokenDataUser struct {
	UserId uuid.UUID
}

type RefreshTokenData struct {
	Id uuid.UUID
}

var (
	ErrCreateAccessToken  = errors.New("error create AccessToken")
	ErrCreateRefreshToken = errors.New("error create RefreshToken")
)

func CreateUserTokens(userId uuid.UUID) (*Tokens, error) {
	tokens := &Tokens{}

	var err error

	tokens.AccessToken, err = genAccessToken(userId)

	if err != nil {
		return nil, ErrCreateAccessToken
	}

	tokens.RefreshToken, err = genRefreshToken(userId)

	if err != nil {
		return nil, ErrCreateRefreshToken
	}

	return tokens, nil
}

func CreateAdminUserTokens(adminUserId uuid.UUID) (*Tokens, error) {
	tokens := &Tokens{}

	var err error

	tokens.AccessToken, err = genAccessTokenAdmin(adminUserId)

	if err != nil {
		return nil, ErrCreateAccessToken
	}

	tokens.RefreshToken, err = genRefreshToken(adminUserId)

	if err != nil {
		return nil, ErrCreateRefreshToken
	}

	return tokens, nil
}

func genAccessToken(userId uuid.UUID) (string, error) {
	accessTokenClaims := jwt.MapClaims{}
	accessTokenClaims["userId"] = userId
	accessTokenClaims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaims)

	accessTokenStr, err := accessToken.SignedString([]byte(viper.GetString("jwt.secret.access")))

	if err != nil {
		return "", err
	}
	return accessTokenStr, nil
}

func genAccessTokenAdmin(adminUserId uuid.UUID) (string, error) {
	accessTokenClaims := jwt.MapClaims{}
	accessTokenClaims["adminId"] = adminUserId
	accessTokenClaims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaims)

	accessTokenStr, err := accessToken.SignedString([]byte(viper.GetString("jwt.secret.access")))

	if err != nil {
		return "", err
	}

	return accessTokenStr, nil
}

func genRefreshToken(id uuid.UUID) (string, error) {
	refreshTokenClaims := jwt.MapClaims{}
	refreshTokenClaims["id"] = id
	refreshTokenClaims["exp"] = time.Now().Add(time.Hour * 24 * 30).Unix()

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenClaims)

	refreshTokenStr, err := refreshToken.SignedString([]byte(viper.GetString("jwt.secret.refresh")))

	if err != nil {
		return "", err
	}

	return refreshTokenStr, nil
}

func VerifyToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(viper.GetString("jwt.secret.access")), nil
	})

	if err != nil {
		return nil, err
	}

	return token, nil
}

func ExtractTokenMetadataUser(token *jwt.Token) (*AccessTokenDataUser, error) {
	claims, ok := token.Claims.(jwt.MapClaims)

	if ok && token.Valid {
		userIdStr, ok := claims["userId"].(string)

		if !ok {
			return nil, errors.New("fail claims extract")
		}

		userId, err := uuid.Parse(userIdStr)

		if err != nil {
			return nil, err
		}

		return &AccessTokenDataUser{
			UserId: userId,
		}, nil
	}
	return nil, errors.New("fail claims extract")
}

func ExtractTokenMetadataAdmin(token *jwt.Token) (*AccessTokenDataAdmin, error) {
	claims, ok := token.Claims.(jwt.MapClaims)

	if ok && token.Valid {
		adminIdStr, ok := claims["adminId"].(string)

		if !ok {
			return nil, errors.New("fail claims extract")
		}

		adminId, err := uuid.Parse(adminIdStr)

		if err != nil {
			return nil, err
		}

		return &AccessTokenDataAdmin{
			AdminId: adminId,
		}, nil
	}
	return nil, errors.New("fail claims extract")
}

func ExtractTokenMetadataRefresh(token *jwt.Token) (*RefreshTokenData, error) {
	claims, ok := token.Claims.(jwt.MapClaims)

	if ok && token.Valid {
		idStr, ok := claims["id"].(string)

		if !ok {
			return nil, errors.New("fail claims extract")
		}

		id, err := uuid.Parse(idStr)

		if err != nil {
			return nil, err
		}

		return &RefreshTokenData{
			Id: id,
		}, nil
	}
	return nil, errors.New("fail claims extract")
}

func ExtractToken(r *http.Request) string {
	token := r.Header.Get("Authorization")
	return token
}

func ExtractTokenHeader(r *http.Request) string {
	token := r.Header.Get("Authorization")
	return token
}

func TokenValid(token *jwt.Token) error {
	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		return errors.New("err valid token")
	}
	return nil
}

func VerifyTokenRefresh(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(viper.GetString("jwt.secret.refresh")), nil
	})

	if err != nil {
		return nil, err
	}

	return token, nil
}
