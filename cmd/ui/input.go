package ui

import (
	"strconv"

	"github.com/CesareDev/tdui/cmd/utils"
	"github.com/gdamore/tcell/v2"

	"github.com/rivo/tview"
)

type Input struct {
    flex *tview.Flex
    form *tview.Form
    inputText *tview.InputField
    inputDays *tview.DropDown
    inputMonths *tview.DropDown
    selectedDay int
    selectedMonth int
}

func (input *Input) Init() {
    input.form = tview.NewForm()
    input.flex = tview.NewFlex()
    input.inputText = tview.NewInputField()
    input.inputDays = tview.NewDropDown()
    input.inputMonths = tview.NewDropDown()

}

func (input *Input) Setup(app *tview.Application, list *tview.List) {
    days := []string{}
    for i := 1; i <= 31; i++ {
        days = append(days, strconv.Itoa(i))
    }
    months := []string{ "Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec" }

    currentTime := utils.GetCurrentTime() 

    var stringTask string 
    input.selectedDay = currentTime.Day()
    input.selectedMonth = int(currentTime.Month())

    input.inputText.SetLabel("Enter a new task")
    input.inputText.SetChangedFunc(func(text string) { stringTask = text })
    input.inputText.SetFocusFunc(func() { 
        list.SetTitle(" Press Ctrl+L to focus the list ")
        input.form.SetTitle(" Use Tab to navigate bewteen the inputs ")
    })

    input.inputDays.SetLabel("Day: ")
    input.inputDays.SetOptions(days, nil)
    input.inputDays.SetCurrentOption(input.selectedDay - 1)
    input.inputDays.SetFocusFunc(func() { 
        list.SetTitle(" Press Ctrl+L to focus the list ")
        input.form.SetTitle(" Use Tab to navigate bewteen the inputs ")
    })

    input.inputMonths.SetLabel("Month: ")
    input.inputMonths.SetOptions(months, nil)
    input.inputMonths.SetCurrentOption(input.selectedMonth - 1)
    input.inputMonths.SetFocusFunc(func() { 
        list.SetTitle(" Press Ctrl+L to focus the list ")
        input.form.SetTitle(" Use Tab to navigate bewteen the inputs ")
    })
    
    // Separated callbacks to prevent stackoverflow
    input.inputDays.SetSelectedFunc(func(text string, index int) {
        input.selectedDay = index + 1
        UpdateDropdownInput(input.inputDays, currentTime, input.selectedMonth, input.selectedDay)
    })

    input.inputMonths.SetSelectedFunc(func(text string, index int) {
        input.selectedMonth = index + 1
        UpdateDropdownInput(input.inputDays, currentTime, input.selectedMonth, input.selectedDay)
    })

    // Button cannot be inserted via the AddFormItem function
    input.form.AddButton("Insert", func() {
        if stringTask == "" {
            return
        }
        list.AddItem(stringTask + " [@ " + strconv.Itoa(input.selectedDay) + "/" + months[input.selectedMonth - 1] + "]", "", '-', func() {
            tmp := list.GetCurrentItem()
            list.RemoveItem(tmp)
        })
    })
    input.form.AddFormItem(input.inputText)
    input.form.AddFormItem(input.inputDays)
    input.form.AddFormItem(input.inputMonths)

    input.form.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
        if event.Key() == tcell.KeyCtrlL {
            app.SetFocus(list)
            input.form.SetTitle(" Press Ctrl+L to focus the input ")
            list.SetTitle(" Press Enter to delete a task ")
            return nil
        }
        return event
    })
    input.form.SetTitle(" Use Tab to navigate bewteen the inputs ").SetBorder(true)

    input.flex.SetDirection(tview.FlexColumn)
    input.flex.AddItem(input.form, 0, 1, true)
}

func (input *Input) GetFlex() *tview.Flex {
    return input.flex
}

func (input *Input) GetInternalInput() *tview.Form {
    return input.form
}

func UpdateDropdownInput(input *tview.DropDown, currentTime utils.Time, month int, day int) {
    if month == 2 {
        if currentTime.IsLeapYear() && day > 29 {
            input.SetCurrentOption(28)
        } else if day > 28 {
            input.SetCurrentOption(27)
        }
        return
    }
    if day == 31 {
        switch month {
        case 4: input.SetCurrentOption(29)
        case 6: input.SetCurrentOption(29)
        case 9: input.SetCurrentOption(29)
        case 11: input.SetCurrentOption(29)
        }
    }
}
