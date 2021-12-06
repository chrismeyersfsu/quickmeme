package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/gotk3/gotk3/gtk"
	"golang.design/x/clipboard"
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

	for {
		if row := app.resultList.GetRowAtIndex(0); row != nil {
			app.resultList.Remove(row)
		} else {
			break
		}
	}

	for _, entry := range app.gm.GetEntries(searchText) {
		app.addSearchResultItem(entry)
	}

	app.resultList.ShowAll()
}

func (app *Application) handleTagTextChange(tagText *gtk.Entry, ge GifEntry) func() {
	return func() {
		tagsStr, err := tagText.GetText()
		panicIf(err)

		tagsStr = strings.TrimSpace(tagsStr)

		endsWithNewlineFlag := false
		if len(tagsStr) > 0 && tagsStr[len(tagsStr)-1] == ',' {
			endsWithNewlineFlag = true
		}
		tags := strings.Split(tagsStr, ",")
		if !endsWithNewlineFlag {
			tags = tags[:len(tags)-1]
		}
		var tagsFiltered []string
		for _, tag := range tags {
			if tag != "" {
				tag = strings.TrimSpace(tag)
				tagsFiltered = append(tagsFiltered, tag)
			}
		}

		app.gm.SetTags(&ge, tagsFiltered)
	}
}

func (app *Application) handleRowSelected() func() {
	return func() {
		i := app.resultList.GetSelectedRow().GetIndex()
		fmt.Println("Clicked on the row! ", i)
		if i == -1 {
			return
		}
		url := UploadGifFile(app.gm.entries[i].gif.Path)
		fmt.Println("URL ", string(url))
		clipboard.Write(clipboard.FmtText, url)
	}
}

func (app *Application) addSearchResultItem(ge GifEntry) {
	row, err := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 10)
	panicIf(err)

	icon, err := gtk.ImageNew()
	panicIf(err)
	icon.SetFromAnimation(ge.pixbuf)

	tagText, err := gtk.EntryNew()
	panicIf(err)

	tags := app.gm.GetTags(&ge)
	if len(tags) > 0 {
		tagText.SetText(strings.Join(tags, ","))
	}
	tagText.Connect("changed", app.handleTagTextChange(tagText, ge))

	row.Add(icon)
	row.Add(tagText)
	app.resultList.Add(row)
}

func NewApplication() *Application {

	gtk.Init(nil)

	// window
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

	// resultList
	list, err := gtk.ListBoxNew()
	panicIf(err)
	list.SetSelectionMode(gtk.SELECTION_SINGLE)

	// window(layout(searchbar, resultList))
	layoutList.Add(searchbar)
	scroll.Add(list)
	layoutList.Add(scroll)
	win.Add(layoutList)

	gifroot := os.Getenv("GIFROOT")
	if gifroot == "" {
		gifroot = "/home/meyers/Downloads/gifs/"
	}

	gm := GifManager{search: "", root: gifroot}
	gm.init()

	app := &Application{
		gm,
		win,
		list,
		searchbar,
	}

	layoutList.Connect("row-activated", app.handleRowSelected())

	// Add all the gifs
	for _, entry := range gm.GetEntries("") {
		app.addSearchResultItem(entry)
	}
	app.resultList.ShowAll()

	win.SetDefaultSize(1024, 1024)
	win.SetPosition(gtk.WIN_POS_CENTER_ALWAYS)

	/*
		win.Connect("clicked", func(window *gtk.Window, event *gdk.Event) {
			i := app.resultList.GetSelectedRow().GetIndex()

			app.searchbar.GrabFocusWithoutSelecting()
		})
	*/

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
