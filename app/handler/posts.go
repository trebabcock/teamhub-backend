package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"teamhub-backend/app/model"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

func GetAllPosts(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	posts := []model.Post{}
	db.Find(&posts)
	RespondJSON(w, http.StatusOK, posts)
}

func GetPost(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	post := getPostByID(db, id, w, r)
	if post == nil {
		return
	}
	RespondJSON(w, http.StatusOK, post)
}

func GetAllPostsFromUser(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	user_id := vars["user_id"]
	posts := getPostsFromUser(db, user_id, w, r)
	if posts == nil {
		return
	}
	RespondJSON(w, http.StatusOK, posts)
}

func CreatePost(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	npost := model.NPost{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&npost); err != nil {
		RespondError(w, http.StatusBadRequest, "")
		return
	}
	defer r.Body.Close()

	post := model.Post{
		Author:   npost.Author,
		Type:     npost.Type,
		Time:     npost.Time,
		Content:  npost.Content,
		UUID:     npost.UUID,
		Comments: []model.Comment{},
	}

	if err := db.Save(&post).Error; err != nil {
		RespondError(w, http.StatusInternalServerError, "")
		return
	}
	RespondJSON(w, http.StatusCreated, post)
}

func UpdatePost(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	post := getPostByID(db, id, w, r)
	if post == nil {
		return
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&post); err != nil {
		RespondError(w, http.StatusBadRequest, "")
		return
	}
	defer r.Body.Close()

	if err := db.Save(&post).Error; err != nil {
		RespondError(w, http.StatusInternalServerError, "")
		return
	}
	RespondJSON(w, http.StatusOK, post)
}

func DeletePost(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	post := getPostByID(db, id, w, r)
	if post == nil {
		return
	}
	if err := db.Delete(&post).Error; err != nil {
		RespondError(w, http.StatusInternalServerError, "")
		return
	}
	RespondJSON(w, http.StatusNoContent, nil)
}

func getPostByID(db *gorm.DB, id string, w http.ResponseWriter, r *http.Request) *model.Post {
	pid, err := strconv.Atoi(id)
	if err != nil {
		return nil
	}
	post := model.Post{}
	if err := db.First(&post, pid).Error; err != nil {
		RespondError(w, http.StatusNotFound, err.Error())
		return nil
	}
	return &post
}

func getPostsFromUser(db *gorm.DB, id string, w http.ResponseWriter, r *http.Request) *[]model.Post {
	/*uid, err := strconv.Atoi(id)
	if err != nil {
		return nil
	}*/
	posts := []model.Post{}
	/*if err := db.Find(&posts, model.Post{}).Error; err != nil {
		RespondError(w, http.StatusNotFound, "")
		return nil
	}*/
	return &posts
}
