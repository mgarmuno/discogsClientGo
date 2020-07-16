package gui

import (
	"fmt"
	"log"

	"github.com/cloudfoundry-attic/jibber_jabber"
	"github.com/gotk3/gotk3/gtk"
	"github.com/mgarmuno/discogsClientGo/db"
	"github.com/mgarmuno/discogsClientGo/file"
)

const (
	defaultLang = "en"
	locales     = "locales"
)

var language string

// InitApp initializes the app GUI
func InitApp() {
	fmt.Println("Initializing GUI application...")
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

	if !file.CheckToken() {
		showAddTokenWindow()
	}

	win.Add(l)

	win.SetDefaultSize(500, 120)

	win.ShowAll()

	gtk.Main()
}

func getAppLang() string {
	dbLang := db.CheckLanguage()
	if dbLang != "" {
		return dbLang
	}
	lang, err := jibber_jabber.DetectLanguage()
	if err != nil {
		log.Println("Can't determinate system language, setting default language")
		return defaultLang
	}
	return lang
}

func showAddTokenWindow() {
	win, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	if err != nil {
		log.Fatal("Error creating token window:", err)
	}

	masterBox, err := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 1)

	buttonsBox, err := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 1)

	win.SetTitle("Add Token")
	win.Connect("destroy", func() {
		win.Destroy()
	})

	addBut, err := gtk.ButtonNewWithLabel("Add Token")
	if err != nil {
		log.Fatal("Error creating button for token:", err)
	}
	canBut, err := gtk.ButtonNewWithLabel("Cancel")
	if err != nil {
		log.Fatal("Error creating button for token:", err)
	}
	addTokenMes, err := gtk.LabelNew("You need to add a token from your Discogs account in order to operate from this application, do you want to add it no?")

	canBut.Connect("clicked", func() {
		win.Destroy()
	})

	masterBox.PackStart(addTokenMes, true, true, 1)
	buttonsBox.PackStart(addBut, true, true, 1)
	buttonsBox.PackStart(canBut, true, true, 1)
	masterBox.PackStart(buttonsBox, true, true, 1)

	win.Add(masterBox)

	win.SetDefaultSize(500, 300)
	win.SetPosition(gtk.WIN_POS_CENTER)

	win.ShowAll()
}
