package components

import (
	"github.com/miteshhc/gonetman/app"

	"github.com/rivo/tview"
)

var MainMenu = tview.NewList()

// NewMainMenu creates main menu and draws it to the screen
func NewMainMenu() *tview.List {
    MainMenu.AddItem(
        "Activate Connection",
        "Activate/Deactivate one of the available connections",
        'a',
        nil,
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
