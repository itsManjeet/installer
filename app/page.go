package app

import "github.com/gotk3/gotk3/gtk"

type Page struct {
	*gtk.Box

	Title    *gtk.Label
	SubTitle *gtk.Label
	Icon     *gtk.Image

	Window     *gtk.Assistant
	GlobalData interface{}
}
