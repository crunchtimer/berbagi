package query_service

import (
	"github.com/hourglasshoro/berbagi/src/app/domain/value_object"
	"github.com/hourglasshoro/berbagi/src/app/usecase/query/read_model"
)

type FetchUserQueryService interface {
	FetchUser(userId uint64, token value_object.AccessToken) (readModel read_model.FetchUserReadModel, err error)
}
