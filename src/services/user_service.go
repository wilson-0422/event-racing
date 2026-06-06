package services

import (
	"event-racing/src/config"
	"event-racing/src/models"

	"golang.org/x/crypto/bcrypt"
)

func CreateUser(username, password, role string) (*models.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	result, err := config.DB.Exec(
		"INSERT INTO users (username, password, role) VALUES (?, ?, ?)",
		username, string(hashedPassword), role,
	)
	if err != nil {
		return nil, err
	}
	id, _ := result.LastInsertId()
	return &models.User{ID: id, Username: username, Role: role}, nil
}

func GetUserByUsername(username string) (*models.User, error) {
	var u models.User
	err := config.DB.QueryRow(
		"SELECT id, username, password, role, created_at FROM users WHERE username = ?",
		username,
	).Scan(&u.ID, &u.Username, &u.Password, &u.Role, &u.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func GetUserByID(id int64) (*models.User, error) {
	var u models.User
	err := config.DB.QueryRow(
		"SELECT id, username, password, role, created_at FROM users WHERE id = ?",
		id,
	).Scan(&u.ID, &u.Username, &u.Password, &u.Role, &u.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func VerifyPassword(username, password string) (*models.User, error) {
	user, err := GetUserByUsername(username)
	if err != nil {
		return nil, err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, err
	}
	return user, nil
}
