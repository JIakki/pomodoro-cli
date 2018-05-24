package components

import (
	"github.com/marcusolsson/tui-go"
)

type PomoUi struct {
	progressBar *tui.StatusBar
	timer       *tui.Label
}

func (p *PomoUi) SetProgressBar(status string) {

	p.progressBar = tui.NewStatusBar(status)
}

func (p *PomoUi) SetTimer(time string) {
	p.timer = tui.NewLabel(time)
	p.timer.SetStyleName("fatal")
}

func (p *PomoUi) Start() {
	theme := tui.NewTheme()
	theme.SetStyle("label.fatal", tui.Style{Bg: tui.ColorDefault, Fg: tui.ColorRed})

	hBox := tui.NewHBox(
		tui.NewSpacer(),
		p.timer,
		tui.NewSpacer(),
	)
	vbox := tui.NewVBox(
		p.progressBar,
		tui.NewSpacer(),
		hBox,
		tui.NewSpacer(),
	)

	ui, err := tui.New(vbox)
	if err != nil {
		panic(err)
	}
	ui.SetKeybinding("Esc", func() { ui.Quit() })
	ui.SetTheme(theme)

	if err := ui.Run(); err != nil {
		panic(err)
	}
}
