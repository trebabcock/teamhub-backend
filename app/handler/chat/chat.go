package chat

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	handler "teamhub-backend/app/handler"
	model "teamhub-backend/app/model"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

func GetAllMessages(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	messages := []model.Message{}
	db.Find(&messages)
	handler.RespondJSON(w, http.StatusOK, messages)
}

func GetSomeMessages(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	limit, err := strconv.Atoi(r.URL.Query()["limit"][0])
	if err != nil {
		handler.RespondError(w, http.StatusInternalServerError, "")
		return
	}
	offset, err := strconv.Atoi(r.URL.Query()["offset"][0])
	if err != nil {
		handler.RespondError(w, http.StatusInternalServerError, "")
		return
	}
	messages := []model.Message{}
	db.Offset(offset).Limit(limit).Find(&messages)
	handler.RespondJSON(w, http.StatusOK, messages)
}

func GetMessage(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	message := getMessageByID(db, id, w, r)
	if message == nil {
		return
	}
	handler.RespondJSON(w, http.StatusOK, message)
}

func UpdateMessage(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]
	message := getMessageByID(db, id, w, r)
	if message == nil {
		return
	}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&message); err != nil {
		handler.RespondError(w, http.StatusBadRequest, "")
		return
	}
	defer r.Body.Close()

	if err := db.Save(&message).Error; err != nil {
		handler.RespondError(w, http.StatusInternalServerError, "")
		return
	}
	handler.RespondJSON(w, http.StatusOK, message)
}

func DeleteMessage(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	message := getMessageByID(db, id, w, r)
	if message == nil {
		return
	}
	if err := db.Delete(&message).Error; err != nil {
		handler.RespondError(w, http.StatusInternalServerError, "")
		return
	}
	handler.RespondJSON(w, http.StatusNoContent, nil)
}

func GetAllChannels(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	channels := []model.Channel{}
	db.Find(&channels)
	handler.RespondJSON(w, http.StatusOK, channels)
}

func GetChannel(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	channel := getChannelByID(db, id, w, r)
	if channel == nil {
		return
	}
	handler.RespondJSON(w, http.StatusOK, channel)
}

func CreateChannel(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	log.Println("Create Channel")
	channel := model.Channel{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&channel); err != nil {
		handler.RespondError(w, http.StatusBadRequest, "")
		return
	}
	defer r.Body.Close()

	if err := db.Save(&channel).Error; err != nil {
		handler.RespondError(w, http.StatusInternalServerError, "")
		return
	}
	handler.RespondJSON(w, http.StatusCreated, channel)
}

func UpdateChannel(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	channel := getChannelByID(db, id, w, r)
	if channel == nil {
		return
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&channel); err != nil {
		handler.RespondError(w, http.StatusBadRequest, "")
		return
	}
	defer r.Body.Close()

	if err := db.Save(&channel).Error; err != nil {
		handler.RespondError(w, http.StatusInternalServerError, "")
		return
	}
	handler.RespondJSON(w, http.StatusOK, channel)
}

func DeleteChannel(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	channel := getChannelByID(db, id, w, r)
	if channel == nil {
		return
	}
	if err := db.Delete(&channel).Error; err != nil {
		handler.RespondError(w, http.StatusInternalServerError, "")
		return
	}
	handler.RespondJSON(w, http.StatusNoContent, nil)
}

func WsHandler(db *gorm.DB, hub *Hub, w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	id := r.URL.Query().Get("id")

	client := &Client{hub: hub, conn: conn, send: make(chan []byte, 256), id: id}
	client.hub.register <- client

	// Allow collection of memory referenced by the caller by doing all work in
	// new goroutines.
	go client.writePump()
	go client.readPump(db)
}

func getMessageByID(db *gorm.DB, id string, w http.ResponseWriter, r *http.Request) *model.Message {
	mid, err := strconv.Atoi(id)
	if err != nil {
		return nil
	}
	message := model.Message{}
	if err := db.First(&message, mid).Error; err != nil {
		handler.RespondError(w, http.StatusNotFound, err.Error())
		return nil
	}
	return &message
}

func getChannelByID(db *gorm.DB, id string, w http.ResponseWriter, r *http.Request) *model.Channel {
	channel := model.Channel{}
	if err := db.First(&channel, &model.Channel{UUID: id}).Error; err != nil {
		handler.RespondError(w, http.StatusNotFound, err.Error())
		return nil
	}
	return &channel
}
