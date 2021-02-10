package wasm

import (
	"encoding/json"
	"fmt"
	"syscall/js" // ignore warning, build script handles this
)

// Run will start frontend processes, returning an error if there is any.
func Run() error {
	var err error
	fmt.Println("Go Web Assembly")

	SetupFuncMap()

	err = SetupApp()
	if err != nil {
		return err
	}

	// Wait for commands on channel
	<-make(chan bool)
	return err
}

// SetupFuncMap will make listed go functions globally accessible to the frontend.
func SetupFuncMap() {
	funcMap := make(map[string]js.Func)

	funcMap["echo"] = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		var str string
		for _, v := range args {
			str += v.String()
		}
		return js.ValueOf(str)
	})

	for k, v := range funcMap {
		js.Global().Set(k, v)
	}
}

// HELPERS

func jsGet(v js.Value, elements ...string) js.Value {
	for _, e := range elements {
		v.Get(e)
	}
	return v
}

func prettyJSON(input string) (string, error) {
	var raw interface{}
	if err := json.Unmarshal([]byte(input), &raw); err != nil {
		return "", err
	}
	pretty, err := json.MarshalIndent(raw, "", "  ")
	if err != nil {
		return "", err
	}
	return string(pretty), nil
}
