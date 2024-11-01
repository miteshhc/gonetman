package app

import (
    "github.com/Wifx/gonetworkmanager/v2"
	"github.com/rivo/tview"
)

var NMInstance gonetworkmanager.NetworkManager
var NMSettings gonetworkmanager.Settings

var App = tview.NewApplication()

