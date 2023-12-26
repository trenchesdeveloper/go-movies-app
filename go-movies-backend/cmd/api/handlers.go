package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt/v5"
)

func (a *application) Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from %s\n", a.Domain)
}

func (app *application) AllMovies(w http.ResponseWriter, r *http.Request) {
	movies, err := app.DB.AllMovies()

	if err != nil {
		app.errorJson(w, err)
		return
	}

	// Marshal the payload to JSON
	_ = app.writeJson(w, http.StatusOK, movies)

}

func (app *application) authenticate(w http.ResponseWriter, r *http.Request) {
	// read a json payload
	var requestPayload struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	err := app.readJson(w, r, &requestPayload)

	if err != nil {
		app.errorJson(w, err, http.StatusBadRequest)
		return
	}

	// validate user against database
	user, err := app.DB.GetUserByEmail(requestPayload.Email)

	if err != nil {
		app.errorJson(w, errors.New("invalid credentials"), http.StatusBadRequest)
		return
	}

	// check password
	valid, err := user.ValidatePassword(requestPayload.Password)

	if !valid || err != nil {
		app.errorJson(w, errors.New("invalid credentials"), http.StatusBadRequest)
		return
	}

	// create a jwt user
	u := jwtUser{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
	}

	//generate token
	tokens, err := app.auth.GenerateTokenPair(&u)

	if err != nil {
		app.errorJson(w, err)
		return
	}

	refreshCookie := app.auth.GetRefreshCookie(
		tokens.RefreshToken,
	)

	http.SetCookie(w, refreshCookie)

	// send token to user
	_ = app.writeJson(w, http.StatusAccepted, tokens)

}

func (app *application) refreshToken(w http.ResponseWriter, r *http.Request) {
	for _, cookie := range r.Cookies() {
		if cookie.Name == app.auth.CookieName {
			claims := &Claims{}
			refreshToken := cookie.Value

			//parse the token
			_, err := jwt.ParseWithClaims(refreshToken, claims, func(token *jwt.Token) (interface{}, error) {
				return []byte(app.auth.Secret), nil
			})

			if err != nil {
				app.errorJson(w, errors.New("unauthorized"), http.StatusUnauthorized)
				return
			}

			// get the userId from the token claims
			userID, err := strconv.Atoi(claims.Subject)
			if err != nil {
				app.errorJson(w, errors.New("unknown user"), http.StatusUnauthorized)
				return
			}

			user, err := app.DB.GetUserById(int64(userID))

			if err != nil {
				app.errorJson(w, errors.New("unknown user"), http.StatusUnauthorized)
				return
			}

			// create a jwt user
			u := jwtUser{
				ID:        user.ID,
				FirstName: user.FirstName,
				LastName:  user.LastName,
			}

			//generate token
			tokens, err := app.auth.GenerateTokenPair(&u)

			if err != nil {
				app.errorJson(w, errors.New("error generating token"), http.StatusUnauthorized)
				return
			}

			http.SetCookie(w, app.auth.GetRefreshCookie(tokens.RefreshToken))

			// send token to user
			_ = app.writeJson(w, http.StatusAccepted, tokens)
		}

	}
}
