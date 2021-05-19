package app

import (
	"fmt"
	"time"

	//"io"
	"log"
	"net/http"

	//"os"
	//"strconv"

	handler "fglhub-backend/app/handler"
	chat "fglhub-backend/app/handler/chat"
	"fglhub-backend/app/model"
	db "fglhub-backend/db"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

type App struct {
	Router *mux.Router
	API    *mux.Router
	DB     *gorm.DB
	Hub    *chat.Hub
}

func (a *App) Init(dbConfig *db.Config) {
	dbURI := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s connect_timeout=15",
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.Username,
		dbConfig.Name,
		dbConfig.Password,
		dbConfig.SSLMode,
	)
	fmt.Println("Connecting to database...")
	db, err := gorm.Open(dbConfig.Dialect, dbURI)
	if err != nil {
		log.Println("Could not connect to database")
		log.Fatal(err.Error())
	}

	fmt.Println("Connected to database")
	fmt.Println("Migrating database...")
	a.DB = model.DBMigrate(db)
	fmt.Println("Migrated database")
	a.Router = mux.NewRouter()
	a.Router.Use(a.Default)
	a.API = a.Router.PathPrefix("/api").Subrouter()
	//a.API.Use(a.Authenticate)
	a.Hub = chat.NewHub()
	a.setv1Routes()
}

func (a *App) setv1Routes() {
	a.get("/api/v1/users", a.getAllUsers)
	a.get("/api/v1/users/{id}", a.getUser)
	a.put("/api/v1/users/{id}", a.updateUser)
	a.delete("/api/v1/users/{id}", a.deleteUser)

	a.post("/auth/register", a.registerUser)
	a.post("/auth/login", a.userLogin)

	a.get("/api/v1/spaces", a.getAllSpaces)
	a.get("/api/v1/spaces/{user_id}", a.getAllSpacesFromUser)
	a.get("/api/v1/spaces/{id}", a.getSpace)
	a.post("/api/v1/spaces/create", a.createSpace)
	a.put("/api/v1/spaces/{id}", a.updateSpace)
	a.delete("/api/v1/spaces/{id}", a.deleteSpace)
	a.get("/api/v1/spaces/{id}/content/{id}", a.getContent)
	a.put("/api/v1/spaces/{id}/content/{id}", a.updateContent)
	a.delete("/api/v1/spaces/{id}/content/{id}", a.deleteContent)

	a.get("/api/v1/posts", a.getAllPosts)
	a.get("/api/v1/posts/{id}", a.getPost)
	a.get("/api/v1/posts/{user_id}", a.getAllPostsFromUser)
	a.post("/api/v1/posts/create", a.createPost)
	a.put("/api/v1/posts/{id}", a.updatePost)
	a.delete("/api/v1/posts/{id}", a.deletePost)

	a.get("/api/v1/chat/messages", a.getAllMessages)
	a.get("/api/v1/chat/messages/pagination", a.getSomeMessages)
	a.get("/api/v1/chat/messages/{id}", a.getMessage)
	a.put("/api/v1/chat/messages/{id}", a.updateMessage)
	a.delete("/api/v1/chat/messages/{id}", a.deleteMessage)
	a.get("/api/v1/chat/channels", a.getAllChannels)
	a.post("/api/v1/chat/channels/new", a.createChannel)
	a.get("/api/v1/chat/channels/{id}", a.getChannel)
	a.put("/api/v1/chat/channels/{id}", a.updateChannel)
	a.delete("/api/v1/chat/channels/{id}", a.deleteChannel)

	a.get("/api/v1/settings/{user_id}", a.getUserSettings)
	a.put("/api/v1/settings/{user_id}", a.updateUserSettings)
	a.get("/api/v1/settings/default", a.getDefaultSettings)

	a.get("/api/v1/roles", a.getAllRoles)
	a.post("/api/v1/roles/create", a.createRole)
	a.get("/api/v1/role/{id}", a.getRole)
	a.put("/api/v1/role/{id}", a.updateRole)
	a.delete("/api/v1/role/{id}", a.deleteRole)

	a.handle("/api/v1/chat", a.wsHandler)
}

func (a *App) get(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("GET")
}

func (a *App) post(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("POST")
}
func (a *App) put(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("PUT")
}

func (a *App) delete(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("DELETE")
}

func (a *App) Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "task failed succesfully", http.StatusOK)
	})
}

func (a *App) Default(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		next.ServeHTTP(w, r)
	})
}

func (a *App) handle(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f)
}

func (a *App) wsHandler(w http.ResponseWriter, r *http.Request) {
	chat.WsHandler(a.DB, a.Hub, w, r)
}

func (a *App) getAllUsers(w http.ResponseWriter, r *http.Request) {
	handler.GetAllUsers(a.DB, w, r)
}

func (a *App) getUser(w http.ResponseWriter, r *http.Request) {
	handler.GetUser(a.DB, w, r)
}

func (a *App) registerUser(w http.ResponseWriter, r *http.Request) {
	handler.RegisterUser(a.DB, w, r)
}

func (a *App) userLogin(w http.ResponseWriter, r *http.Request) {
	handler.UserLogin(a.DB, w, r)
}

