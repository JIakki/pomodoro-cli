package models

import (
	"time"
)

const WorkStatus = "work"
const RestStatus = "rest"

var instance *Pomo

type Pomo struct {
	ticker         *time.Ticker
	pomodoros      int
	donePomodoros  int
	workDuration   int
	restDuration   int
	durationRemain int
	status         string // work, rest
	started        bool
}

func (p *Pomo) Start(c chan int) {
	p.started = true

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
	if p.status == WorkStatus {
		p.durationRemain = p.restDuration
		p.status = RestStatus
	} else {
		p.durationRemain = p.workDuration
		p.status = WorkStatus
	}

	p.Stop()

}

func (p *Pomo) Stop() {
	p.ticker.Stop()
	p.started = false
}

func (p *Pomo) GetStatus() string {
	return p.status
}

func (p *Pomo) GetWorkDuration() int {
	return p.workDuration
}

func (p *Pomo) IsStarted() bool {
	return p.started
}

// Singleton
func NewPomo() *Pomo {
	println(instance == nil)
	if instance != nil {
		return instance
	}

	instance := &Pomo{
		pomodoros:      4,
		workDuration:   30 * 60,
		restDuration:   5 * 60,
		durationRemain: 30 * 60,
		status:         WorkStatus,
		started:        false,
	}

	return instance
}
