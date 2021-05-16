package handler

import (
	"fglhub-backend/app/model"
	"net/http"

	"github.com/jinzhu/gorm"
)

func GetAllSpaces(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	spaces := []model.Space{}
	db.Find(&spaces)
	RespondJSON(w, http.StatusOK, spaces)
}

func getAllPublicSpaces(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	spaces := []model.Space{}
	db.Find(&spaces)
	RespondJSON(w, http.StatusOK, spaces)
}

func GetAllSpacesFromUser(db *gorm.DB, w http.ResponseWriter, r *http.Request) {

}

func GetSpace(db *gorm.DB, w http.ResponseWriter, r *http.Request) {

}

func CreateSpace(db *gorm.DB, w http.ResponseWriter, r *http.Request) {

}

func UpdateSpace(db *gorm.DB, w http.ResponseWriter, r *http.Request) {

}

func DeleteSpace(db *gorm.DB, w http.ResponseWriter, r *http.Request) {

}

func GetContent(db *gorm.DB, w http.ResponseWriter, r *http.Request) {

}

func UpdateContent(db *gorm.DB, w http.ResponseWriter, r *http.Request) {

}

func DeleteContent(db *gorm.DB, w http.ResponseWriter, r *http.Request) {

}
