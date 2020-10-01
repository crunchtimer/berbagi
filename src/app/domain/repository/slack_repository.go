package repository

import "github.com/hourglasshoro/berbagi/src/app/domain/value_object"

type SlackRepository interface {
	Send(tallyCount []value_object.TallyCountItem) (err error)
}
