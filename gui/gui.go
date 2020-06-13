package gui

import (
	"fmt"
	"log"

	"github.com/gotk3/gotk3/gtk"
)

// InitApp initializes the app GUI
func InitApp() {
	fmt.Println("Iniziation GUI application...")
	gtk.Init(nil)

	win, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	if err != nil {
		log.Fatal("Unable to create new window:", err)
	}

	win.SetTitle("MyDiscogs Application")
	win.Connect("destroy", func() {
		gtk.MainQuit()
	})

	l, err := gtk.LabelNew("Hello, Discogs!")
	if err != nil {
		log.Fatal("Unable to create label:", err)
	}

	win.Add(l)

	win.SetDefaultSize(2000, 1200)

	win.ShowAll()

	gtk.Main()
}
