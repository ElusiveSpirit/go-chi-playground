package auth

import (
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"log"
	"net/http"
)

type Repository struct{}

type Controller struct {
	Repository Repository
}

// Get Authentication token POST /
func (c *Controller) GetToken(w http.ResponseWriter, req *http.Request) {
	var user User
	_ = json.NewDecoder(req.Body).Decode(&user)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Username,
		"password": user.Password,
	})

	log.Println("Username: " + user.Username)
	log.Println("Password: " + user.Password)
	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		fmt.Println(err)
	}
	json.NewEncoder(w).Encode(JwtToken{Token: tokenString})
}
