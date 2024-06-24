package db

import (
	"fmt"
	"postgres/chat/pkg/models"
)

func CreateChats(userID int, targetUserID int) (int, error) {
	var chat models.Chats

	query := `
		INSERT INTO chats
	`
	return 0, nil
}

func GetChats(userID int) ([]models.Chats, error) {
	var chats []models.Chats

	query := `
		SELECT A.id, A.type_chat
		FROM chats A 
		LEFT JOIN chat_members B ON A.id = B.chat_id
		WHERE A.type_chat = 'public' OR (A.type_chat = 'private' AND B.user_id = $1)
	`
	rows, err := DBClient.Query(query, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve chat list: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var chat models.Chats
		if err := rows.Scan(&chat.ID, &chat.ChatType); err != nil {
			return nil, fmt.Errorf("failed to scan chat: %w", err)
		}
		chats = append(chats, chat)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error during row iteration: %w", err)
	}

	return chats, nil
}

func GetChatInfo() {}
