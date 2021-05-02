package handler

import (
	"encoding/json"
	model "fglhub-backend/app/model"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/jinzhu/gorm"
)

func GetAllMessages(db *gorm.DB, w http.ResponseWriter, r *http.Request) {

}

func GetMessage(db *gorm.DB, w http.ResponseWriter, r *http.Request) {

}

func UpdateMessage(db *gorm.DB, w http.ResponseWriter, r *http.Request) {

}

func DeleteMessage(db *gorm.DB, w http.ResponseWriter, r *http.Request) {

}

func GetAllChannels(db *gorm.DB, w http.ResponseWriter, r *http.Request) {

}

func GetChannel(db *gorm.DB, w http.ResponseWriter, r *http.Request) {

}

func CreateChannel(db *gorm.DB, w http.ResponseWriter, r *http.Request) {

}

func UpdateChannel(db *gorm.DB, w http.ResponseWriter, r *http.Request) {

}

func DeleteChannel(db *gorm.DB, w http.ResponseWriter, r *http.Request) {

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
