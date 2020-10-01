package controller

import "github.com/hourglasshoro/berbagi/src/app/usecase/command/interactor"

type AppController struct {
	tallyCountInteractor interactor.TallyCountInteractor
}

func NewAppController(tallyCountInteractor interactor.TallyCountInteractor) *AppController {
	ctrl := new(AppController)
	ctrl.tallyCountInteractor = tallyCountInteractor
	return ctrl
}

func (ctrl *AppController) TallyCount() (err error) {
	err = ctrl.tallyCountInteractor.Invoke()
	return
}
