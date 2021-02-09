package wasm

import (
	"fmt"
	"syscall/js" // ignore warning, build script handles this
)

// Run will start frontend processes, returning an error if there is any.
func Run() error {
	var err error
	fmt.Println("Go Web Assembly")

	SetupFuncMap()

	// Wait for commands on channel
	<-make(chan bool)
	return err
}

// SetupFuncMap will make listed go functions globally accessible to the frontend.
func SetupFuncMap() {
	funcMap := make(map[string]js.Func)

	funcMap["echo"] = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		for i, v := range args {
			fmt.Printf("%d: %v\n", i, v)
		}
		return ""
	})

	for k, v := range funcMap {
		js.Global().Set(k, v)
	}
}
