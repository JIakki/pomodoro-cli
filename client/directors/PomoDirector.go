package directors

import (
	"github.com/JIakki/pomodoro-cli/client/builders"
	"github.com/JIakki/pomodoro-cli/client/components/TimeHumanize"
)

type PomoDirector struct {
	builder builders.PomoUiBuilder
}

func (pd *PomoDirector) Construct() {
	pd.builder.NewPomoUi()
	pd.builder.BuildProgressBar()
	pd.builder.SetTimeHumanizeProvider(humanize.NewTimeHumanize3D())
	pd.builder.BuildTimer()
}

func NewPomoDirector(builder builders.PomoUiBuilder) *PomoDirector {
	return &PomoDirector{builder}
}
