import axios from 'axios'

declare namespace tasks {
  // Tasks holds exported methods related to tasks.
  export interface Tasks {
    Create: (e: Event) => void
    Complete: (e: Event, id: number) => void
  }
}
export { tasks }

// NewTasks is a constructor function that returns Tasks, along with all methods of that type.
export function NewTasks(): tasks.Tasks {
  const t: tasks.Tasks = {
    Create: async (e: Event) => {
      e.preventDefault()
      type formInput = HTMLElement & { value?: string }
      const input: formInput = document.getElementById("newTask")!
      if (typeof input.value === 'undefined') {
        return null
      }
      const task = input.value
      try {
        const res = await axios.request({
          method: 'POST',
          url: "/tasks",
          data: {
            task: task
          }
        })
        // TODO push task to list
        const taskTpl = taskTplGet(res.data.id, task)
        const tasks = document.getElementById("tasks")
        tasks!.innerHTML += taskTpl
      } catch (err) {
        console.log(err)
      }
    },
    Complete: async (e: Event, id: number) => {
      e.preventDefault()
      console.log("TODO: implement")
    }
  }
  return t
}

// Templates

function taskTplGet(id: number, task: string): string {
  const html = `<div>
  <span>
    <input type="checkbox" onchange="app.Tasks.Complete(event, ${id})">${task}
  </span>
</div>`
return html
}