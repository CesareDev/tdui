package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func main() {

    app := tview.NewApplication()
    grid := tview.NewGrid()

    app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
        if event.Rune() == 'q' {
            modal := tview.NewModal().
                SetText("Do you want to quit the application?").
                AddButtons([]string{"Quit", "Cancel"}).
                SetDoneFunc(func(buttonIndex int, buttonLabel string) {
                    if buttonLabel == "Quit" {
                        app.Stop()
                    }
                    if buttonLabel == "Cancel" {
                        app.SetRoot(grid, true)
                    }
                })
            app.SetRoot(modal, true)
        }
        return event
    })

    grid.SetColumns(-1, -1, -1)
    grid.SetRows(-1, -2, -1)

    body := tview.NewTextView()
    body.SetBorder(true).SetTitle("Body")

    input := tview.NewInputField()
    input.SetLabel("Enter a new task ")
    input.SetDoneFunc(func(key tcell.Key) { 
        body.SetText(input.GetText())
        input.SetText("")
    })
    input.SetTitle("To Do").SetBorder(true)

    grid.AddItem(input, 3, 0, 1, 3, 0, 0, false)
    grid.AddItem(body, 0, 0, 3, 3, 0, 0, false)

    err := app.SetRoot(grid, true).EnableMouse(true).Run()

    if err != nil {
        panic(err)
    }
}
