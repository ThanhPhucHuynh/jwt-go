package handler

import (
	"encoding/json"
	"fmt"
	repoImpl "jwt-go/repository/repoimpl"
	driver "jwt-go/src/driver"
	models "jwt-go/src/model/user"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("thanhphuchuynh1999")

type Claims struct {
	Email string `json:"email"`
	DisplayName string `json:"display"`
	jwt.StandardClaims
}

func Register(w http.ResponseWriter, r *http.Request){
	var regData models.RegistrationData
	err := json.NewDecoder(r.Body).Decode(&regData)
	if err != nil {
		ResponseErr(w, http.StatusBadRequest)
		return
	}

	_, err = repoImpl.NewUserRepo(driver.Mongo.Client.Database("database_test")).FindUserByEmail(regData.Email)
	if err != models.ERR_USER_NOT_FOUND {
			ResponseErr(w, http.StatusConflict)
			return
		}

	user := models.User{
		Email:       regData.Email,
		Password:    regData.Password,
		DisplayName: regData.DisplayName,
	}
	err = repoImpl.NewUserRepo(driver.Mongo.Client.Database("database_test")).Insert(user)
	
	if err != nil {
		ResponseErr(w, http.StatusInternalServerError)
		return
	} 

	var tokenString string
	tokenString, err = GenToken(user)
	if err != nil {
		ResponseErr(w, http.StatusInternalServerError)
		return
	}

	ResponseSusscessfully(w, models.RegisterResponse{
		Token:  tokenString,
		Status: http.StatusOK,
	})
}

func Login(w http.ResponseWriter, r *http.Request){
	fmt.Println("Login")
	var loginData models.LoginData

	err := json.NewDecoder(r.Body).Decode(&loginData)
	fmt.Println(err)
	fmt.Println(loginData)
	if err != nil {
		ResponseErr(w, http.StatusBadRequest,)
		return
	}

	var user models.User
	user, err = repoImpl.NewUserRepo(driver.Mongo.Client.Database("database_test")).CheckLoginInfo(loginData.Email,loginData.Password)
	if  err != nil {
		ResponseErr(w, http.StatusUnauthorized)
		return
	}
	var tokenString string
	tokenString, err = GenToken(user)
	if err != nil {
		ResponseErr(w, http.StatusInternalServerError)
		return
	}
	// var user_get 
	user_get := models.UserGet{ID: user.ID, Email: user.Email, DisplayName: user.DisplayName}
	
	ResponseSusscessfully(w,models.RegisterResponse{
		User: user_get,
		Token: tokenString,
		Status: http.StatusOK,
	})
}

func ResponseErr(w http.ResponseWriter, statusCode int)  {
	jData, err := json.Marshal(models.Error{
		Status:  statusCode,
		Message: http.StatusText(statusCode),
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jData)
}

func ResponseSusscessfully(w http.ResponseWriter, data interface{})  {
	if data == nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	jData, err := json.Marshal(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jData)
}
func GenToken(user models.User) (string, error) {
	expirationTime := time.Now().Add(120 * time.Second)
	claims := &Claims{
		Email:       user.Email,
		DisplayName: user.DisplayName,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}