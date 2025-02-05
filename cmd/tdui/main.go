package main

import (
	"strconv"
    "time"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func IsLeapYear(date time.Time) bool {
    if date.Year() % 400 == 0 {
        return true
    }
    return date.Year() % 4 == 0 && date.Year() % 100 != 0
}

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
    input := tview.NewForm()

    list.SetSelectedFocusOnly(true)
    list.ShowSecondaryText(false)
    list.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
        if event.Key() == tcell.KeyCtrlL {
            app.SetFocus(input)
            list.SetTitle(" Press Ctrl+L to focus the list ")
            input.SetTitle(" Use Tab to navigate bewteen the inputs ")
            return nil
        }
        return event
    })
    list.SetFocusFunc(func() {
        input.SetTitle(" Press Ctrl+L to focus the input ")
        list.SetTitle(" Press Enter to delete a task ")
    })
    list.SetTitle(" Press Tab to focus the list ").SetBorder(true)

    days := []string{}
    for i := 1; i <= 31; i++ {
        days = append(days, strconv.Itoa(i))
    }
    months := []string{ "Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec" }

    currentTime := time.Now().Local()

    var stringTask string 
    var selectedDay int = currentTime.Day()
    var selectedMonth int = int(currentTime.Month())

    inputText := tview.NewInputField()
    inputText.SetLabel("Enter a new task")
    inputText.SetChangedFunc(func(text string) { stringTask = text })
    inputText.SetFocusFunc(func() { 
        list.SetTitle(" Press Ctrl+L to focus the list ")
        input.SetTitle(" Use Tab to navigate bewteen the inputs ")
    })

    inputDays := tview.NewDropDown()
    inputDays.SetLabel("Day: ")
    inputDays.SetOptions(days, nil)
    inputDays.SetCurrentOption(selectedDay - 1)
    inputDays.SetFocusFunc(func() { 
        list.SetTitle(" Press Ctrl+L to focus the list ")
        input.SetTitle(" Use Tab to navigate bewteen the inputs ")
    })

    inputMonths := tview.NewDropDown()
    inputMonths.SetLabel("Month: ")
    inputMonths.SetOptions(months, nil)
    inputMonths.SetCurrentOption(selectedMonth - 1)
    inputMonths.SetFocusFunc(func() { 
        list.SetTitle(" Press Ctrl+L to focus the list ")
        input.SetTitle(" Use Tab to navigate bewteen the inputs ")
    })
    
    // Separated callbacks to prevent stackoverflow
    inputDays.SetSelectedFunc(func(text string, index int) {
        selectedDay = index + 1
        if selectedMonth == 2 {
            if IsLeapYear(currentTime) && selectedDay > 29 {
                inputDays.SetCurrentOption(28)
            } else if selectedDay > 28 {
                inputDays.SetCurrentOption(27)
            }
            return
        }
        if selectedDay == 31 {
            switch selectedMonth {
            case 4: inputDays.SetCurrentOption(29)
            case 6: inputDays.SetCurrentOption(29)
            case 9: inputDays.SetCurrentOption(29)
            case 11: inputDays.SetCurrentOption(29)
            }
        }
    })

    inputMonths.SetSelectedFunc(func(text string, index int) {
        selectedMonth = index + 1
        if selectedMonth == 2 {
            if IsLeapYear(currentTime) && selectedDay > 29 {
                inputDays.SetCurrentOption(28)
            } else if selectedDay > 28 {
                inputDays.SetCurrentOption(27)
            }
            return
        }
        if selectedDay == 31 {
            switch selectedMonth {
            case 4: inputDays.SetCurrentOption(29)
            case 6: inputDays.SetCurrentOption(29)
            case 9: inputDays.SetCurrentOption(29)
            case 11: inputDays.SetCurrentOption(29)
            }
        }
    })

    // Button cannot be inserted via the AddFormItem function
    input.AddButton("Insert", func() {
        if stringTask == "" {
            return
        }
        list.AddItem(stringTask + " [@ " + strconv.Itoa(selectedDay) + "/" + months[selectedMonth - 1] + "]", "", '-', func() {
            tmp := list.GetCurrentItem()
            list.RemoveItem(tmp)
        })
    })
    input.AddFormItem(inputText)
    input.AddFormItem(inputDays)
    input.AddFormItem(inputMonths)

    input.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
        if event.Key() == tcell.KeyCtrlL {
            app.SetFocus(list)
            input.SetTitle(" Press Ctrl+L to focus the input ")
            list.SetTitle(" Press Enter to delete a task ")
            return nil
        }
        return event
    })
    input.SetTitle(" Use Tab to navigate bewteen the inputs ").SetBorder(true)
    
    inputFlex := tview.NewFlex()
    inputFlex.SetDirection(tview.FlexColumn)
    inputFlex.AddItem(input, 0, 1, true)

    listFlex := tview.NewFlex()
    listFlex.SetDirection(tview.FlexColumn)
    listFlex.AddItem(list, 0, 1, false)

    flex.AddItem(listFlex, 0, 5, false)
    flex.AddItem(inputFlex, 0, 2, true)

    err := app.SetRoot(flex, true).EnableMouse(true).Run()

    if err != nil {
        panic(err)
    }
}
