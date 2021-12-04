package main

import (
	"strings"

	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/gtk"
)

type Application struct {
	gm         GifManager
	win        *gtk.Window
	resultList *gtk.ListBox
	searchbar  *gtk.Entry
}

func (app *Application) handleSearch() {
	searchText, err := app.searchbar.GetText()
	panicIf(err)

	searchText = strings.TrimSpace(searchText)

	app.resultList.ShowAll()
}

func (app *Application) addSearchResultItem(pixbuf *gdk.PixbufAnimation) {
	row, err := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 10)
	panicIf(err)
	icon, err := gtk.ImageNew()
	panicIf(err)
	icon.Show()

	icon.SetFromAnimation(pixbuf)

	row.Add(icon)
	app.resultList.Add(row)
	app.resultList.ShowAll()
	//app.win.ShowAll()
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
	entry, err := gtk.EntryNew()
	panicIf(err)

	// resultList
	list, err := gtk.ListBoxNew()
	panicIf(err)
	list.SetSelectionMode(gtk.SELECTION_SINGLE)

	// window(layout(searchbar, resultList))
	layoutList.Add(entry)
	scroll.Add(list)
	layoutList.Add(scroll)
	win.Add(layoutList)

	gm := GifManager{nil, "", "/home/meyers/Downloads/gifs/", nil}
	gm.init()

	app := &Application{
		gm,
		win,
		list,
		entry,
	}

	// Add all the gifs
	for _, entry := range gm.GetEntries() {
		app.addSearchResultItem(entry.pixbuf)
	}

	win.SetDefaultSize(1024, 1024)
	win.SetPosition(gtk.WIN_POS_CENTER_ALWAYS)

	entry.Connect("changed", func() {
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
