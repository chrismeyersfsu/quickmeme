package main

import (
	"fmt"
	"strings"

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
	tagGifEntryMap map[string]map[uint]*GifEntry
}

func (ge *GifEntry) loadPixelBuf() {
	p, err := gdk.PixbufAnimationNewFromFile(ge.gif.Path)
	panicIf(err)
	ge.pixbuf = p
}

func (gm *GifManager) printTagMap() {
	fmt.Println("Tag Map state:")
	for k, v := range gm.tagGifEntryMap {
		fmt.Print(k, " = ", v, " : ")
		var ids []uint
		for id := range v {
			ids = append(ids, id)
		}
		fmt.Println(strings.Join(strings.Fields(fmt.Sprint(ids)), ","))
	}
}

func (gm *GifManager) updateTagMap(ge *GifEntry, tagsBefore []string, tagsAfter []string) {
	fmt.Println("Before:")
	gm.printTagMap()
	for _, tag := range tagsBefore {
		if _, ok := gm.tagGifEntryMap[tag][ge.gif.ID]; ok {
			delete(gm.tagGifEntryMap[tag], ge.gif.ID)
		}
	}
	for _, tag := range tagsAfter {
		if v, ok := gm.tagGifEntryMap[tag]; ok {
			v[ge.gif.ID] = ge
		} else {
			gm.tagGifEntryMap[tag] = make(map[uint]*GifEntry)
			gm.tagGifEntryMap[tag][ge.gif.ID] = ge
		}
	}
	fmt.Println("After:")
	gm.printTagMap()
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

	gm.tagGifEntryMap = make(map[string]map[uint]*GifEntry)

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
			results = append(results, *ge)
		}
	}
	return results
}
