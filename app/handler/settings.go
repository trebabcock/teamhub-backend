package handler

import (
	"fglhub-backend/app/model"
	"net/http"

	"github.com/jinzhu/gorm"
)

func GetUserSettings(db *gorm.DB, w http.ResponseWriter, r *http.Request) {

}

func UpdateUserSettings(db *gorm.DB, w http.ResponseWriter, r *http.Request) {

}

func GetDefaultSettings(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	defaultSettings := &model.Settings{
		Options: map[string]bool{"dark_mode": false, "notifications": true},
	}
	RespondJSON(w, http.StatusOK, defaultSettings)
}
