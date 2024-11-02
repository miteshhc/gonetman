package components

import (
	"github.com/miteshhc/gonetman/app"
	"github.com/rivo/tview"
)

// NewActivateConnection Creates Activate Connection submenu
func NewActivateConnection() *tview.List {
    connectionsList := tview.NewList().ShowSecondaryText(false)

    connectionsList.AddItem(
        "",
        "",
        0,
        nil,
        ).
        AddItem(
        "Go Back",
        "",
        'B',
        func() {
            app.App.SetFocus(MainMenu)
            Flex.RemoveItem(connectionsList)
        },
        )

    connectionsList.SetBorder(true).SetTitle("Connections")
    Flex.AddItem(connectionsList, 0, 1, false)
    return connectionsList
}
