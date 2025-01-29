package main

import(
    "github.com/rivo/tview"
)

func main() {

    app := tview.NewApplication()

    grid := tview.NewGrid()
    grid.SetColumns(-1)
    grid.SetRows(-1, -2, -1)

    head := tview.NewTextView().SetTitle("Head").SetBorder(true)
    body := tview.NewTextView().SetTitle("Body").SetBorder(true)
    footer := tview.NewTextView().SetTitle("Footer").SetBorder(true)

    grid.AddItem(head, 0, 0, 1, 1, 0, 0, false)
    grid.AddItem(body, 1, 0, 2, 1, 0, 0, false)
    grid.AddItem(footer, 3, 0, 1, 1, 0, 0, false)

    err := app.SetRoot(grid, true).EnableMouse(true).Run()

    if err != nil {
        panic(err)
    }
}
