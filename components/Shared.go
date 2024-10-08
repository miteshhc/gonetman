package components

import (
    "github.com/miteshhc/gonetman/app"

	"github.com/rivo/tview"
	"github.com/Wifx/gonetworkmanager/v2"
)

var NMInstance gonetworkmanager.NetworkManager
var NMSettings gonetworkmanager.Settings

func init() {
    var err error
    NMInstance, err = gonetworkmanager.NewNetworkManager()

    if err != nil {
        ErrorModal(err, MainMenu)
    }

    NMSettings, err = gonetworkmanager.NewSettings()

    if err != nil {
        ErrorModal(err, MainMenu)
    }
}

func ErrorModal(err error, root tview.Primitive) {
    errorModal := tview.NewModal().
    SetText("Error: " + err.Error()).
    AddButtons([]string{"Got it"}).
    SetDoneFunc(func(buttonIndex int, buttonLabel string) {
        if buttonIndex == 0 {
            app.App.SetRoot(root, true).SetFocus(root)
        }
    })

    app.App.SetRoot(errorModal, true).SetFocus(errorModal)
}
