package components

import (
	"github.com/miteshhc/gonetman/app"
	"github.com/miteshhc/gonetman/consts"
	"github.com/miteshhc/gonetman/helpers"

	"github.com/Wifx/gonetworkmanager/v2"
	"github.com/rivo/tview"
)

/*
    TODO: Add stuff from AccessPoint
*/

// NewSettings creates Settings submenu
func NewSettings() *tview.Form {
    hostname, err := app.NMSettings.GetPropertyHostname()
    if err != nil {
        panic("Failed to get hostname: " + err.Error())
    }

    form := tview.NewForm()
    Flex.AddItem(form, 0, 1, false)

    form.SetBorder(true).SetTitle("Settings")
    form.AddInputField("Hostname: ", hostname, 40, nil, nil)

    form.AddButton("<Back>", func() {
        app.App.SetFocus(MainMenu)
        Flex.RemoveItem(form)
    })

    form.AddButton("<Reload>", func() {
        reloadModal := tview.NewModal()

        reloadModal.SetText("What do you want to reload?").
            AddButtons([]string{"Nothing",
                                "Everything",
                                "NetworkManager.conf",
                                "DNS config",
                                "DNS plugin"}).
            SetDoneFunc(func(buttonIndex int, buttonLabel string) {
                switch buttonIndex {
                case 1:
                    reload(app.NMInstance, reloadModal, consts.ReloadEverything)
                case 2:
                    reload(app.NMInstance, reloadModal, consts.ReloadNetworkManager)
                case 3:
                    reload(app.NMInstance, reloadModal, consts.ReloadDNSConfig)
                case 4:
                    reload(app.NMInstance, reloadModal, consts.ReloadDNSPlugin)
                default:
                    app.App.SetRoot(Flex, true).SetFocus(form)
                }
                app.App.SetRoot(Flex, true).SetFocus(form)
            })

        app.App.SetRoot(reloadModal, true).SetFocus(reloadModal)
    })

    form.AddButton("<OK>", func() {
        newHostname := form.GetFormItem(0).(*tview.InputField).GetText()
        if err := app.NMSettings.SaveHostname(newHostname); err != nil {
            helpers.ErrorModal(err, MainMenu, app.App)
        } else {
            app.App.SetFocus(MainMenu)
        }
    })

    return form
}

// reload Reloads settings as per provided flag
func reload(nmInstance gonetworkmanager.NetworkManager, reloadModal *tview.Modal, flag uint32) {
    err := nmInstance.Reload(flag)

    if err != nil {
        helpers.ErrorModal(err, reloadModal, app.App)
    } else {
        app.App.SetFocus(reloadModal)
    }
}
