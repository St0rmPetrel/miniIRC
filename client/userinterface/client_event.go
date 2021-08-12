package userinterface

import (
	"unicode"
	"unicode/utf8"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

var (
	text string
)

func inputWidgetInit(width, height int) *widgets.Paragraph {
	p := widgets.NewParagraph()
	p.Text = prettyPrint("")
	p.SetRect(0, 0, width, 3)
	ui.Render(p)
	return p
}

func handleClientEvent(l *look, e string,
	clientEvents chan string, quit chan int) {
	input_p := l.input_p
	switch {
	case e == "<C-c>":
		quit <- 1
		return
	case isToPrint(e):
		text += e
		input_p.Text = prettyPrint(text)
		ui.Render(input_p)
	case e == "<Space>":
		text += " "
		input_p.Text = prettyPrint(text)
		ui.Render(input_p)
	case e == "<Backspace>":
		if len(text) == 0 {
			break
		}
		text = trimLastChar(text)
		input_p.Text = prettyPrint(text)
		ui.Render(input_p)
	case e == "<Enter>":
		if len(text) == 0 {
			break
		}
		clientEvents <- text
		text = ""
		input_p.Text = prettyPrint(text)
		ui.Render(input_p)
	}
}

func isToPrint(s string) bool {
	var (
		i int
		r rune
	)
	for i, r = range s {
		if i > 0 {
			return false
		}
	}
	if unicode.IsPrint(r) {
		return true
	}
	return false
}

func trimLastChar(s string) string {
	r, size := utf8.DecodeLastRuneInString(s)
	if r == utf8.RuneError && (size == 0 || size == 1) {
		size = 0
	}
	return s[:len(s)-size]
}

func prettyPrint(s string) string {
	return "-> " + s + "|"
}
