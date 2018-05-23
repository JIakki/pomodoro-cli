package builders

import (
	"github.com/JIakki/pomodoro-cli/client/components"
)

type ClassicPomoBuilder struct {
	pomo *components.Pomo
}

func (pb *ClassicPomoBuilder) NewPomo() {
	pb.pomo = &components.Pomo{}
}

func (pb *ClassicPomoBuilder) BuildProgressBar() {
	pb.pomo.SetProgressBar("1/4")
}

func (pb *ClassicPomoBuilder) GetResult() *components.Pomo {
	return pb.pomo
}

func NewClassicPomoBuilder() *ClassicPomoBuilder {
	return &ClassicPomoBuilder{}

}
