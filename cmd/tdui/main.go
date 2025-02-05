package main

import "github.com/CesareDev/tdui/cmd/app"

func main() {
    var app app.App
    app.Init()
    app.Setup()
    app.Run()
}
