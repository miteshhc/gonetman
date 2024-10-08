package components

import (
    "github.com/rivo/tview"
)

var Flex = tview.NewFlex()

func init() {
    Flex.AddItem(NewMainMenu(), 0, 1, true)
}
