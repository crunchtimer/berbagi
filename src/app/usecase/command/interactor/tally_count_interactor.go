package interactor

import (
	"github.com/hourglasshoro/berbagi/src/app/domain/entity"
	"github.com/hourglasshoro/berbagi/src/app/domain/repository"
	"github.com/hourglasshoro/berbagi/src/app/domain/value_object"
	query_service "github.com/hourglasshoro/berbagi/src/app/usecase/query"
	"github.com/hourglasshoro/berbagi/src/app/usecase/query/read_model"
	"github.com/labstack/gommon/log"
	"sort"
	"sync"
	"time"
)

const AccessTokenTTL = time.Hour

type TallyCountInteractor struct {
	twitterRepo     repository.TwitterAccessTokenRepository
	getTwitterToken query_service.GetTwitterAccessTokenQueryService
	twitterAuth     query_service.AuthorizeTwitterQueryService
	fetchTweets     query_service.FetchTweetsQueryService
	fetchUser       query_service.FetchUserQueryService
	slackRepo       repository.SlackRepository
}

func NewTallyCountInteractor(
	twitterRepo repository.TwitterAccessTokenRepository,
	getTwitterToken query_service.GetTwitterAccessTokenQueryService,
	twitterAuth query_service.AuthorizeTwitterQueryService,
	fetchTweets query_service.FetchTweetsQueryService,
	fetchUser query_service.FetchUserQueryService,
	slackRepo repository.SlackRepository,
) *TallyCountInteractor {
	intr := new(TallyCountInteractor)
	intr.twitterRepo = twitterRepo
	intr.getTwitterToken = getTwitterToken
	intr.twitterAuth = twitterAuth
	intr.fetchTweets = fetchTweets
	intr.fetchUser = fetchUser
	intr.slackRepo = slackRepo
	return intr
}

func (intr *TallyCountInteractor) Invoke() (err error) {

	// Get access token for Twitter
	twitterToken, err := intr.getTwitterToken.GetTwitterToken()
	log.Print("Get access token for Twitter")

	if err == query_service.NoTwitterAccessTokenExistException {
		newToken, e := intr.twitterAuth.AuthTwitter()
		log.Print("Authorize Twitter")

		if e != nil {
			err = e
			return
		}
		err = intr.twitterRepo.SetTwitterToken(newToken, AccessTokenTTL)
		log.Print("Set Twitter access token")

		twitterToken, err = intr.getTwitterToken.GetTwitterToken()
		log.Print("Get access token for Twitter")
	}
	if err != nil {
		return
	}

	// Fetch tweets
	aggregates, err := intr.fetchTweets.FetchTweets(twitterToken)
	log.Print("Fetch tweets")

	if err != nil {
		return
	}

	wg := &sync.WaitGroup{}
	mutex := &sync.Mutex{}

	var results []value_object.TallyCountItem
	for _, agg := range aggregates.Aggregates {
		wg.Add(1)
		go func(aggregate read_model.AggregateReadModel) {
			defer wg.Done()
			author, err := intr.fetchUser.FetchUser(aggregate.AuthorId, twitterToken)
			if err != nil {
				return
			}
			authorValue, err := entity.NewAuthor(author.DisplayName, author.UserName, author.ImageUrl)
			if err != nil {
				return
			}
			item, err := value_object.NewTallyCountItem(*authorValue, aggregate.LikeCount)
			if err != nil {
				return
			}
			mutex.Lock()
			results = append(results, *item)
			mutex.Unlock()

		}(agg)
	}
	wg.Wait()

	sort.Slice(results, func(i, j int) bool { return results[i].LikeCount > results[j].LikeCount })

	log.Print(results)
	if len(results) > 0 {
		err = intr.slackRepo.Send(results)
	}
	return
}
