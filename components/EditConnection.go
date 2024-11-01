package components

import (
	"github.com/Wifx/gonetworkmanager/v2"
	"github.com/miteshhc/gonetman/app"
	"github.com/miteshhc/gonetman/helpers"

	// "github.com/Wifx/gonetworkmanager/v2"
	"github.com/rivo/tview"
)

// NewEditConnection creates Edit Connection submenu
func NewEditConnection() *tview.List {
    connectionsList := tview.NewList().ShowSecondaryText(false)

    savedConnections, err := app.NMSettings.ListConnections()

    if err != nil {
        panic(err)
    }

    for i, connection := range savedConnections {
        connectionSettings, err := connection.GetSettings()

        if err != nil {
            panic(err)
        }

        connectionID, ok := connectionSettings["connection"]["id"].(string)

        if !ok {
            connectionID = "Unknown"
        }

        connectionsList.AddItem(
            connectionID,
            "",
            helpers.ConvertToChar(i),
            func() {
                connectionSubMenu(connectionSettings, connectionsList)
            })
    }

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

// connectionSubMenu create submenu of given connection dynamically
func connectionSubMenu(connectionSettings gonetworkmanager.ConnectionSettings, connectionsList *tview.List) {
    connectionForm := tview.NewForm()
    Flex.AddItem(connectionForm, 0, 3, false)

    timestamp, _ := connectionSettings["connection"]["timestamp"].(uint64)
    lastConnected := helpers.GetLocalTime(timestamp)

    id, ok := connectionSettings["connection"]["id"].(string)
    if !ok {
        id = "Unknown"
    }

    connectionForm.
        AddTextView("ID: ", id, 30, 1, false, false).
        AddTextView("Last Connected: ", lastConnected, 30, 1, false, false).
        AddButton("<Back>", func() {
            app.App.SetFocus(connectionsList)
            Flex.RemoveItem(connectionForm)
        })

    connectionForm.SetBorder(true).SetTitle(id)

    app.App.SetFocus(connectionForm)
}
