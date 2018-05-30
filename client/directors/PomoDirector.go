package directors

import (
	"github.com/JIakki/pomodoro-cli/client/builders"
	"github.com/JIakki/pomodoro-cli/client/components/TimeHumanize"
	"github.com/JIakki/pomodoro-cli/models"
)

type PomoDirector struct {
	builder builders.PomoUiBuilder
}

func (pd *PomoDirector) Construct() {
	pd.builder.NewPomoUi()

	pd.builder.SetTimeHumanizeProvider(humanize.NewTimeHumanize3D())
	pd.builder.SetPomoProvider(models.NewPomo(4, 30, 5))

	pd.builder.BuildProgressBar()
	pd.builder.BuildStatus()
	pd.builder.BuildTimer()
}

func NewPomoDirector(builder builders.PomoUiBuilder) *PomoDirector {
	return &PomoDirector{builder}
}
