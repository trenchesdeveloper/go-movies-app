package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Auth struct {
	Issuer        string
	Audience      string
	Secret        string
	TokenExpiry   time.Duration
	RefreshExpiry time.Duration
	CookieDomain  string
	CookiePath    string
	CookieName    string
}

type jwtUser struct {
	ID        int64  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type TokenPairs struct {
	Token        string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type Claims struct {
	jwt.RegisteredClaims
}

func (j *Auth) GenerateTokenPair(user *jwtUser) (TokenPairs, error) {
	// create a token
	token := jwt.New(jwt.SigningMethodHS256)
	// set the claims
	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = fmt.Sprintf("%s %s", user.FirstName, user.LastName)
	claims["sub"] = fmt.Sprint(user.ID)
	claims["aud"] = j.Audience
	claims["iss"] = j.Issuer
	claims["iat"] = time.Now().UTC().Unix()
	claims["typ"] = "JWT"
	// set the expiry for JWT
	claims["exp"] = time.Now().UTC().Add(j.TokenExpiry).Unix()

	//create a signed token
	signedAccessToken, err := token.SignedString([]byte(j.Secret))
	if err != nil {
		return TokenPairs{}, err
	}

	// create a refresh token and set claims
	refreshToken := jwt.New(jwt.SigningMethodHS256)
	refreshTokenClaims := refreshToken.Claims.(jwt.MapClaims)
	refreshTokenClaims["sub"] = fmt.Sprint(user.ID)
	refreshTokenClaims["iat"] = time.Now().UTC().Unix()
	// set the expiry for refresh token
	refreshTokenClaims["exp"] = time.Now().UTC().Add(j.RefreshExpiry).Unix()

	// create a signed refresh token
	signedRefreshToken, err := refreshToken.SignedString([]byte(j.Secret))
	if err != nil {
		return TokenPairs{}, err

	}
	// return token pair and populate the token pair struct
	return TokenPairs{
		Token:        signedAccessToken,
		RefreshToken: signedRefreshToken,
	}, nil
}

func (j *Auth) GetRefreshCookie(refreshToken string) *http.Cookie {
	// create a cookie
	cookie := http.Cookie{
		Name:     j.CookieName,
		Value:    refreshToken,
		HttpOnly: true,
		Secure:   true,
		Domain:   j.CookieDomain,
		Path:     j.CookiePath,
		SameSite: http.SameSiteStrictMode,
		Expires: time.Now().Add(j.RefreshExpiry),
		MaxAge: int(j.RefreshExpiry.Seconds()),
	}
	return &cookie
}

func (j *Auth) GetExpiredRefreshCookie() *http.Cookie {
	// create a cookie
	cookie := http.Cookie{
		Name:     j.CookieName,
		Value:    "",
		HttpOnly: true,
		Secure:   false,
		Domain:   j.CookieDomain,
		Path:     j.CookiePath,
		SameSite: http.SameSiteLaxMode,
		Expires: time.Now().Add(-1 * time.Hour),
		MaxAge: -1,
	}
	return &cookie
}
