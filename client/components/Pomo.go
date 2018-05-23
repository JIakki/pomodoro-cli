package components

import (
	"github.com/marcusolsson/tui-go"
)

type Pomo struct {
	progressBar *tui.StatusBar
}

func (p *Pomo) SetProgressBar(status string) {
	p.progressBar = tui.NewStatusBar(status)
}

func (p *Pomo) Start() {
	box := tui.NewVBox(
		p.progressBar,
	)

	ui, err := tui.New(box)
	if err != nil {
		panic(err)
	}
	ui.SetKeybinding("Esc", func() { ui.Quit() })

	if err := ui.Run(); err != nil {
		panic(err)
	}
}