func (a *App) updateUser(w http.ResponseWriter, r *http.Request) {
	handler.UpdateUser(a.DB, w, r)
}

func (a *App) deleteUser(w http.ResponseWriter, r *http.Request) {
	handler.DeleteUser(a.DB, w, r)
}

func (a *App) getAllSpaces(w http.ResponseWriter, r *http.Request) {
	handler.GetAllSpaces(a.DB, w, r)
}

func (a *App) getAllSpacesFromUser(w http.ResponseWriter, r *http.Request) {
	handler.GetAllSpacesFromUser(a.DB, w, r)
}

func (a *App) getSpace(w http.ResponseWriter, r *http.Request) {
	handler.GetSpace(a.DB, w, r)
}

func (a *App) createSpace(w http.ResponseWriter, r *http.Request) {
	handler.CreateSpace(a.DB, w, r)
}

func (a *App) updateSpace(w http.ResponseWriter, r *http.Request) {
	handler.UpdateSpace(a.DB, w, r)
}

func (a *App) deleteSpace(w http.ResponseWriter, r *http.Request) {
	handler.DeleteSpace(a.DB, w, r)
}

func (a *App) getContent(w http.ResponseWriter, r *http.Request) {
	handler.GetContent(a.DB, w, r)
}

func (a *App) updateContent(w http.ResponseWriter, r *http.Request) {
	handler.UpdateContent(a.DB, w, r)
}

func (a *App) deleteContent(w http.ResponseWriter, r *http.Request) {
	handler.DeleteContent(a.DB, w, r)
}

func (a *App) getAllPosts(w http.ResponseWriter, r *http.Request) {
	handler.GetAllPosts(a.DB, w, r)
}

func (a *App) getPost(w http.ResponseWriter, r *http.Request) {
	handler.GetPost(a.DB, w, r)
}

func (a *App) getAllPostsFromUser(w http.ResponseWriter, r *http.Request) {
	handler.GetAllPostsFromUser(a.DB, w, r)
}

func (a *App) createPost(w http.ResponseWriter, r *http.Request) {
	handler.CreatePost(a.DB, w, r)
}

func (a *App) updatePost(w http.ResponseWriter, r *http.Request) {
	handler.UpdatePost(a.DB, w, r)
}

func (a *App) deletePost(w http.ResponseWriter, r *http.Request) {
	handler.DeletePost(a.DB, w, r)
}

func (a *App) getAllMessages(w http.ResponseWriter, r *http.Request) {
	chat.GetAllMessages(a.DB, w, r)
}

func (a *App) getMessage(w http.ResponseWriter, r *http.Request) {
	chat.GetMessage(a.DB, w, r)
}

func (a *App) updateMessage(w http.ResponseWriter, r *http.Request) {
	chat.UpdateMessage(a.DB, w, r)
}

func (a *App) deleteMessage(w http.ResponseWriter, r *http.Request) {
	chat.DeleteMessage(a.DB, w, r)
}

func (a *App) getAllChannels(w http.ResponseWriter, r *http.Request) {
	chat.GetAllChannels(a.DB, w, r)
}

func (a *App) getSomeMessages(w http.ResponseWriter, r *http.Request) {
	chat.GetSomeMessages(a.DB, w, r)
}

func (a *App) getChannel(w http.ResponseWriter, r *http.Request) {
	chat.GetChannel(a.DB, w, r)
}

func (a *App) createChannel(w http.ResponseWriter, r *http.Request) {
	chat.CreateChannel(a.DB, w, r)
}

func (a *App) updateChannel(w http.ResponseWriter, r *http.Request) {
	chat.UpdateChannel(a.DB, w, r)
}

func (a *App) deleteChannel(w http.ResponseWriter, r *http.Request) {
	chat.DeleteChannel(a.DB, w, r)
}

func (a *App) getUserSettings(w http.ResponseWriter, r *http.Request) {
	handler.GetUserSettings(a.DB, w, r)
}

func (a *App) updateUserSettings(w http.ResponseWriter, r *http.Request) {
	handler.UpdateUserSettings(a.DB, w, r)
}

func (a *App) getDefaultSettings(w http.ResponseWriter, r *http.Request) {
	handler.GetDefaultSettings(a.DB, w, r)
}

func (a *App) getAllRoles(w http.ResponseWriter, r *http.Request) {
	handler.GetAllRoles(a.DB, w, r)
}

func (a *App) createRole(w http.ResponseWriter, r *http.Request) {
	handler.CreateRole(a.DB, w, r)
}

func (a *App) getRole(w http.ResponseWriter, r *http.Request) {
	handler.GetRole(a.DB, w, r)
}

func (a *App) updateRole(w http.ResponseWriter, r *http.Request) {
	handler.UpdateRole(a.DB, w, r)
}

func (a *App) deleteRole(w http.ResponseWriter, r *http.Request) {
	handler.DeleteRole(a.DB, w, r)
}

// Run starts the server
func (a *App) Run(host string) {
	go a.Hub.Run()
	server := &http.Server{
		Addr:         host,
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      a.Router,
	}
	fmt.Println("Server running at", host)
	log.Fatal(server.ListenAndServe())
}
