package handler

import (
	"encoding/json"
	model "fglhub-backend/app/model"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"github.com/jinzhu/gorm"
)

func GetAllMessages(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	messages := []model.Message{}
	db.Find(&messages)
	RespondJSON(w, http.StatusOK, messages)
}

func GetSomeMessages(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	limit, err := strconv.Atoi(r.URL.Query()["limit"][0])
	if err != nil {
		RespondError(w, http.StatusInternalServerError, "")
		return
	}
	offset, err := strconv.Atoi(r.URL.Query()["offset"][0])
	if err != nil {
		RespondError(w, http.StatusInternalServerError, "")
		return
	}
	messages := []model.Message{}
	db.Offset(offset).Limit(limit).Find(&messages)
	RespondJSON(w, http.StatusOK, messages)
}

func GetMessage(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	message := getMessageByID(db, id, w, r)
	if message == nil {
		return
	}
	RespondJSON(w, http.StatusOK, message)
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
		RespondError(w, http.StatusBadRequest, "")
		return
	}
	defer r.Body.Close()

	if err := db.Save(&message).Error; err != nil {
		RespondError(w, http.StatusInternalServerError, "")
		return
	}
	RespondJSON(w, http.StatusOK, message)
}

func DeleteMessage(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	message := getMessageByID(db, id, w, r)
	if message == nil {
		return
	}
	if err := db.Delete(&message).Error; err != nil {
		RespondError(w, http.StatusInternalServerError, "")
		return
	}
	RespondJSON(w, http.StatusNoContent, nil)
}

func GetAllChannels(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	channels := []model.Channel{}
	db.Find(&channels)
	RespondJSON(w, http.StatusOK, channels)
}

func GetChannel(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	channel := getChannelByID(db, id, w, r)
	if channel == nil {
		return
	}
	RespondJSON(w, http.StatusOK, channel)
}

func CreateChannel(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	channel := model.Channel{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&channel); err != nil {
		RespondError(w, http.StatusBadRequest, "")
		return
	}
	defer r.Body.Close()

	if err := db.Save(&channel).Error; err != nil {
		RespondError(w, http.StatusInternalServerError, "")
		return
	}
	RespondJSON(w, http.StatusCreated, channel)
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
		RespondError(w, http.StatusBadRequest, "")
		return
	}
	defer r.Body.Close()

	if err := db.Save(&channel).Error; err != nil {
		RespondError(w, http.StatusInternalServerError, "")
		return
	}
	RespondJSON(w, http.StatusOK, channel)
}

func DeleteChannel(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	channel := getChannelByID(db, id, w, r)
	if channel == nil {
		return
	}
	if err := db.Delete(&channel).Error; err != nil {
		RespondError(w, http.StatusInternalServerError, "")
		return
	}
	RespondJSON(w, http.StatusNoContent, nil)
}

func WsHandler(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	var upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin:     func(r *http.Request) bool { return true },
	}

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}

	log.Println("Client connected")

	reader(db, ws)
}

func reader(db *gorm.DB, conn *websocket.Conn) {
	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}

		log.Println(string(p))
		saveMessage(db, p)

		if err := conn.WriteMessage(messageType, p); err != nil {
			log.Println(err)
			return
		}
	}
}

func saveMessage(db *gorm.DB, m []byte) {
	message := model.Message{}

	if err := json.Unmarshal(m, &message); err != nil {
		log.Println(err)
		return
	}

	if err := db.Save(&message).Error; err != nil {
		log.Println(err)
		return
	}
}

func getMessageByID(db *gorm.DB, id string, w http.ResponseWriter, r *http.Request) *model.Message {
	mid, err := strconv.Atoi(id)
	if err != nil {
		return nil
	}
	message := model.Message{}
	if err := db.First(&message, mid).Error; err != nil {
		RespondError(w, http.StatusNotFound, err.Error())
		return nil
	}
	return &message
}

func getChannelByID(db *gorm.DB, id string, w http.ResponseWriter, r *http.Request) *model.Channel {
	cid, err := strconv.Atoi(id)
	if err != nil {
		return nil
	}
	channel := model.Channel{}
	if err := db.First(&channel, cid).Error; err != nil {
		RespondError(w, http.StatusNotFound, err.Error())
		return nil
	}
	return &channel
}
