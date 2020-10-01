package value_object

import (
	"github.com/hourglasshoro/berbagi/src/app/domain/entity"
)

type TallyCountItem struct {
	Author    entity.Author
	LikeCount uint64
}

func NewTallyCountItem(author entity.Author, likeCount uint64) (item *TallyCountItem, err error) {
	item = new(TallyCountItem)
	item.Author = author
	item.LikeCount = likeCount
	return item, nil
}
