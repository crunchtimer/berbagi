package main

import (
	"github.com/hourglasshoro/berbagi/src/app/infrastructure/redis"
	slack_webhook "github.com/hourglasshoro/berbagi/src/app/infrastructure/web/slack"
	"github.com/hourglasshoro/berbagi/src/app/infrastructure/web/twitter"
	"github.com/hourglasshoro/berbagi/src/app/presentation/controller"
	"github.com/hourglasshoro/berbagi/src/app/usecase/command/interactor"
)

type Controllers struct {
	App controller.AppController
}

func NewControllers() (ctrls *Controllers, err error) {
	ctrls = new(Controllers)
	redisInst, err := redis.NewRedis()
	if err != nil {
		return
	}

	twitterRepo := redis.NewTwitterAccessTokenRepository(redisInst)
	getTwitterToken := redis.NewGetTwitterAccessTokenQueryService(redisInst)
	twitterAuth := twitter.NewAuthorizeTwitterQueryService()
	fetchTweets := twitter.NewFetchTweetsQueryService()
	fetchUser := twitter.NewFetchUserQueryService()
	slackRepo := slack_webhook.NewSlackRepository()

	searchIntr := *interactor.NewTallyCountInteractor(
		twitterRepo,
		getTwitterToken,
		twitterAuth,
		fetchTweets,
		fetchUser,
		slackRepo,
	)

	appCtrl := controller.NewAppController(searchIntr)
	ctrls.App = *appCtrl
	return ctrls, nil
}
