package models

import (
	"time"
)

const WorkType = "work"
const RestType = "rest"
const PlayStatus = "play"
const PauseStatus = "pause"

type Pomo struct {
	ticker         *time.Ticker
	pomodoros      int
	donePomodoros  int
	workDuration   int
	restDuration   int
	durationRemain int
	pomoType       string // work, rest
	status         string
}

func (p *Pomo) Start(c chan int) {
	p.ticker = time.NewTicker(1000 * time.Millisecond)
	for range p.ticker.C {
		p.durationRemain -= 1

		c <- p.durationRemain

		if p.durationRemain <= 0 {
			p.Done()
		}

	}
}

func (p *Pomo) Done() {
	if p.pomoType == WorkType {
		p.durationRemain = p.restDuration
		p.pomoType = RestType
	} else {
		p.durationRemain = p.workDuration
		p.pomoType = WorkType

	}

}

func (p *Pomo) Stop() {
	p.ticker.Stop()
}

func (p *Pomo) GetType() string {
	return p.pomoType
}

func NewPomo(pomodoros int, work int, rest int) *Pomo {
	return &Pomo{
		pomodoros:      pomodoros,
		workDuration:   work * 60,
		restDuration:   rest * 60,
		durationRemain: work * 60,
		pomoType:       WorkType,
		status:         PlayStatus,
	}
}
