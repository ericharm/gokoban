package states

import (
	"github.com/rthornton128/goncurses"
	"log"
	"sync"
)

type Application struct {
	states  []State
	running bool
	window  *goncurses.Window
}

var appInstance *Application
var once sync.Once

func GetApplication() *Application {
	once.Do(func() {
		appInstance = &Application{
			states:  []State{},
			running: true,
		}
	})

	return appInstance
}

func (app *Application) SetWindow(window *goncurses.Window) {
	app.window = window
}

func (app *Application) GetWindow() *goncurses.Window {
	return app.window
}

func (app *Application) Run(window *goncurses.Window) {
	if len(app.states) < 1 {
		log.Fatal("No states to run")
	}

	for app.running {
		currentState := app.states[len(app.states)-1]
		window.Clear()
		currentState.Draw(window)
		window.Refresh()
		char := window.GetChar()
		if char == 'q' {
			app.running = false
		} else {
			currentState.HandleInput(char)
		}
	}
}

func (app *Application) PushState(state State) {
	app.states = append(app.states, state)
}

func (app *Application) popState() {
	if len(app.states) > 0 {
		app.states = app.states[:len(app.states)-1]
	}
}

func (app *Application) SwapState(state State) {
	if len(app.states) > 0 {
		app.popState()
	}
	app.PushState(state)
}
