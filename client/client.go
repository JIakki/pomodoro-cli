package client

import (
	"github.com/JIakki/pomodoro-cli/client/builders"
	"github.com/JIakki/pomodoro-cli/client/directors"
)

func NewClient() {
	c := make(chan int)
	classicPomoBuilder := builders.NewPomoTuiBuilder()
	director := directors.NewPomoDirector(classicPomoBuilder)
	director.Construct()

	pomoUi := classicPomoBuilder.GetResult()
	go func(c chan int) {
		for {
			pomoUi.UpdateTimer(<-c)
			pomoUi.UpdateStatus()
			pomoUi.UpdateUi()
		}
	}(c)

	pomoUi.Start(c)
}
