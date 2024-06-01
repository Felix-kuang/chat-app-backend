package main

import (
	"net/http"
	"postgres/chat/pkg/db"
	"postgres/chat/pkg/handler"

	"github.com/gorilla/mux"
)

func main() {
	//database connect
	if err := db.InitDB(); err != nil {
		panic(err)
	}
	defer db.DBClient.Close()

	//router creation
	router := mux.NewRouter()

	//user management
	router.HandleFunc("/api/register",handler.SignUpHandler).Methods("POST")
	router.HandleFunc("/api/login",handler.SignInHandler).Methods("POST")

	//TODO: CREATE CHAT MANAGEMENT
	//chat management
	// router.HandleFunc("/api/chats").Methods("GET")
	// router.HandleFunc("/api/chats").Methods("POST")
	// router.HandleFunc("/api/chats/{chat_id}").Methods("GET")

	//TODO: CREATE MESSAGE MANAGEMENT
	//message management
	// router.HandleFunc("/api/chats/{chat_id}/messages").Methods("GET")
	// router.HandleFunc("/api/chats/{chat_id}/messages").Methods("POST")

	//TODO: CREATE USER SELECTION
	//user selection
	// router.HandleFunc("/api/users").Methods("GET")

	//TODO: WEBSOCKET IMPLEMENTATION
	//ws
	// router.HandleFunc("/api/ws")

	http.ListenAndServe(":4869", router)
}

func TestHandler(w http.ResponseWriter, r http.Request) {}