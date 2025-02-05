package app

import (
	"github.com/CesareDev/tdui/cmd/ui"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type App struct {
    internalApp *tview.Application
    flex *tview.Flex
    tasklist ui.List
    input ui.Input
}

func (app *App) Init() {
    app.internalApp = tview.NewApplication()
    app.flex = tview.NewFlex()
    app.tasklist.Init()
    app.input.Init()
}

func (app *App) Setup() {
    app.tasklist.Setup(app.internalApp, app.input.GetInternalInput())
    app.input.Setup(app.internalApp, app.tasklist.GetInternalList())

    app.internalApp.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
        if event.Key() == tcell.KeyEsc {
            modal := tview.NewModal().
                SetText("Do you want to quit the application?").
                AddButtons([]string{"Quit", "Cancel"}).
                SetDoneFunc(func(buttonIndex int, buttonLabel string) {
                    if buttonLabel == "Quit" {
                        app.internalApp.Stop()
                    }
                    if buttonLabel == "Cancel" {
                        app.internalApp.SetRoot(app.flex, true)
                    }
                })
            app.internalApp.SetRoot(modal, true)
        }
        return event
    })
    app.flex.SetDirection(tview.FlexRow)

    app.flex.AddItem(app.tasklist.GetFlex(), 0, 5, false)
    app.flex.AddItem(app.input.GetFlex(), 0, 2, true)
}

func (app App) Run() {
    err := app.internalApp.SetRoot(app.flex, true).EnableMouse(true).Run()
    if err != nil { panic(err) }
}
