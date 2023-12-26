package main

import (
	"fmt"
	"net/http"
	"strings"
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
		Expires:  time.Now().Add(j.RefreshExpiry),
		MaxAge:   int(j.RefreshExpiry.Seconds()),
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
		Expires:  time.Now().Add(-1 * time.Hour),
		MaxAge:   -1,
	}
	return &cookie
}

func (j *Auth) GetTokenFromHeaderAndVerify(w http.ResponseWriter, r *http.Request) (string, *Claims, error) {
	w.Header().Add("Vary", "Authorization")

	// get the authorization header
	authHeader := r.Header.Get("Authorization")

	// check if the authorization header is empty
	if authHeader == "" {
		return "", nil, fmt.Errorf("no authorization header provided")
	}

	// split the authorization header on space
	headerParts := strings.Split(authHeader, " ")

	// check if the authorization header has two parts
	if len(headerParts) != 2 {
		return "", nil, fmt.Errorf("invalid authorization header provided")
	}

	// check if the authorization header has bearer as the first part
	if headerParts[0] != "Bearer" {
		return "", nil, fmt.Errorf("invalid authorization header provided")
	}

	// get the token from the second part of the authorization header
	token := headerParts[1]

	// declare empty claims
	claims := &Claims{}

	// parse the token
	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		// check the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("invalid signing method: %v", token.Header["alg"])
		}
		// return the secret key
		return []byte(j.Secret), nil
	})


	if err != nil {
		if strings.HasPrefix(err.Error(), "token is expired") {
			return "", nil, fmt.Errorf("token is expired")
		}
		return "", nil, err
	}

	// check if the token is valid
	if claims.Issuer != j.Issuer {
		return "", nil, fmt.Errorf("invalid token issuer")
	}

	return token, claims, nil

}
