package builders

import (
	"github.com/JIakki/pomodoro-cli/client/components"
)

type PomoBuilder interface {
	NewPomo()
	BuildProgressBar()
	GetResult() *components.Pomo
}
