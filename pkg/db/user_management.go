package db

import (
	"fmt"
	"postgres/chat/pkg/jwt"
	"postgres/chat/pkg/models"

	"golang.org/x/crypto/bcrypt"
)

func SignUp(user models.User) (string, error) {
	var userID int

	hashedPassword, err := HashPassword(user.Password)
	if err != nil {
		return "", fmt.Errorf("failed to hash password: %w", err)
	}

	query := "INSERT INTO users (username,password) VALUES ($1,$2) RETURNING id"
	err = DBClient.QueryRow(query, user.Username, hashedPassword).Scan(&userID)
	if err != nil {
		return "", fmt.Errorf("failed to sign up: %w", err)
	}

	user.ID = userID

	token, err := SignIn(user)
	if err != nil {
		return "", fmt.Errorf("failed to sign in after sign up: %w", err)
	}

	return token, nil
}

func SignIn(user models.User) (string, error) {
	storedPassword, err := getPassword(user.Username)
	if err != nil {
		return "", fmt.Errorf("failed to sign in: %w", err)
	}

	match := CheckPassword(user.Password, storedPassword)
	if match {
		// Authentication Success, generate JWT
		token, err := jwt.GenerateToken(user.ID)
		if err != nil {
			return "", fmt.Errorf("failed to generate JWT token: %w", err)
		}

		return token, nil
	}
	return "", fmt.Errorf("invalid password")

}

func getPassword(username string) (string, error) {
	var password string
	err := DBClient.QueryRow("SELECT password FROM users where username = $1", username).Scan(&password)
	if err != nil {
		return "", err
	}
	return password, nil
}

// HashPassword hashes the given password using bcrypt
func HashPassword(password string) (string, error) {
	// Generate a salt and hash the password with bcrypt
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

// CheckPassword compares the given password with the stored hashed password
func CheckPassword(password, hashedPassword string) bool {
	// Compare the provided password with the stored hashed password
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
