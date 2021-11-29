package user

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	configdb "github.com/hzprog/restapi/DBConfig"
	dbconfig "github.com/hzprog/restapi/DBConfig"
	Env "github.com/hzprog/restapi/Helpers"
	Response "github.com/hzprog/restapi/Helpers"
	User "github.com/hzprog/restapi/Models/user"
	"golang.org/x/crypto/bcrypt"
)

// swagger:route POST /signup Auth signup
// signup to the api.
// responses:
//   200: authResponse
//
// swagger:response authResponse

//sign up
func Signup(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user User.User

	json.NewDecoder(r.Body).Decode(&user)
	dbconfig.Db.First(&user, "Username = ?", user.Username)
	if user.ID != 0 {
		Response.HttpError(w, http.StatusInternalServerError, "username already exists", nil)
		return
	}

	password := []byte(Env.GetEnvVar("SALT") + user.Password)

	// Hashing the password with the default cost of 10
	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		Response.HttpError(w, http.StatusInternalServerError, "Error couldn't create the user", err)
	}
	user.Password = string(hashedPassword)

	err = configdb.Db.Create(&user).Error
	if err != nil {
		Response.HttpError(w, http.StatusInternalServerError, "Error couldn't create the user", err)
		return
	}

	tokenString, err := GenerateJWT(user.Username)
	if err != nil {
		Response.HttpError(w, http.StatusInternalServerError, "Error couldn't create the user", err)
		return
	}

	Response.HttpResponse(w, http.StatusOK, tokenString)
}

// swagger:route POST /login Auth login
// login to the book api.
// responses:
//   200: authResponse
//
// swagger:response authResponse

//create a book
func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user User.User
	json.NewDecoder(r.Body).Decode(&user)
	password := []byte(Env.GetEnvVar("SALT") + user.Password)
	err := configdb.Db.Find(&user, "Username = ?", user.Username).Error
	if err != nil {
		Response.HttpError(w, http.StatusInternalServerError, "internal server error", err)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), password)
	if err != nil {
		Response.HttpError(w, http.StatusForbidden, "username or password is incorrect", err)
		return
	}
	tokenString, err := GenerateJWT(user.Username)
	if err != nil {
		Response.HttpError(w, http.StatusInternalServerError, "Error couldn't create the user", err)
		return
	}

	Response.HttpResponse(w, http.StatusOK, tokenString)
}

//Todo: implements UpdateUser
// func UpdateUser(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")

// 	json.NewEncoder(w).Encode("working fine")
// }

//Todo: implements DeleteUser
// func DeleteUser(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")

// 	json.NewEncoder(w).Encode("The user has been deleted successfully")
// }

func GenerateJWT(username string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS512)

	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["client"] = username
	claims["intial_time"] = time.Now().Unix()
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString([]byte(Env.GetEnvVar("SIGNED_STRING")))

	if err != nil {
		fmt.Printf("Something Went Wrong: %s", err.Error())
		return "", err
	}

	tokenString = "Bearer " + tokenString

	return tokenString, nil
}
