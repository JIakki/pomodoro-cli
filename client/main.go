package client

import (
	"github.com/JIakki/pomodoro-cli/client/builders"
	"github.com/JIakki/pomodoro-cli/client/directors"
)

func NewClient() {
	classicPomoBuilder := builders.NewClassicPomoBuilder()
	director := directors.NewPomoDirector(classicPomoBuilder)
	director.Construct()

	pomo := classicPomoBuilder.GetResult()
	pomo.Start()
}
