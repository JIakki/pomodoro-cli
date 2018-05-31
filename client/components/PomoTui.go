package components

import (
	"github.com/JIakki/pomodoro-cli/client/components/TimeHumanize"
	"github.com/JIakki/pomodoro-cli/controllers"
	"github.com/JIakki/pomodoro-cli/models"
	"github.com/marcusolsson/tui-go"
)

type PomoUi struct {
	ui                   tui.UI
	progressBar          *tui.StatusBar
	status               *tui.Label
	timer                *tui.Label
	timeHumanizeProvider humanize.TimeHumanize
	pomoProvider         *models.Pomo
	controllersProvider  *controllers.Controllers
}

func (p *PomoUi) SetTimeHumanizeProvider(provider humanize.TimeHumanize) {
	p.timeHumanizeProvider = provider
}

func (p *PomoUi) SetPomoProvider(provider *models.Pomo) {
	p.pomoProvider = provider
}

func (p *PomoUi) SetControllersProvider(provider *controllers.Controllers) {
	p.controllersProvider = provider
}

func (p *PomoUi) SetProgressBar(status string) {
	p.progressBar = tui.NewStatusBar(status)
}

func (p *PomoUi) SetStatus() {
	p.status = tui.NewLabel(p.pomoProvider.GetStatus())
	p.status.SetStyleName("info")
}

func (p *PomoUi) UpdateStatus() {
	p.status.SetText(p.pomoProvider.GetStatus())

}

func (p *PomoUi) SetTimer(time int) {
	humanizedTime := p.timeHumanizeProvider.Convert(time)
	p.timer = tui.NewLabel(humanizedTime)
	p.timer.SetStyleName("fatal")
}

func (p *PomoUi) UpdateTimer(time int) {
	humanizedTime := p.timeHumanizeProvider.Convert(time)

	if p.pomoProvider.GetStatus() == models.WorkStatus {
		p.timer.SetStyleName("fatal")
	} else {
		p.timer.SetStyleName("success")
	}

	p.timer.SetText(humanizedTime)
}

func (p *PomoUi) ToggleTimer() {
	if p.controllersProvider.IsStarted() == true {
		p.controllersProvider.Stop()
		return
	}

	c := p.controllersProvider.Start()
	go func(c chan int) {
		for {
			p.UpdateTimer(<-c)
			p.UpdateStatus()
			p.UpdateUi()
		}
	}(c)

}

func (p *PomoUi) Render() {
	theme := tui.NewTheme()
	theme.SetStyle("label.fatal", tui.Style{Bg: tui.ColorDefault, Fg: tui.ColorRed})
	theme.SetStyle("label.info", tui.Style{Bg: tui.ColorDefault, Fg: tui.ColorYellow})
	theme.SetStyle("label.success", tui.Style{Bg: tui.ColorDefault, Fg: tui.ColorGreen})

	timer := tui.NewHBox(
		tui.NewSpacer(),
		p.timer,
		tui.NewSpacer(),
	)

	vBox := tui.NewVBox(
		p.progressBar,
		p.status,
		tui.NewSpacer(),
		timer,
		tui.NewSpacer(),
	)

	ui, err := tui.New(vBox)
	p.ui = ui

	if err != nil {
		panic(err)
	}
	ui.SetKeybinding("Esc", func() { ui.Quit() })
	ui.SetKeybinding("p", func() { p.ToggleTimer() })
	ui.SetTheme(theme)

	if err := ui.Run(); err != nil {
		panic(err)
	}
}

func (p *PomoUi) UpdateUi() {
	p.ui.Update(func() {})
}
