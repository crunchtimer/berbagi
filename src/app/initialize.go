package main

import (
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

	twitterAuth := twitter.NewAuthorizeTwitterQueryService()
	fetchTweets := twitter.NewFetchTweetsQueryService()
	fetchUser := twitter.NewFetchUserQueryService()
	slackRepo := slack_webhook.NewSlackRepository()

	searchIntr := *interactor.NewTallyCountInteractor(
		twitterAuth,
		fetchTweets,
		fetchUser,
		slackRepo,
	)

	appCtrl := controller.NewAppController(searchIntr)
	ctrls.App = *appCtrl
	return ctrls, nil
}
