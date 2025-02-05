package ui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type List struct {
    flex *tview.Flex
    innerList *tview.List
}

func (list *List) Init() {
    list.innerList = tview.NewList()
    list.flex = tview.NewFlex()
}

func (list *List) Setup(app *tview.Application, input *tview.Form) {
    list.innerList.SetSelectedFocusOnly(true)
    list.innerList.ShowSecondaryText(false)
    list.innerList.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
        if event.Key() == tcell.KeyCtrlL {
            app.SetFocus(input)
            list.innerList.SetTitle(" Press Ctrl+L to focus the list ")
            input.SetTitle(" Use Tab to navigate bewteen the inputs ")
            return nil
        }
        return event
    })
    list.innerList.SetFocusFunc(func() {
        input.SetTitle(" Press Ctrl+L to focus the input ")
        list.innerList.SetTitle(" Press Enter to delete a task ")
    })
    list.innerList.SetTitle(" Press Tab to focus the list ").SetBorder(true)

    list.flex.SetDirection(tview.FlexColumn)
    list.flex.AddItem(list.innerList, 0, 1, false)
}

func (list *List) GetFlex() *tview.Flex {
    return list.flex
}

func (list *List) GetInternalList() *tview.List {
    return list.innerList
}
