package main

import (
	"github.com/gotk3/gotk3/gdk"
)

type GifEntry struct {
	gif    *Gif
	tags   []string
	pixbuf *gdk.PixbufAnimation
}

type GifManager struct {
	entries        []GifEntry
	search         string
	root           string
	gifDb          *GifDb
	tagGifEntryMap map[string]map[*GifEntry]GifEntry
}

func (ge *GifEntry) loadPixelBuf() {
	p, err := gdk.PixbufAnimationNewFromFile(ge.gif.Path)
	panicIf(err)
	ge.pixbuf = p
}

func (gm *GifManager) updateTagMap(ge *GifEntry, tagsBefore []string, tagsAfter []string) {
	for _, tag := range tagsBefore {
		if _, ok := gm.tagGifEntryMap[tag][ge]; ok {
			delete(gm.tagGifEntryMap[tag], ge)
		}
	}
}

func (gm *GifManager) GetTags(ge *GifEntry) []string {
	tagsBefore := ge.tags[:]
	ge.tags = gm.gifDb.GetTags(ge.gif)
	gm.updateTagMap(ge, tagsBefore, ge.tags)
	return ge.tags
}

func (gm *GifManager) SetTags(ge *GifEntry, names []string) {
	gm.updateTagMap(ge, ge.tags, names)
	gm.gifDb.SetTags(ge.gif, names)
	ge.tags = names
}

func (gm *GifManager) init() {
	gm.gifDb = &GifDb{}
	gm.gifDb.init()

	gif_paths := GetGifPaths(gm.root)
	for _, gif_path := range gif_paths {
		gif, _ := gm.gifDb.GetOrCreate(gif_path)
		ge := GifEntry{gif: gif}
		ge.loadPixelBuf()

		gm.entries = append(gm.entries, ge)
	}
}

func (gm *GifManager) GetEntries(search string) []GifEntry {
	results := []GifEntry{}
	if search == "" {
		return gm.entries
	}
	if val, ok := gm.tagGifEntryMap[search]; ok {
		for _, ge := range val {
			results = append(results, ge)
		}
	}
	return results
}
