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

func inputWidgetInit() *widgets.Paragraph {
	p := widgets.NewParagraph()
	p.Text = prettyPrint("")
	p.SetRect(0, 0, 50, 3)
	ui.Render(p_my)
	return p
}

func handleClientEvent(input_p *widgets.Paragraph, e string, quit chan int) {
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
		ui.Render(p_my)
	case e == "<Backspace>":
		if len(text) == 0 {
			continue
		}
		text = trimLastChar(text)
		input_p.Text = prettyPrint(text)
		ui.Render(p_my)
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
