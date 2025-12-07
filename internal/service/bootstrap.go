package service

import (
	"context"
	"log"

	"github.com/example/pns/internal/repository"
)

// Bootstrap ensures a default admin user exists for first-time access.
func Bootstrap(ctx context.Context, users *repository.UserRepository, auth *AuthService, email, password string) {
	if email == "" || password == "" {
		return
	}
	_, err := users.FindByEmail(ctx, email)
	if err == nil {
		return
	}
	hash, err := auth.HashPassword(password)
	if err != nil {
		log.Printf("cannot hash bootstrap password: %v", err)
		return
	}
	if _, err := users.Create(ctx, email, hash); err != nil {
		log.Printf("cannot create bootstrap user: %v", err)
	}
}

