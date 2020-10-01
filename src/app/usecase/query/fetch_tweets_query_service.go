package query_service

import (
	"errors"
	"github.com/hourglasshoro/berbagi/src/app/domain/value_object"
	"github.com/hourglasshoro/berbagi/src/app/usecase/query/read_model"
)

var CannotFetchTweetsException = errors.New("cannot fetch tweets")

type FetchTweetsQueryService interface {
	FetchTweets(token value_object.AccessToken) (readModel read_model.FetchTweetsReadModel, err error)
}
