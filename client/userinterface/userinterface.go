package userinterface

import (
	"log"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	tsize "github.com/kopoli/go-terminal-size"
)

type look struct {
	input_p  *widgets.Paragraph
	output_p *widgets.Paragraph
}

func Userinterface(clientEvents, serverEvents chan string, quit chan int) {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()
	s, err := tsize.GetSize()
	if err != nil {
		s.Width = 60
		s.Height = 25
	}
	l := widgetInit(s.Width, s.Height)
	printEvents := ui.PollEvents()
	for {
		select {
		case e := <-printEvents:
			handleClientEvent(l, e.ID, clientEvents, quit)
		case msg := <-serverEvents:
			handleServerEvent(l, msg)
		case <-quit:
			return
		}
	}
}

func widgetInit(width, height int) *look {
	return &look{inputWidgetInit(width, height),
		outputWidgetInit(width, height)}
}
