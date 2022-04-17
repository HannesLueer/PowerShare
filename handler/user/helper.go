package user

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"os"
	"strconv"
	"time"
)

func generateHashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func generateJWT(email string, role int64) (string, error) {
	var mySigningKey = []byte(os.Getenv("SIGNING_KEY"))
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["email"] = email
	claims["role"] = role
	minutes, err := strconv.Atoi(os.Getenv("SERVER_PORT"))
	if err != nil {
		minutes = 30
	}
	claims["exp"] = time.Now().Add(time.Minute * time.Duration(minutes)).Unix()

	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		fmt.Errorf("something went wrong: %s", err.Error())
		return "", err
	}
	return tokenString, nil
}

func IsAuthorized(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if r.Header["Token"] == nil {
			http.Error(w, "no token found", http.StatusBadRequest)
			return
		}

		var mySigningKey = []byte(os.Getenv("SIGNING_KEY"))

		token, err := jwt.Parse(r.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("there was an error in parsing")
			}
			return mySigningKey, nil
		})

		if err != nil {
			http.Error(w, "Your Token has been expired", http.StatusUnauthorized)
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			r.Header.Set("Role", fmt.Sprintf("%v", claims["role"]))
			handler.ServeHTTP(w, r)
			return
		}
		http.Error(w, "Not authorized", http.StatusUnauthorized)
	}
}

func GetEmailFromToken(tokenStr string) (string, error) {
	var mySigningKey = []byte(os.Getenv("SIGNING_KEY"))
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("there was an error in parsing")
		}
		return mySigningKey, nil
	})

	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%v", token.Claims.(jwt.MapClaims)["email"]), nil
}

func GetToken(r *http.Request) (token string, errCode int, err error) {
	if r.Header["Token"] == nil {
		return "", http.StatusBadRequest, fmt.Errorf("no token found")
	}
	var tokenStr = r.Header["Token"][0]
	return tokenStr, http.StatusOK, nil
}
