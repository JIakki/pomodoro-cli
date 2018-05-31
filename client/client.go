package client

import (
	"github.com/JIakki/pomodoro-cli/client/builders"
	"github.com/JIakki/pomodoro-cli/client/directors"
)

func NewClient() {
	classicPomoBuilder := builders.NewPomoTuiBuilder()
	director := directors.NewPomoDirector(classicPomoBuilder)
	director.Construct()

	pomoUi := classicPomoBuilder.GetResult()
	pomoUi.Render()
}
