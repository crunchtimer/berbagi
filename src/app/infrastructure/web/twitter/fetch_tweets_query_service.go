package twitter

import (
	"encoding/json"
	"fmt"
	"github.com/hourglasshoro/berbagi/src/app/domain/value_object"
	query_service "github.com/hourglasshoro/berbagi/src/app/usecase/query"
	"github.com/hourglasshoro/berbagi/src/app/usecase/query/read_model"
	"net/http"
	"net/url"
	"strconv"
)

const MaxTweetsResultNumber = 100
const Hashtag = "#CrunchtimerShare"

type FetchTweetsQueryService struct {
}

func NewFetchTweetsQueryService() *FetchTweetsQueryService {
	qs := new(FetchTweetsQueryService)
	return qs
}

type TweetsResponse struct {
	Data []struct {
		Id            string `json:"id"`
		Text          string `json:"text"`
		AuthorId      string `json:"author_id"`
		PublicMetrics struct {
			RetweetCount int `json:"retweet_count"`
			ReplyCount   int `json:"reply_count"`
			LikeCount    int `json:"like_count"`
			QuoteCount   int `json:"quote_count"`
		} `json:"public_metrics"`
	} `json:"data"`
	Meta struct {
		NewestId    string `json:"newest_id"`
		OldestId    string `json:"oldest_id"`
		ResultCount int    `json:"result_count"`
		NextToken   string `json:"next_token"`
	} `json:"meta"`
}

func (qs *FetchTweetsQueryService) FetchTweets(token value_object.AccessToken) (readModel read_model.FetchTweetsReadModel, err error) {
	u := fmt.Sprintf("https://api.twitter.com/2/tweets/search/recent?max_results=%d&query=%s&tweet.fields=author_id,public_metrics", MaxTweetsResultNumber, url.QueryEscape(Hashtag))
	req, err := http.NewRequest("GET", u, nil)
	if err != nil {
		return
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token.Token))
	client := new(http.Client)
	res, err := client.Do(req)
	if err != nil {
		return
	}
	if res.StatusCode != http.StatusOK {
		err = query_service.CannotFetchTweetsException
		return
	}
	defer res.Body.Close()

	var tweetsRes TweetsResponse
	err = json.NewDecoder(res.Body).Decode(&tweetsRes)
	if err != nil {
		return
	}
	var aggregates []read_model.AggregateReadModel
	var aggregateMap = map[int]int{}
	for _, v := range tweetsRes.Data {
		authorId, e := strconv.Atoi(v.AuthorId)
		if e != nil {
			err = e
			return
		}
		aggregateMap[authorId] += v.PublicMetrics.LikeCount
	}

	for key, value := range aggregateMap {
		aggregates = append(aggregates, read_model.AggregateReadModel{
			AuthorId:  uint64(key),
			LikeCount: uint64(value),
		})
	}
	readModel = read_model.FetchTweetsReadModel{Aggregates: aggregates}
	return
}
