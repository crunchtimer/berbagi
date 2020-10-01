package repository

import (
	"github.com/hourglasshoro/berbagi/src/app/domain/value_object"
	"time"
)

type TwitterAccessTokenRepository interface {
	SetTwitterToken(token value_object.AccessToken, ttl time.Duration) (err error)
}
