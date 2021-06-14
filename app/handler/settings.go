package handler

import (
	"net/http"
	"teamhub-backend/app/model"

	"github.com/jinzhu/gorm"
)

func GetUserSettings(db *gorm.DB, w http.ResponseWriter, r *http.Request) {

}

func UpdateUserSettings(db *gorm.DB, w http.ResponseWriter, r *http.Request) {

}

func GetDefaultSettings(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	settings := defaultSettings()
	RespondJSON(w, http.StatusOK, settings)
}

func defaultSettings() []model.Settings {
	return []model.Settings{
		{Option: "dark_mode", Enabled: false},
		{Option: "notifications", Enabled: true},
	}
}
