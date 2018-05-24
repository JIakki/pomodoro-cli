package directors

import (
	"github.com/JIakki/pomodoro-cli/client/builders"
)

type PomoDirector struct {
	builder builders.PomoUiBuilder
}

func (pd *PomoDirector) Construct() {
	pd.builder.NewPomoUi()
	pd.builder.BuildProgressBar()
	pd.builder.BuildTimer()
}

func NewPomoDirector(builder builders.PomoUiBuilder) *PomoDirector {
	return &PomoDirector{builder}
}
