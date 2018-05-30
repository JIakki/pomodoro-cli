package models

import (
	"time"
)

type Pomo struct {
}

func (p *Pomo) Start(c chan int) {
	for range time.Tick(1000 * time.Millisecond) {
		c <- (time.Now().Minute() * 60) + time.Now().Second()
	}
}

func NewPomo() *Pomo {
	return &Pomo{}
}
