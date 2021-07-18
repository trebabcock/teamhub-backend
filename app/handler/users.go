package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	model "teamhub-backend/app/model"

	"github.com/gbrlsnchs/jwt/v3"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/joho/godotenv/autoload"
	"golang.org/x/crypto/bcrypt"
)

// GetUser returns a user
func GetUser(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	user := getUserByID(db, id, w, r)
	if user == nil {
		return
	}
	RespondJSON(w, http.StatusOK, user)
}

func getUserPrivate(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	user := getUserByID(db, id, w, r)
	if user == nil {
		return
	}
	RespondJSON(w, http.StatusOK, user)
}

func getUserPublic(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	user := getUserByID(db, id, w, r)
	if user == nil {
		return
	}
	public := model.PublicUser{
		Name:     user.Name,
		Username: user.Username,
		UUID:     user.UUID,
	}
	RespondJSON(w, http.StatusOK, public)
}

// GetAllUsers returns all users
func GetAllUsers(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	users := []model.User{}
	db.Find(&users)
	RespondJSON(w, http.StatusOK, users)
}

// UserLogin handles user login
func UserLogin(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	fmt.Println("login")

	t := r.URL.Query().Get("token")
	if t

	credentials := model.Credentials{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&credentials); err != nil {
		RespondError(w, http.StatusUnauthorized, "")
		return
	}

	user := getUserByName(db, credentials.Username, w, r)
	if user == nil {
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(credentials.Password)); err != nil {
		RespondError(w, http.StatusUnauthorized, "")
		return
	}

	token, err := generateToken(user.UUID)
	if err != nil {
		RespondError(w, http.StatusInternalServerError, "")
		return
	}

	lr := model.LoginResponse{
		Name:     user.Name,
		Username: user.Username,
		UUID:     user.UUID,
		Token:    token,
	}

	RespondJSON(w, http.StatusOK, lr)
}

// RegisterUser handles user registration
func RegisterUser(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	fmt.Println("register")
	ruser := model.RUser{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&ruser); err != nil {
		RespondError(w, http.StatusBadRequest, "")
		log.Println("decode:", err.Error())
		return
	}
	defer r.Body.Close()

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(ruser.Password), 8)
	if err != nil {
		RespondError(w, http.StatusInternalServerError, "")
		log.Println("hash:", err.Error())
		return
	}

	id, err := uuid.NewUUID()
	if err != nil {
		RespondError(w, http.StatusInternalServerError, "")
	}

	user := model.User{
		Name:     ruser.Name,
		Username: ruser.Username,
		Password: string(hashedPassword),
		UUID:     id.String(),
	}

	if err := db.Save(&user).Error; err != nil {
		RespondError(w, http.StatusInternalServerError, "")
		log.Println("save:", err.Error())
		return
	}
	RespondJSON(w, http.StatusCreated, user)
}

// UpdateUser updates user information for specified user
func UpdateUser(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]
	user := getUserByID(db, id, w, r)
	if user == nil {
		return
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		RespondError(w, http.StatusBadRequest, "")
		return
	}
	defer r.Body.Close()

	if err := db.Save(&user).Error; err != nil {
		RespondError(w, http.StatusInternalServerError, "")
		return
	}
	RespondJSON(w, http.StatusOK, user)
}

// DeleteUser deletes specified user
func DeleteUser(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]
	user := getUserByID(db, id, w, r)
	if user == nil {
		return
	}
	if err := db.Delete(&user).Error; err != nil {
		RespondError(w, http.StatusInternalServerError, "")
		return
	}
	RespondJSON(w, http.StatusNoContent, nil)
}

func getUserByID(db *gorm.DB, id string, w http.ResponseWriter, r *http.Request) *model.User {
	uid, err := strconv.Atoi(id)
	if err != nil {
		return nil
	}
	user := model.User{}
	if err := db.First(&user, uid).Error; err != nil {
		RespondError(w, http.StatusNotFound, "")
		return nil
	}
	return &user
}

func getUserByName(db *gorm.DB, username string, w http.ResponseWriter, r *http.Request) *model.User {
	user := model.User{}
	if err := db.First(&user, model.User{Username: username}).Error; err != nil {
		RespondError(w, http.StatusNotFound, "")
		return nil
	}
	return &user
}

func generateToken(userID string) (string, error) {
	secret := os.Getenv("JWT_SECRET")
	hs := jwt.NewHS256([]byte(secret))

	id, err := uuid.NewUUID()
	if err != nil {
		return "", err
	}

	payload := jwt.Payload{
		Audience: jwt.Audience{username},
		JWTID:    id.String(),
	}

	token, err := jwt.Sign(payload, hs)
	if err != nil {
		return "", err
	}

	return string(token), nil
}

func verifyToken(token []byte) (bool, error) {
	secret := os.Getenv("JWT_SECRET")
	hs := jwt.NewHS256([]byte(secret))

	payload := jwt.Payload{}

	hd, err := jwt.Verify(token, hs, &payload)
	if err != nil {
		return false, err
	}
	jwt.Header
	hd.KeyID
	return true, nil
}
