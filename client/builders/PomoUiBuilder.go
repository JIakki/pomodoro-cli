package builders

import (
	"github.com/JIakki/pomodoro-cli/client/components"
)

type PomoUiBuilder interface {
	NewPomoUi()
	BuildProgressBar()
	BuildTimer()
	GetResult() *components.PomoUi
}
