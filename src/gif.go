package main

import (
	"github.com/gotk3/gotk3/gdk"
)

type GifEntry struct {
	path   string
	pixbuf *gdk.PixbufAnimation
}

func (gif *GifEntry) Start() *gdk.PixbufAnimation {
	if gif.pixbuf != nil {
		return gif.pixbuf
	}
	gif.load()
	return gif.pixbuf
}

func (gif *GifEntry) load() {
	p, err := gdk.PixbufAnimationNewFromFile(gif.path)
	panicIf(err)
	gif.pixbuf = p
}
