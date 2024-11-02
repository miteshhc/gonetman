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

// NewSettings Creates Settings submenu
func NewSettings() *tview.Form {
    settingsForm := tview.NewForm()
    Flex.AddItem(settingsForm, 0, 1, false)

    hostname, err := app.NMSettings.GetPropertyHostname()

    if err != nil {
        helpers.ErrorModal(err, MainMenu, app.App)
        return settingsForm
    }

    isWirelessEnabledProperty, err := app.NMInstance.GetPropertyWirelessEnabled()

    if err != nil {
        helpers.ErrorModal(err, MainMenu, app.App)
        return settingsForm
    }

    isWirelessHWEnabledProperty, err := app.NMInstance.GetPropertyWirelessHardwareEnabled()

    if err != nil {
        helpers.ErrorModal(err, MainMenu, app.App)
        return settingsForm
    }

    isWirelessHWEnabled := "Disabled"
    if isWirelessHWEnabledProperty {
        isWirelessHWEnabled = "Enabled"
    }

    isWirelessEnabled := int8(1)
    if isWirelessEnabledProperty {
        isWirelessEnabled = 0
    }

    settingsForm.SetBorder(true).SetTitle("Settings")
    settingsForm.AddInputField("Hostname: ", hostname, 18, nil, nil)
    settingsForm.AddTextView("Wireless HW: ", isWirelessHWEnabled, 30, 1, false, false)
    settingsForm.AddDropDown("Wireless: ", []string{"Enable", "Disable"}, int(isWirelessEnabled), func(option string, optionIndex int) {
        if optionIndex == 0 {
            isWirelessEnabledProperty = true
        } else {
            isWirelessEnabledProperty = false
        }
    })

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
        saveSettings(settingsForm, hostname, isWirelessEnabledProperty)
    })

    return settingsForm
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

// saveSettings Saves all the changed settings
func saveSettings(settingsForm *tview.Form,
                originalHostname string,
                isWirelessEnabled bool,
                ) {
    newHostname := settingsForm.GetFormItem(0).(*tview.InputField).GetText()
    
    if newHostname != originalHostname {
        if err := app.NMSettings.SaveHostname(newHostname); err != nil {
            helpers.ErrorModal(err, MainMenu, app.App)
            return
        }
    }

    if err := app.NMInstance.SetPropertyWirelessEnabled(isWirelessEnabled); err != nil {
        helpers.ErrorModal(err, MainMenu, app.App)
        return
    }

    // Return to main menu on success
    Flex.RemoveItem(settingsForm)
    app.App.SetFocus(MainMenu)
}
