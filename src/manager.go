package main

import (
	"github.com/gotk3/gotk3/gdk"
)

type GifEntry struct {
	gif    *Gif
	pixbuf *gdk.PixbufAnimation
}

type GifManager struct {
	entries []GifEntry
	search  string
	root    string
	gifDb   *GifDb
}

func (ge *GifEntry) loadPixelBuf() {
	p, err := gdk.PixbufAnimationNewFromFile(ge.gif.Path)
	panicIf(err)
	ge.pixbuf = p
}

func (gm *GifManager) GetTags(ge *GifEntry) []string {
	return gm.gifDb.GetTags(ge.gif)
}

func (gm *GifManager) AddTags(ge *GifEntry, names []string) {
	gm.gifDb.AddTags(ge.gif, names)
}

func (gm *GifManager) init() {
	gm.gifDb = &GifDb{}
	gm.gifDb.init()

	gif_paths := GetGifPaths(gm.root)
	for _, gif_path := range gif_paths {
		gif, _ := gm.gifDb.GetOrCreate(gif_path)
		ge := GifEntry{gif, nil}
		ge.loadPixelBuf()

		gm.entries = append(gm.entries, ge)
	}
}

func (gm *GifManager) GetEntries() []GifEntry {
	return gm.entries
}
