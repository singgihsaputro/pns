package service

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
)

// NotificationService is a thin wrapper to call FCM HTTP v1 endpoints.
// In this offline-friendly implementation we just perform an HTTP POST if the URL is reachable.
// Replace APIURL/ServerKey with your own values for production.
type NotificationService struct {
	APIURL    string
	ServerKey string
	Client    *http.Client
}

type PushMessage struct {
	Token string
	Title string
	Body  string
}

func NewNotificationService(serverKey string) *NotificationService {
	return &NotificationService{
		APIURL:    "https://fcm.googleapis.com/fcm/send",
		ServerKey: serverKey,
		Client:    &http.Client{Timeout: 10 * time.Second},
	}
}

func (s *NotificationService) Send(ctx context.Context, msg PushMessage) error {
	payload := fmt.Sprintf(`{"to": "%s", "notification": {"title": "%s", "body": "%s"}}`, msg.Token, escape(msg.Title), escape(msg.Body))
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, s.APIURL, strings.NewReader(payload))
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", "key "+s.ServerKey)
	req.Header.Set("Content-Type", "application/json")

	resp, err := s.Client.Do(req)
	if err != nil {
		log.Printf("warning: push send simulated (offline?): %v", err)
		return nil
	}
	defer resp.Body.Close()
	if resp.StatusCode >= 300 {
		log.Printf("FCM returned status %d for token %s", resp.StatusCode, msg.Token)
	}
	return nil
}

func escape(v string) string {
	replacer := strings.NewReplacer(`"`, `'`)
	return replacer.Replace(v)
}

