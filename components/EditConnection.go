package components

import (
	"github.com/miteshhc/gonetman/app"

	// "github.com/Wifx/gonetworkmanager/v2"
	"github.com/rivo/tview"
)

func NewEditConnection() *tview.Form {
    form := tview.NewForm()
    Flex.AddItem(form, 0, 1, false)

    form.SetBorder(true).SetTitle("Settings")

    form.AddButton("<Back>", func() {
        app.App.SetFocus(MainMenu)
        Flex.RemoveItem(form)
    })

    return form
}

