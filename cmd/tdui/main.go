package main

import (
    "github.com/gdamore/tcell/v2"
    "github.com/rivo/tview"
)

func main() {

    app := tview.NewApplication()
    flex := tview.NewFlex().SetDirection(tview.FlexRow)

    app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
        if event.Key() == tcell.KeyEsc {
            modal := tview.NewModal().
                SetText("Do you want to quit the application?").
                AddButtons([]string{"Quit", "Cancel"}).
                SetDoneFunc(func(buttonIndex int, buttonLabel string) {
                    if buttonLabel == "Quit" {
                        app.Stop()
                    }
                    if buttonLabel == "Cancel" {
                        app.SetRoot(flex, true)
                    }
                })
            app.SetRoot(modal, true)
        }
        return event
    })

    title := tview.NewTextView()
    list := tview.NewList()
    input := tview.NewInputField()

    title.SetText("tdui - To Do interactive application")
    title.SetTextAlign(tview.AlignCenter)
    title.SetBorder(true)

    list.SetSelectedFocusOnly(true)
    list.SetBorder(true)
    list.ShowSecondaryText(false)
    list.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
        if event.Key() == tcell.KeyTab {
            app.SetFocus(input)
            return nil
        }
        return event
    })

    input.SetLabel("Enter a new task: ")
    input.SetDoneFunc(func(key tcell.Key) {
        if input.GetText() == "" {
            return
        }
        list.AddItem(input.GetText(), "", '-', func() {
            tmp := list.GetCurrentItem()
            list.RemoveItem(tmp)
        })
        input.SetText("")
    })
    input.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
        if event.Key() == tcell.KeyTab {
            app.SetFocus(list)
            return nil
        }
        return event
    })
    input.SetBorder(true)

    flex.AddItem(title, 0, 1, false)
    flex.AddItem(list, 0, 6, false)
    flex.AddItem(input, 0, 1, true)

    err := app.SetRoot(flex, true).EnableMouse(true).Run()

    if err != nil {
        panic(err)
    }
}
