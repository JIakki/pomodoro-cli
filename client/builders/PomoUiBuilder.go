package builders

import (
	"github.com/JIakki/pomodoro-cli/client/components"
	"github.com/JIakki/pomodoro-cli/client/components/TimeHumanize"
)

type PomoUiBuilder interface {
	NewPomoUi()
	BuildProgressBar()
	BuildTimer()
	SetTimeHumanizeProvider(humanize.TimeHumanize)
	GetResult() *components.PomoUi
}
