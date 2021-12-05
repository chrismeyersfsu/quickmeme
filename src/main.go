package main

import (
	"fmt"
	"strings"

	"github.com/gotk3/gotk3/gtk"
)

type Application struct {
	gm         GifManager
	win        *gtk.Window
	resultList *gtk.ListBox
	searchbar  *gtk.Entry
	tagbar     *gtk.Entry
}

func (app *Application) handleSearch() {
	searchText, err := app.searchbar.GetText()
	panicIf(err)

	searchText = strings.TrimSpace(searchText)

	app.resultList.ShowAll()
}

func (app *Application) handleTagTextChange(tagText *gtk.Entry, ge GifEntry) func() {
	return func() {
		tagsStr, err := tagText.GetText()
		panicIf(err)

		tagsStr = strings.TrimSpace(tagsStr)

		endsWithNewlineFlag := false
		if tagsStr[len(tagsStr)-1] == ',' {
			endsWithNewlineFlag = true
		}
		tags := strings.Split(tagsStr, ",")
		if !endsWithNewlineFlag {
			tags = tags[:len(tags)-1]
		}
		app.gm.AddTags(&ge, tags)
	}
}

func (app *Application) addSearchResultItem(ge GifEntry) {
	row, err := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 10)
	panicIf(err)

	icon, err := gtk.ImageNew()
	panicIf(err)
	icon.SetFromAnimation(ge.pixbuf)

	tagText, err := gtk.EntryNew()
	panicIf(err)

	tags := app.gm.GetTags(&ge)
	fmt.Println("The tags are: ", strings.Join(tags, ","))

	tagText.SetText(strings.Join(tags, ","))

	tagText.Connect("changed", app.handleTagTextChange(tagText, ge))

	row.Add(tagText)
	row.Add(icon)
	app.resultList.Add(row)
}

func NewApplication() *Application {

	gtk.Init(nil)

	win, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	panicIf(err)

	win.SetTitle("Simple Example")
	win.Connect("destroy", func() {
		gtk.MainQuit()
	})

	// layout
	layoutList, err := gtk.ListBoxNew()
	panicIf(err)
	layoutList.SetSelectionMode(gtk.SELECTION_NONE)

	// scroll
	hadj, err := gtk.AdjustmentNew(0, 0, 640, 50, 640, 640)
	panicIf(err)
	vadj, err := gtk.AdjustmentNew(0, 0, 640, 50, 640, 640)
	panicIf(err)
	scroll, err := gtk.ScrolledWindowNew(hadj, vadj)
	scroll.SetPolicy(gtk.POLICY_NEVER, gtk.POLICY_AUTOMATIC)
	scroll.SetSizeRequest(-1, 1024)
	panicIf(err)

	// searchbar
	searchbar, err := gtk.EntryNew()
	panicIf(err)

	// tagbar
	tagbar, err := gtk.EntryNew()
	panicIf(err)

	// resultList
	list, err := gtk.ListBoxNew()
	panicIf(err)
	list.SetSelectionMode(gtk.SELECTION_SINGLE)

	// window(layout(searchbar, resultList))
	layoutList.Add(searchbar)
	layoutList.Add(tagbar)
	scroll.Add(list)
	layoutList.Add(scroll)
	win.Add(layoutList)

	gm := GifManager{nil, "", "/home/meyers/Downloads/gifs/", nil}
	gm.init()

	app := &Application{
		gm,
		win,
		list,
		searchbar,
		tagbar,
	}

	// Add all the gifs
	for _, entry := range gm.GetEntries() {
		app.addSearchResultItem(entry)
	}
	app.resultList.ShowAll()

	win.SetDefaultSize(1024, 1024)
	win.SetPosition(gtk.WIN_POS_CENTER_ALWAYS)

	searchbar.Connect("changed", func() {
		app.handleSearch()
	})

	return app
}

func (app *Application) Main() {
	app.win.ShowAll()
	gtk.Main()
}

func main() {
	app := NewApplication()
	app.Main()

}
