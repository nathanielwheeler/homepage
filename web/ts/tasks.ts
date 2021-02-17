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
      let input = <HTMLInputElement> document.getElementById("newTask")!
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
        // Add task to list and clear input
        const taskTpl = taskTplGet(res.data.id, task)
        const tasks = document.getElementById("tasks")
        tasks!.innerHTML += taskTpl
        input.value = ""
      } catch (err) {
        // TODO display error to user
        console.log(err)
      }
    },
    Complete: async (e: Event, id: number) => {
      e.preventDefault()
      let task = document.getElementById(`t-${id}`)!
      try {
        const res = await axios.request({
          method: 'POST',
          url: `/tasks/${id}/complete`,
          data: {
            id: id
          }
        })
        task.remove()
      } catch (err) { 
        // TODO display error to user
        console.log(err)
        let taskCheck = <HTMLInputElement> task.querySelector(".task-check")!
        taskCheck.checked = false
      }
    }
  }
  return t
}

// Templates

function taskTplGet(id: number, task: string): string {
  const html = `<div id="t-${id}">
  <span>
    <input class="task-check" type="checkbox" onchange="app.Tasks.Complete(event, ${id})">${task}
  </span>
</div>`
return html
}