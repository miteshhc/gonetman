package components

import (
	"github.com/miteshhc/gonetman/app"
	"github.com/miteshhc/gonetman/consts"

	"github.com/Wifx/gonetworkmanager/v2"
	"github.com/rivo/tview"
)

/*
    TODO: Add stuff from AccessPoint
*/

// NewSettings Creates Settings submenu
func NewSettings() *tview.Form {
    settingsForm := tview.NewForm()
    Flex.AddItem(settingsForm, 0, 1, false)

    hostname, err := app.NMSettings.GetPropertyHostname()
    if err != nil {
        ErrorModal(err, MainMenu)
        return settingsForm
    }

    isNMEnabledProperty, err := app.NMInstance.GetPropertyNetworkingEnabled()
    if err != nil {
        ErrorModal(err, MainMenu)
        return settingsForm
    }

    isWirelessHWEnabledProperty, err := app.NMInstance.GetPropertyWirelessHardwareEnabled()
    if err != nil {
        ErrorModal(err, MainMenu)
        return settingsForm
    }

    isWirelessEnabledProperty, err := app.NMInstance.GetPropertyWirelessEnabled()
    if err != nil {
        ErrorModal(err, MainMenu)
        return settingsForm
    }

    isWWanHWEnabledProperty, err := app.NMInstance.GetPropertyWwanHardwareEnabled()
    if err != nil {
        ErrorModal(err, MainMenu)
        return settingsForm
    }

    isWWanEnabledProperty, err := app.NMInstance.GetPropertyWwanEnabled()
    if err != nil {
        ErrorModal(err, MainMenu)
        return settingsForm
    }

    isNMEnabled := 1
    if isNMEnabledProperty {
        isNMEnabled = 0
    }

    isWirelessHWEnabled := "Disabled"
    if isWirelessHWEnabledProperty {
        isWirelessHWEnabled = "Enabled"
    }

    isWirelessEnabled := 1
    if isWirelessEnabledProperty {
        isWirelessEnabled = 0
    }

    isWWanHWEnabled := "Disabled"
    if isWWanHWEnabledProperty {
        isWWanHWEnabled = "Enabled"
    }

    isWWanEnabled := "Disabled"
    if isWWanEnabledProperty {
        isWWanEnabled = "Enabled"
    }

    settingsForm.SetBorder(true).SetTitle("Settings")
    settingsForm.AddInputField("Hostname: ", hostname, 18, nil, nil)
    settingsForm.AddDropDown("Network Manager: ", []string{"Enable", "Disable"}, isNMEnabled, func(option string, optionIndex int) {
        if optionIndex == 0 {
            isNMEnabledProperty = true
        } else {
            isNMEnabledProperty = false
        }
    })
    settingsForm.AddTextView("Wireless HW: ", isWirelessHWEnabled, 30, 1, false, false)
    settingsForm.AddDropDown("Wireless: ", []string{"Enable", "Disable"}, isWirelessEnabled, func(option string, optionIndex int) {
        if optionIndex == 0 {
            isWirelessEnabledProperty = true
        } else {
            isWirelessEnabledProperty = false
        }
    })
    settingsForm.AddTextView("WWan HW: ", isWWanHWEnabled, 30, 1, false, false)
    settingsForm.AddTextView("WWan: ", isWWanEnabled, 30, 1, false, false)
    settingsForm.AddButton("<Back>", func() {
        app.App.SetFocus(MainMenu)
        Flex.RemoveItem(settingsForm)
    })

    settingsForm.AddButton("<Reload>", func() {
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
                    app.App.SetRoot(Flex, true).SetFocus(settingsForm)
                }
                app.App.SetRoot(Flex, true).SetFocus(settingsForm)
            })

        app.App.SetRoot(reloadModal, true).SetFocus(reloadModal)
    })


    settingsForm.AddButton("<OK>", func() {
        saveSettings(settingsForm, hostname, isNMEnabledProperty, isWirelessEnabledProperty)
    })

    return settingsForm
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

// saveSettings Saves all the changed settings
func saveSettings(settingsForm *tview.Form,
                originalHostname string,
                isNMEnabled bool,
                isWirelessEnabled bool,
                ) {
    newHostname := settingsForm.GetFormItem(0).(*tview.InputField).GetText()
    
    if newHostname != originalHostname {
        if err := app.NMSettings.SaveHostname(newHostname); err != nil {
            ErrorModal(err, MainMenu)
            return
        }
    }

    if wirelessEnabled, _ := app.NMInstance.GetPropertyWirelessEnabled(); wirelessEnabled != isWirelessEnabled {
        if err := app.NMInstance.SetPropertyWirelessEnabled(isWirelessEnabled); err != nil {
            ErrorModal(err, MainMenu)
            return
        }
    }

    if nmEnabled, _ := app.NMInstance.GetPropertyNetworkingEnabled(); nmEnabled != isNMEnabled {
        if err := app.NMInstance.Enable(isNMEnabled); err != nil {
            ErrorModal(err, MainMenu)
            return
        }
    }

    Flex.RemoveItem(settingsForm)
    app.App.SetFocus(MainMenu)
}
