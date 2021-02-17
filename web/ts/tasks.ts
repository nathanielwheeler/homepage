import axios from 'axios'

declare namespace tasks {
  // Tasks holds exported methods related to tasks.
  export interface Tasks {
    Create: (e: Event) => void
  }
}
export { tasks }

// NewTasks is a constructor function that returns Tasks, along with all methods of that type.
export function NewTasks(): tasks.Tasks {
  const t: tasks.Tasks = {
    Create: async (e: Event) => {
      e.preventDefault()
      type formInput = HTMLElement & { value?: string }
      let input: formInput = document.getElementById("newTask")!
      if (typeof input.value === 'undefined') {
        return null
      }
      try {
        const res = await axios.request({
          method: 'POST',
          url: "/tasks",
          data: {
            task: input.value
          }
        })
        console.log(res)
        // TODO push task to list
      } catch (error) {
        console.log(error)
      }
    }
  }
  return t
}