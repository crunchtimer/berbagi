package query_service

import (
	"errors"
	"github.com/hourglasshoro/berbagi/src/app/domain/value_object"
)

var CannotGetTwitterAccessTokenException = errors.New("cannot get twitter access token")

type AuthorizeTwitterQueryService interface {
	AuthTwitter() (token value_object.AccessToken, err error)
}
