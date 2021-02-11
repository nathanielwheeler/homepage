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
      console.log(e.target)

      // interface request {
      //   task: string
      // }
      // const req: request = {
      //   task: taskName,
      // }
      // try {
      //   const res = await axios.post("/tasks", req)
      //   // TODO push task to list
      // } catch (error) {
      //   console.log(error)
      // }
    }
  }
  return t
}