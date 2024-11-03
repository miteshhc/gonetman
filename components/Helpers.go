package components

import (
	"time"

	"github.com/miteshhc/gonetman/app"
	"github.com/rivo/tview"
)

// ErrorModal Creates new modal with Got it button
func ErrorModal(err error, root tview.Primitive) {
    errorModal := tview.NewModal()

    errorModal.SetText("Error: " + err.Error()).
                AddButtons([]string{"Got it"}).
                SetDoneFunc(func(buttonIndex int, buttonLabel string) {
                        app.App.SetFocus(root)
                        Flex.RemoveItem(errorModal)
                        Flex.AddItem(root, 0, 1, true)
                })

    Flex.Clear()
    Flex.AddItem(errorModal, 0, 1, true)
    app.App.SetFocus(errorModal)
}

// ConvertToRune converts given integer into rune
func ConvertToRune(i int) rune {
	return rune(i + 'a')
}

// GetLocalTime retrieves local time from given UNIX timestamp
func GetLocalTime(timestamp uint64) string {
    if timestamp == 0 {
        return "Unknown"
    }

    localTime := time.Unix(int64(timestamp), 0).Local()

    return localTime.String()
}
