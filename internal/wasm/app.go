package wasm

import (
	"fmt"
	"syscall/js"
)

// SetupApp will attach the App to the browser's window.
func SetupApp() (err error) {
	fmt.Println("Hello from SetupApp!")

	jsWin := js.Global().Get("window") // truthy
	jsWin.Set("app", NewApp())

	return err
}

// App will hold methods and variables for the frontend
type App struct {
	BaseURL string
	Axios   js.Value
}

// NewApp is a constructor function for a POJO that will hold frontend controllers.
func NewApp() js.Value {
	a := App{
		BaseURL: "/api/",
		Axios:   js.Global().Get("axios"),
	}
	app := js.ValueOf(make(map[string]interface{}))

	app.Set("taskController", a.NewTaskController())

	return app
}
