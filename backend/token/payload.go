package token

import (
	"errors"

	"github.com/google/uuid"

	"time"
)

var (
	ErrInvalidToken = errors.New("token is invalid")
	ErrExpiredToken = errors.New("token has expired")
)

// Payload is a struct that represents the token payload
type Payload struct {
	ID       uuid.UUID `json:"id"`
	Username string    `json:"username"`
	IssuedAt time.Time `json:"issued_at"`
	Expired  time.Time `json:"expired_at"`
}

// NewPayload creates a new payload for a token
func NewPayload(username string, duration time.Duration) (*Payload, error) {
	tokenID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	payload := &Payload{
		ID:       tokenID,
		Username: username,
		IssuedAt: time.Now(),
		Expired:  time.Now().Add(duration),
	}

	return payload, nil

}

// Valid checks if the token payload is valid or not
func (payload Payload) Valid() error {
	if time.Now().After(payload.Expired) {
		return ErrExpiredToken
	}
	return nil
}
