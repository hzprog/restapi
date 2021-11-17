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
	User "github.com/hzprog/restapi/Models/user"
	"golang.org/x/crypto/bcrypt"
)

//get all books
func Signup(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user User.User

	json.NewDecoder(r.Body).Decode(&user)
	dbconfig.Db.First(&user, "Username = ?", user.Username)
	if user.ID != 0 {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("username already exists")
		return
	}

	password := []byte(Env.GetEnvVar("SALT") + user.Password)

	// Hashing the password with the default cost of 10
	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	user.Password = string(hashedPassword)

	err = configdb.Db.Create(&user).Error
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("Error couldn't create the user")
		fmt.Println(err)
		return
	}

	tokenString, err := GenerateJWT(user.Username)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("Error couldn't create the user")
		fmt.Println(err)
		return
	}

	resp := make(map[string]string)
	resp["token"] = tokenString
	resp["message"] = "sign up successfully"
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		fmt.Printf("Error happened in JSON marshal. Err: %s", err)
	}

	w.Write(jsonResp)
}

//create a book
func Signin(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user User.User
	json.NewDecoder(r.Body).Decode(&user)
	password := []byte(Env.GetEnvVar("SALT") + user.Password)
	err := configdb.Db.Find(&user, "Username = ?", user.Username).Error
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("internal server error")
		fmt.Println(err)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), password)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusNonAuthoritativeInfo)
		json.NewEncoder(w).Encode("username or password is incorrect")
		return
	}
	tokenString, err := GenerateJWT(user.Username)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("Error couldn't create the user")
		fmt.Println(err)
		return
	}

	resp := make(map[string]string)
	resp["token"] = tokenString
	resp["message"] = "success"
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		fmt.Printf("Error happened in JSON marshal. Err: %s", err)
	}

	w.Write(jsonResp)
}

//update a book
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode("working fine")
}

//delete a book
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode("The user has been deleted successfully")
}

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

	return tokenString, nil
}
