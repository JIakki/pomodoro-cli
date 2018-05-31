package builders

import (
	"github.com/JIakki/pomodoro-cli/client/components"
	"github.com/JIakki/pomodoro-cli/client/components/TimeHumanize"
	"github.com/JIakki/pomodoro-cli/controllers"
	"github.com/JIakki/pomodoro-cli/models"
)

type PomoUiBuilder interface {
	NewPomoUi()
	BuildProgressBar()
	BuildStatus()
	BuildTimer(int)
	SetTimeHumanizeProvider(humanize.TimeHumanize)
	SetPomoProvider(*models.Pomo)
	SetControllersProvider(*controllers.Controllers)
	GetResult() *components.PomoUi
}
