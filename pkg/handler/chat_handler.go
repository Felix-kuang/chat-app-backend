package handler

import (
	"encoding/json"
	"net/http"
	"postgres/chat/pkg/db"
)

func CreateChatHandler(w http.ResponseWriter, r *http.Request) {
	
}

func GetChatHandler(w http.ResponseWriter, r *http.Request) {
	// Extract user ID from the context set by the JWT middleware
	userID, ok := r.Context().Value("userID").(int)
	if !ok {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "user ID not found in context"})
		return
	}

	// Fetch chat list for the user
	chat_list, err := db.GetChats(userID)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(chat_list)
}

func GetChatInfoHandler(w http.ResponseWriter, r *http.Request) {

}
