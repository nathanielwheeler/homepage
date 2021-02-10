package wasm

import (
	"fmt"
	"syscall/js"
)

// NewTaskController will return a POJO controller that will handle tasks on the frontend.
func (a *App) NewTaskController() js.Value {
	tc := make(map[string]interface{})
	fmt.Println("Hello from Tasks!")
	tc["createTask"] = a.TaskCreate()

	return js.ValueOf(tc)
}

// TaskCreate will make an async call to the server for a new task and then update the page.
func (a *App) TaskCreate() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		// task := make(map[string]interface{})

		// POST /api/tasks
		// a.Axios.Call("post", a.BaseURL+"tasks", js.ValueOf(task)).
		// 	Call("then")

		return "TODO implement"
	})
}
