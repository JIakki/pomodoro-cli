package directors

import (
	"github.com/JIakki/pomodoro-cli/client/builders"
)

type PomoDirector struct {
	builder builders.PomoBuilder
}

func (pd *PomoDirector) Construct() {
	pd.builder.NewPomo()
	pd.builder.BuildProgressBar()
}

func NewPomoDirector(builder builders.PomoBuilder) *PomoDirector {
	return &PomoDirector{builder}
}
