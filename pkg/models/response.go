package models

type Response struct {
	Token  string `json:"token,omitempty"`
	Error  string `json:"error,omitempty"`
	UserID int    `json:"user_id,omitempty"`
}
