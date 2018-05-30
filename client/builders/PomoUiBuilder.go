package builders

import (
	"github.com/JIakki/pomodoro-cli/client/components"
	"github.com/JIakki/pomodoro-cli/client/components/TimeHumanize"
	"github.com/JIakki/pomodoro-cli/models"
)

type PomoUiBuilder interface {
	NewPomoUi()
	BuildProgressBar()
	BuildStatus()
	BuildTimer()
	SetTimeHumanizeProvider(humanize.TimeHumanize)
	SetPomoProvider(*models.Pomo)
	GetResult() *components.PomoUi
}
