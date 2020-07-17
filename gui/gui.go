package gui

import (
	"fmt"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
	"github.com/mgarmuno/discogsClientGo/api"
	"github.com/mgarmuno/discogsClientGo/db"
	"github.com/mgarmuno/discogsClientGo/file"
)

var (
	addTokenLabel  *widget.Label
	addTokenButton *widget.Button
)

// InitApp initializes the app GUI
func InitApp() {
	showMainWindow()
}

func showAddTokenWindow() {
	w := fyne.CurrentApp().NewWindow("Add Token")
	entry := widget.NewEntry()
	entry.SetPlaceHolder("Enter Token here...")
	container := fyne.NewContainerWithLayout(
		layout.NewMaxLayout(),
		entry)
	w.SetContent(widget.NewVBox(
		container,
		widget.NewButton("Save Token", func() {
			if !file.SaveToken(entry.Text) {
				showErrorSavingToken()
			} else {
				addTokenButton.SetText("Change token")
				addTokenLabel.Hide()
			}
			w.Close()
		}),
		widget.NewButton("Close", func() {
			w.Close()
		}),
	))
	w.Resize(fyne.NewSize(600, 300))
	w.Show()
}

func showMainWindow() {
	app := app.New()
	w := app.NewWindow("Hello")
	box := widget.NewVBox()
	userName := getUserName()
	box.Append(widget.NewLabel(fmt.Sprintf("Hello %s!", userName)))
	addTokenLabel = widget.NewLabel("No Token present, it's necessary for operate with Discogs, please add one with the following button!")
	addTokenLabel.Hide()
	addTokenButton = widget.NewButton("Change Token", showAddTokenWindow)
	box.Append(addTokenButton)
	if !file.CheckToken() {
		addTokenButton.SetText("Add Token")
		addTokenLabel.Show()
		box.Append(addTokenLabel)
	}
	box.Append(widget.NewButton("User Sync", func() {
		if file.CheckToken() {
			api.UpdateUserInfo()
		}
	}))
	box.Append(widget.NewButton("Quit", func() {
		app.Quit()
	}))
	w.SetContent(box)
	w.Resize(fyne.NewSize(600, 300))
	w.ShowAndRun()
}

func getUserName() string {
	user := db.GetUserInfo()
	if user.UserName == "" {
		return "User"
	}
	return user.UserName
}

func showErrorSavingToken() {
	w := fyne.CurrentApp().NewWindow("Error saving token")
	w.SetContent(widget.NewVBox(
		widget.NewLabel("There was an error trying to save the token."),
		widget.NewButton("Close", func() {
			w.Close()
		}),
	))
	w.Show()
}
