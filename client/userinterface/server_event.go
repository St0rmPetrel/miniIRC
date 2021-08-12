package userinterface

import (
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

func handleServerEvent(l *look, msg string) {
	output_p := l.output_p
	output_p.Text = msg + output_p.Text
	if len(output_p.Text) > 2048 {
		output_p.Text = output_p.Text[:1024]
	}
	ui.Render(output_p)
}

func outputWidgetInit(width, height int) *widgets.Paragraph {
	p := widgets.NewParagraph()
	p.Text = ""
	p.SetRect(0, 3, width, height)
	ui.Render(p)
	return p
}
