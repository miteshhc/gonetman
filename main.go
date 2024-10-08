package main

import (
    "github.com/miteshhc/gonetman/app"
    "github.com/miteshhc/gonetman/components"
)

func main() {
    if err := app.App.
            SetRoot(components.Flex, true).
            Run();
            err != nil {
                panic("Failed to run app: " + err.Error())
    }
}
