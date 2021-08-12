package userinterface

import (
	"log"

	ui "github.com/gizak/termui/v3"
)

func Userinterface(clientEvents, serverEvents chan string, quit chan int) {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()
	input_p := inputWidgetInit()
	printEvents := ui.PollEvents()
	for {
		select {
		case e := <-printEvents:
			//
			handleClientEvent(input_p, e.ID, quit)
		case msg := <-serverEvents:
			handleServerEvent(msg)
		case <-quit:
			return
		}
	}
}
