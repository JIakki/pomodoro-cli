package controllers

import (
	"github.com/JIakki/pomodoro-cli/models"
)

type Controllers struct {
	pomoInstance *models.Pomo
}

func NewControllers(instance *models.Pomo) *Controllers {
	return &Controllers{instance}
}

func (ctrl *Controllers) GetWorkDuration() int {
	return ctrl.pomoInstance.GetWorkDuration()
}

func (ctrl *Controllers) IsStarted() bool {
	return ctrl.pomoInstance.IsStarted()
}
