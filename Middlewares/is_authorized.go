package middlewares

import (
	"fmt"
	"net/http"

	"strings"

	jwt "github.com/dgrijalva/jwt-go"
	Env "github.com/hzprog/restapi/Helpers"
)

var mySigningKey = []byte(Env.GetEnvVar("SIGNED_STRING"))

func IsAuthorized(endpoint func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		if r.Header["Authorization"] != nil {
			reqToken := r.Header.Get("Authorization")
			splitToken := strings.Split(reqToken, "Bearer ")
			reqToken = splitToken[1]
			fmt.Println(r.Header["Authorization"])
			token, err := jwt.Parse(reqToken, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("There was an error")
				}
				return mySigningKey, nil
			})

			if err != nil {
				fmt.Fprintf(w, err.Error())
			}

			if token.Valid {
				endpoint(w, r)
			}
		} else {

			fmt.Fprintf(w, "Not Authorized")
		}
	}
}
