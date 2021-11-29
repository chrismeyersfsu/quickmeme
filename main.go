package main

import (
	"context"
	"os"
	"os/signal"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"

	xwidget "fyne.io/x/fyne/widget"
	"github.com/gotk3/gotk3/gtk"
)

func CancelAll(cancel []context.CancelFunc) {
	for _, c := range cancel {
		c()
	}
}

func _ConsumeSignals(cancel []context.CancelFunc, sigchan <-chan os.Signal) {
	// TODO: Should we loop over the signals or just break after the first?
	// Maybe we should have an if condition for SIGINT
	for range sigchan {
		CancelAll(cancel)
	}
	gtk.MainQuit()
	os.Exit(0)
}

func StartSignalHandlers(cancel []context.CancelFunc) {
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, os.Interrupt)

	go _ConsumeSignals(cancel, sigchan)
}

type GUI struct {
	search *widget.Entry
	grid   *fyne.Container
}

func createImage(path string) *xwidget.AnimatedGif {
	img, err := xwidget.NewAnimatedGif(storage.NewFileURI(path))
	if err != nil {
		println("Error creating animated gif")
	}
	img.Start()
	return img
}

func main() {
	gui := GUI{}
	//rootCtx := context.Background()
	//clipCtx, clipCancel := context.WithCancel(rootCtx)

	//StartSignalHandlers([]context.CancelFunc{clipCancel})

	// 				clipboard.Write(clipboard.FmtText, t)

	a := app.New()
	w := a.NewWindow("Hello")

	//hello := widget.NewLabel("Hello Fyne!")
	gui.search = widget.NewEntry()
	gui.search.SetText("")

	var images []*xwidget.AnimatedGif
	images = append(images, createImage("/home/meyers/Downloads/gifs/1.gif"))
	images = append(images, createImage("/home/meyers/Downloads/gifs/2.gif"))

	gui.grid = container.NewGridWithColumns(4)
	gui.grid.Add(images[0])
	gui.grid.Add(images[1])
	gui.grid.Add(createImage("/home/meyers/Downloads/gifs/1.gif"))
	gui.grid.Add(createImage("/home/meyers/Downloads/gifs/1.gif"))
	gui.grid.Add(createImage("/home/meyers/Downloads/gifs/1.gif"))
	gui.grid.Add(createImage("/home/meyers/Downloads/gifs/1.gif"))
	gui.grid.Add(createImage("/home/meyers/Downloads/gifs/1.gif"))
	gui.grid.Add(createImage("/home/meyers/Downloads/gifs/1.gif"))
	gui.grid.Add(createImage("/home/meyers/Downloads/gifs/1.gif"))
	gui.grid.Add(createImage("/home/meyers/Downloads/gifs/1.gif"))
	gui.grid.Add(createImage("/home/meyers/Downloads/gifs/1.gif"))
	gui.grid.Add(createImage("/home/meyers/Downloads/gifs/1.gif"))
	gui.grid.Add(createImage("/home/meyers/Downloads/gifs/1.gif"))
	gui.grid.Add(createImage("/home/meyers/Downloads/gifs/1.gif"))
	gui.grid.Add(createImage("/home/meyers/Downloads/gifs/1.gif"))
	gui.grid.Add(createImage("/home/meyers/Downloads/gifs/1.gif"))

	w.SetContent(container.NewVScroll(container.New(layout.NewBorderLayout(gui.search, nil, nil, nil), gui.search, gui.grid)))
	w.Resize(fyne.NewSize(1200, 600))
	w.Canvas().Focus(gui.search)
	w.ShowAndRun()
}

// {"hello": "world"}
