package directors

import (
	"github.com/JIakki/pomodoro-cli/client/builders"
	"github.com/JIakki/pomodoro-cli/client/components/TimeHumanize"
	"github.com/JIakki/pomodoro-cli/controllers"
	"github.com/JIakki/pomodoro-cli/models"
)

type PomoDirector struct {
	builder builders.PomoUiBuilder
}

func (pd *PomoDirector) Construct() {
	pomo := models.NewPomo()
	ctrls := controllers.NewControllers(pomo)
	pd.builder.NewPomoUi()

	pd.builder.SetTimeHumanizeProvider(humanize.NewTimeHumanize3D())
	pd.builder.SetPomoProvider(pomo)
	pd.builder.SetControllersProvider(ctrls)

	pd.builder.BuildProgressBar()
	pd.builder.BuildStatus()
	pd.builder.BuildTimer(ctrls.GetWorkDuration())
}

func NewPomoDirector(builder builders.PomoUiBuilder) *PomoDirector {
	return &PomoDirector{builder}
}
