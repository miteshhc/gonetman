package components

import (
	"github.com/Wifx/gonetworkmanager/v2"
	"github.com/miteshhc/gonetman/app"

	"github.com/rivo/tview"
)

var MainMenu = tview.NewList()

func init() {
    var err error
    app.NMInstance, err = gonetworkmanager.NewNetworkManager()

    if err != nil {
        ErrorModal(err, MainMenu)
    }

    app.NMSettings, err = gonetworkmanager.NewSettings()

    if err != nil {
        ErrorModal(err, MainMenu)
    }
}

// NewMainMenu Creates main menu and draws it to the screen
func NewMainMenu() *tview.List {
    MainMenu.AddItem(
        "Activate Connection",
        "Activate/Deactivate one of the available connections",
        'a',
        func() {
            activateConnection := NewActivateConnection()
            app.App.SetFocus(activateConnection)
        },
    ).AddItem(
        "Edit Connection(s)",
        "Manage saved connections",
        'e',
        func() {
            editConnection := NewEditConnection()
            app.App.SetFocus(editConnection)
        },
    ).AddItem(
        "Settings",
        "Change settings",
        's',
        func() {
            setting := NewSettings()
            app.App.SetFocus(setting)
        },
    ).AddItem(
        "Quit",
        "Quit gonetman",
        'q',
        func() {
            app.App.Stop()
        },
    )

    MainMenu.SetBorder(true).SetTitle("Main Menu")

    return MainMenu
}
