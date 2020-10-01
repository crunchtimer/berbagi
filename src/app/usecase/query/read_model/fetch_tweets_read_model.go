package read_model

type FetchTweetsReadModel struct {
	Aggregates []AggregateReadModel
}

type AggregateReadModel struct {
	AuthorId  uint64
	LikeCount uint64
}
