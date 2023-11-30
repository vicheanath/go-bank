package token

import (
	"time"
)

// Token is a struct that represents an authentication token
type Maker interface {
	CreateToken(username string, duration time.Duration) (string, error)
	VerifyToken(token string) (*Payload, error)
}