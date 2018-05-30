package client

import (
	"github.com/JIakki/pomodoro-cli/client/builders"
	"github.com/JIakki/pomodoro-cli/client/directors"
	"github.com/JIakki/pomodoro-cli/models"
)

func NewClient() {
	c := make(chan int)
	pomo := models.NewPomo(4, 30, 5)
	classicPomoBuilder := builders.NewPomoTuiBuilder()
	director := directors.NewPomoDirector(classicPomoBuilder)
	director.Construct()

	pomoUi := classicPomoBuilder.GetResult()
	go pomo.Start(c)
	go func(c chan int) {
		for {
			pomoUi.UpdatePomoType(pomo.GetType())
			pomoUi.UpdateTimer(<-c)
			pomoUi.UpdateUi()
		}
	}(c)

	pomoUi.Start()
}
