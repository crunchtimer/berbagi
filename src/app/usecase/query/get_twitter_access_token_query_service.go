package query_service

import (
	"errors"
	"github.com/hourglasshoro/berbagi/src/app/domain/value_object"
)

var NoTwitterAccessTokenExistException = errors.New("no twitter access token exists")

type GetTwitterAccessTokenQueryService interface {
	GetTwitterToken() (token value_object.AccessToken, err error)
}
