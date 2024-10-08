package components

import (
	"github.com/miteshhc/gonetman/app"
	"github.com/miteshhc/gonetman/consts"

	"github.com/rivo/tview"
	"github.com/Wifx/gonetworkmanager/v2"
)

// NewSettings creates and draws settings submenu
func NewSettings() *tview.Form {
    hostname, err := NMSettings.GetPropertyHostname()
    if err != nil {
        panic("Failed to get hostname: " + err.Error())
    }

    form := tview.NewForm()
    Flex.AddItem(form, 0, 1, false)

    form.SetBorder(true).SetTitle("Settings")
    form.AddInputField("Hostname: ", hostname, 40, nil, nil)

    /*
    // Move this block to Edit Connections, as this property shows whether the
    // connections can be modified or not

    var isModifiable string

    if canModify, err := nmSettings.GetPropertyCanModify(); err != nil {
        ErrorModal(err, form)
    } else {
        if canModify {
            isModifiable = "Yes"
        } else {
            isModifiable = "No"
        }
    }

    form.AddTextView("Modifiable: ", isModifiable, 10, 10, false, false)
    */

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
                    reload(NMInstance, reloadModal, consts.ReloadEverything)
                case 2:
                    reload(NMInstance, reloadModal, consts.ReloadNetworkManager)
                case 3:
                    reload(NMInstance, reloadModal, consts.ReloadDNSConfig)
                case 4:
                    reload(NMInstance, reloadModal, consts.ReloadDNSPlugin)
                default:
                    app.App.SetRoot(Flex, true).SetFocus(form)
                }
                app.App.SetRoot(Flex, true).SetFocus(form)
            })

        app.App.SetRoot(reloadModal, true).SetFocus(reloadModal)
    })

    form.AddButton("<OK>", func() {
        newHostname := form.GetFormItem(0).(*tview.InputField).GetText()
        if err := NMSettings.SaveHostname(newHostname); err != nil {
            ErrorModal(err, MainMenu)
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
        ErrorModal(err, reloadModal)
    } else {
        app.App.SetFocus(reloadModal)
    }
}
