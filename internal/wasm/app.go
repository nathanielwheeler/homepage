package wasm

import (
	"fmt"
	"syscall/js"
)

// NewApp is a constructor function for a javascript object
func NewApp() js.Value {
  a := make(map[string]interface{})
  app := js.ValueOf(a)

  app.Set("acknowledge", Acknowledge())

  return app
}

// SetupApp will attach the App to the browser's window.
func SetupApp() (err error) {
	fmt.Println("Hello from SetupApp!")

	jsWin := js.Global().Get("window") // truthy
  jsWin.Set("app", NewApp())

	return err
}

// Acknowledge is a test function to see if I can attach methods to the frontend. Currently works.
func Acknowledge() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		fmt.Println("Method acknowledged")
		return "Success!"
	})
}
