package auth

import (
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
	"todo-list/db"
	"todo-list/user"
)

type Handler struct {
	Email string `json:"email"`
	ID    uint   `json:"id"`
	jwt.StandardClaims
}

func (h *Handler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	var u user.User
	_ = json.NewDecoder(r.Body).Decode(&u)

	foundUser := db.Db.Where("email = ?", u.Email).First(&u)

	if foundUser.Error != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Email is already registered")
		return
	}

	hashedPass, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	u.Password = string(hashedPass)

	result := db.Db.Create(&u)

	if result.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(result.Error)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(u)

}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var u user.User
	_ = json.NewDecoder(r.Body).Decode(&u)

	var password = u.Password

	result := db.Db.Where("email = ?", u.Email).First(&u)

	if result.Error != nil {
		fmt.Println("No user")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("Invalid credential")
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))

	if err != nil {
		fmt.Println("Passwor hashing priblem", password, u.Password)
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("Invalid credentials")
		return
	}

	expirationTime := time.Now().Add(24 * time.Hour)
	h.Email = u.Email
	h.ID = u.ID
	h.StandardClaims = jwt.StandardClaims{
		ExpiresAt: expirationTime.Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, h)

	tokenString, err := token.SignedString([]byte("your-secret-key"))

	if err != nil {
		fmt.Println("token error")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("Error generating token")
		return
	}

	response := map[string]string{
		"token": tokenString,
	}

	json.NewEncoder(w).Encode(response)
}
