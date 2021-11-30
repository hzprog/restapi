package user

import (
	"encoding/json"
	"errors"
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
	"gorm.io/gorm"
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
	user.Balance = 2000

	if err := configdb.Db.Create(&user).Error; err != nil {
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

	if err := configdb.Db.Find(&user, "Username = ?", user.Username).Error; err != nil {
		Response.HttpError(w, http.StatusInternalServerError, "internal server error", err)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), password); err != nil {
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

// swagger:route POST /transfer User transfer
// login to the book api.
// responses:
//   200: noContent
//
// swagger:response noContent
func TransferMoney(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	type Body struct {
		SenderId   int     `json:"sender_id"`
		ReceiverId int     `json:"receiver_id"`
		Amount     float64 `json:"amount"`
	}

	var body Body
	var sender User.User
	var receiver User.User

	json.NewDecoder(r.Body).Decode(&body)

	if body.Amount < 0 {
		Response.HttpError(w, http.StatusInternalServerError, "the amount you want to transfer is negative", errors.New("the amount you want to transfer is negative"))
		return
	}

	if err := configdb.Db.First(&sender, body.SenderId).Error; err != nil {
		Response.HttpError(w, http.StatusInternalServerError, "Can't find a sender with that id", err)
		return
	}

	if err := configdb.Db.First(&receiver, body.ReceiverId).Error; err != nil {
		Response.HttpError(w, http.StatusInternalServerError, "Can't find a receiver with that id", err)
		return
	}

	if err := transferMoneyTransaction(configdb.Db, sender, receiver, body.Amount); err != nil {
		Response.HttpError(w, http.StatusInternalServerError, err.Error(), err)
		return
	}

	Response.HttpResponse(w, http.StatusOK, "Transfer successded")
}

func transferMoneyTransaction(db *gorm.DB, sender User.User, receiver User.User, amount float64) error {
	// Note the use of tx as the database handle once you are within a transaction
	tx := db.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	if sender.Balance < amount {
		return errors.New("the transer amount exceeds the balance")
	}

	sender.Balance = sender.Balance - amount

	if err := tx.Save(&sender).Error; err != nil {
		tx.Rollback()
		return err
	}

	receiver.Balance = receiver.Balance + amount
	if err := tx.Save(&receiver).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
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
