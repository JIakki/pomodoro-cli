package components

import (
	"github.com/JIakki/pomodoro-cli/client/components/TimeHumanize"
	"github.com/marcusolsson/tui-go"
)

type PomoUi struct {
	progressBar          *tui.StatusBar
	timer                *tui.Label
	timeHumanizeProvider humanize.TimeHumanize
}

func (p *PomoUi) SetTimeHumanizeProvider(provider humanize.TimeHumanize) {
	p.timeHumanizeProvider = provider
}

func (p *PomoUi) SetProgressBar(status string) {
	p.progressBar = tui.NewStatusBar(status)
}

func (p *PomoUi) SetTimer(time int) {
	humanizedTime := p.timeHumanizeProvider.Convert(time)
	p.timer = tui.NewLabel(humanizedTime)
	p.timer.SetStyleName("fatal")
	p.timer.SizeHint()
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
