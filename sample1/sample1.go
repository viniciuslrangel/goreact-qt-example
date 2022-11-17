package main

import (
	"fmt"
	. "github.com/viniciuslrangel/goreact"
	"github.com/viniciuslrangel/goreact-qt"
	. "github.com/viniciuslrangel/goreact-qt/components"
)

type AppState struct {
	count int
}

func main() {

	app := FCS("App", AppState{}, func(state AppState, updateState func(appState AppState)) Node {
		return Window.New(
			GridLayout.New(
				GridCell(0, 0,
					Label.New(LabelProps{
						Text: fmt.Sprintf("Hello World. You Have clicked %d", state.count),
					}),
				),
				GridCell(1, 0,
					Button.New(ButtonProps{
						Label: fmt.Sprintf("Click count %d", state.count),
						OnClicked: func() {
							updateState(AppState{
								count: state.count + 1,
							})
						},
					}),
				),
			),
		)
	})

	err := goreact_qt.Render(app.New())
	if err != nil {
		panic(err)
	}
}
