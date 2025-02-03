package main

import (
	"strconv"

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

    list := tview.NewList()
    inputFlex := tview.NewFlex()
    input := tview.NewForm()

    list.SetSelectedFocusOnly(true)
    list.SetBorder(true)
    list.ShowSecondaryText(false)
    list.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
        if event.Key() == tcell.KeyCtrlL {
            app.SetFocus(input)
            list.SetTitle(" Press Ctrl+L to focus the list ")
            inputFlex.SetTitle(" Use Tab to navigate bewteen the inputs ")
            return nil
        }
        return event
    })
    list.SetTitle(" Press Tab to focus the list ")

    days := []string{}
    for i := 1; i <= 31; i++ {
        days = append(days, strconv.Itoa(i))
    }

    months := []string{}
    for i := 1; i <= 12; i++ {
        months = append(months, strconv.Itoa(i))
    }

    var stringTask string
    var stringDay string
    var stringMonth string

    input.AddInputField("Enter a new task: ", "", 0, nil, func(text string) {
        stringTask = text
    })
    input.AddDropDown("Day: ", days, 0, func(option string, optionIndex int) {
        stringDay = option 
    })
    input.AddDropDown("Month: ", months, 0, func(option string, optionIndex int) {
        stringMonth = option
    })
    input.AddButton("Insert", func() {
        if stringTask == "" {
            return
        }
        list.AddItem(stringTask + " [@ " + stringDay + "/" + stringMonth + "]", "", '-', func() {
            tmp := list.GetCurrentItem()
            list.RemoveItem(tmp)
        })
    })
    input.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
        if event.Key() == tcell.KeyCtrlL {
            app.SetFocus(list)
            inputFlex.SetTitle(" Press Ctrl+L to focus the input ")
            list.SetTitle(" Press Enter to delete a task ")
            return nil
        }
        return event
    })

    inputFlex.SetDirection(tview.FlexRow)
    inputFlex.AddItem(input, 0, 2, true)
    inputFlex.SetBorder(true)
    inputFlex.SetTitle(" Use Tab to navigate bewteen the inputs ")

    flex.AddItem(list, 0, 5, false)
    flex.AddItem(inputFlex, 0, 2, true)

    err := app.SetRoot(flex, true).EnableMouse(true).Run()

    if err != nil {
        panic(err)
    }
}
