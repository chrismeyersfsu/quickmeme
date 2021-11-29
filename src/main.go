package main

import (
	"github.com/gotk3/gotk3/gtk"
)

type Application struct {
	win        *gtk.Window
	resultList *gtk.ListBox
	searchbar  *gtk.Entry
}

func (app *Application) addSearchResultByPath(path string) {
	gif := &GifEntry{path, nil}
	app.addSearchResultItem(gif)
}

func (app *Application) addSearchResultItem(item *GifEntry) {
	row, err := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 10)
	panicIf(err)
	icon, err := gtk.ImageNew()
	panicIf(err)
	icon.Show()

	icon.SetFromAnimation(item.Start())

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

	// Image

	app := &Application{
		win,
		list,
		entry,
	}

	app.addSearchResultByPath("/home/meyers/Downloads/gifs/1.gif")
	app.addSearchResultByPath("/home/meyers/Downloads/gifs/2.gif")
	app.addSearchResultByPath("/home/meyers/Downloads/gifs/3.gif")
	app.addSearchResultByPath("/home/meyers/Downloads/gifs/4.gif")

	win.SetDefaultSize(1024, 1024)
	win.SetPosition(gtk.WIN_POS_CENTER_ALWAYS)

	return app
}

func (app *Application) Main() {
	app.win.ShowAll()
	gtk.Main()

	app.addSearchResultByPath("/home/meyers/Downloads/gifs/5.gif")
	app.addSearchResultByPath("/home/meyers/Downloads/gifs/6.gif")
	app.addSearchResultByPath("/home/meyers/Downloads/gifs/7.gif")
	app.addSearchResultByPath("/home/meyers/Downloads/gifs/8.gif")
	app.addSearchResultByPath("/home/meyers/Downloads/gifs/9.gif")

}

func main() {
	app := NewApplication()
	app.Main()

}
