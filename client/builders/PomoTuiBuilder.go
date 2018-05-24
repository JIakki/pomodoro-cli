package builders

import (
	"github.com/JIakki/pomodoro-cli/client/components"
)

type PomoTuiBuilder struct {
	pomoUi *components.PomoUi
}

func (pb *PomoTuiBuilder) NewPomoUi() {
	pb.pomoUi = &components.PomoUi{}
}

func (pb *PomoTuiBuilder) BuildProgressBar() {
	pb.pomoUi.SetProgressBar("1/4")
}

func (pb *PomoTuiBuilder) BuildTimer() {
	pb.pomoUi.SetTimer("30:00")
}

func (pb *PomoTuiBuilder) GetResult() *components.PomoUi {
	return pb.pomoUi
}

func NewPomoTuiBuilder() *PomoTuiBuilder {
	return &PomoTuiBuilder{}

}
