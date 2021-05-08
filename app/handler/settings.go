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
	defaultSettings := []model.Settings{
		{Option: "dark_mode", Enabled: false},
		{Option: "notifications", Enabled: true},
	}

	RespondJSON(w, http.StatusOK, defaultSettings)
}
