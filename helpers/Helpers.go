package helpers

import (
	"time"

	"github.com/rivo/tview"
)

func ErrorModal(err error, root tview.Primitive, App *tview.Application) {
    errorModal := tview.NewModal().
    SetText("Error: " + err.Error()).
    AddButtons([]string{"Got it"}).
    SetDoneFunc(func(buttonIndex int, buttonLabel string) {
            App.SetRoot(root, true).SetFocus(root)
    })

    App.SetRoot(errorModal, true).SetFocus(errorModal)
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
