package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"teamhub-backend/app/model"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

func GetAllRoles(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	roles := []model.Role{}
	db.Find(&roles)
	RespondJSON(w, http.StatusOK, roles)
}

func CreateRole(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	role := model.Role{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&role); err != nil {
		RespondError(w, http.StatusBadRequest, "")
		return
	}
	defer r.Body.Close()

	if err := db.Save(&role).Error; err != nil {
		RespondError(w, http.StatusInternalServerError, "")
		return
	}
	RespondJSON(w, http.StatusCreated, role)
}

func GetRole(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	role := getRoleByID(db, id, w, r)
	if role == nil {
		return
	}
	RespondJSON(w, http.StatusOK, role)
}

func UpdateRole(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	role := getRoleByID(db, id, w, r)
	if role == nil {
		return
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&role); err != nil {
		RespondError(w, http.StatusBadRequest, "")
		return
	}
	defer r.Body.Close()

	if err := db.Save(&role).Error; err != nil {
		RespondError(w, http.StatusInternalServerError, "")
		return
	}
	RespondJSON(w, http.StatusOK, role)
}

func DeleteRole(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	role := getRoleByID(db, id, w, r)
	if role == nil {
		return
	}
	if err := db.Delete(&role).Error; err != nil {
		RespondError(w, http.StatusInternalServerError, "")
		return
	}
	RespondJSON(w, http.StatusNoContent, nil)
}

func getRoleByID(db *gorm.DB, id string, w http.ResponseWriter, r *http.Request) *model.Role {
	rid, err := strconv.Atoi(id)
	if err != nil {
		return nil
	}
	role := model.Role{}
	if err := db.First(&role, rid).Error; err != nil {
		RespondError(w, http.StatusNotFound, err.Error())
		return nil
	}
	return &role
}
