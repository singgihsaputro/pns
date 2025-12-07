package model

import "time"

type User struct {
	ID           string    `json:"id"`
	Email        string    `json:"email"`
	PasswordHash string    `json:"-"`
	CreatedAt    time.Time `json:"created_at"`
}

type DeviceToken struct {
	ID        int64     `json:"id"`
	DeviceID  string    `json:"device_id"`
	FCMToken  string    `json:"fcm_device_token"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}

type Promotion struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	ImageURL    string    `json:"image_url"`
	Description string    `json:"description"`
	StartDate   time.Time `json:"start_date"`
	EndDate     time.Time `json:"end_date"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
}

