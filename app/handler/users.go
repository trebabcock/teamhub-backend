package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	model "teamhub-backend/app/model"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
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
		Roles:    user.Roles,
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
	credentials := model.Credentials{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&credentials); err != nil {
		RespondError(w, http.StatusUnauthorized, "")
		return
	}

	user := getUserByID(db, credentials.Username, w, r)
	if user == nil {
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(credentials.Password)); err != nil {
		RespondError(w, http.StatusUnauthorized, "")
		return
	}

	RespondJSON(w, http.StatusOK, user.Username)
}

// RegisterUser handles user registration
func RegisterUser(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	fmt.Println("register")
	user := model.User{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		RespondError(w, http.StatusBadRequest, "")
		return
	}
	defer r.Body.Close()

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 8)
	if err != nil {
		RespondError(w, http.StatusInternalServerError, "")
		return
	}
	user.Password = string(hashedPassword)
	if err := db.Save(&user).Error; err != nil {
		RespondError(w, http.StatusInternalServerError, "")
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
