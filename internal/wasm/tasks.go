package wasm

import (
	"fmt"
	"syscall/js"
)

// NewTaskController will return a POJO controller that will handle tasks on the frontend.
func (a *App) NewTaskController() js.Value {
	tc := make(map[string]interface{})
	fmt.Println("Hello from Tasks!")
	tc["taskCreate"] = a.TaskCreate()

	return js.ValueOf(tc)
}

// TaskCreate will make an async call to the server for a new task and then update the page.
func (a *App) TaskCreate() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
    // Get form elements
    event := args[0]
    event.Call("preventDefault")

		req := make(map[string]interface{})
    req["task"] = jsGet(event, "target", "elements", "value")

		// POST /api/tasks
		a.Axios.Call("post", a.BaseURL+"tasks", js.ValueOf(req)).
			Call("then", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
        res := args[0].Get("data")
        id := res.Get("id").Int()

        // TODO push changes to DOM
        fmt.Println("id:", id)

        return nil
      })).
      Call("catch", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
        err := args[0]

        if err.Get("response").Truthy() {
          fmt.Println(jsGet(err, "response", "status").String())
          fmt.Println(jsGet(err, "response", "data").String())
        } else if err.Get("request").Truthy() {
          fmt.Println("no response")
          fmt.Println(err.Get("request"))
        } else {
          fmt.Println( err.Get("message").String())
        }

        
        return nil
      }))

		return nil
	})
}
