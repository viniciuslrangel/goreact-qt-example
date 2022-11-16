package main

import (
	. "github.com/viniciuslrangel/goreact"
	"github.com/viniciuslrangel/goreact-qt"
	. "github.com/viniciuslrangel/goreact-qt/components"
)

func main() {

	app := FC("App", func() Node {
		return Window.New(
			Button.New(ButtonProps{
				Label: "Click me",
			}),
		)
	})

	err := goreact_qt.Render(app.New())
	if err != nil {
		panic(err)
	}
}
