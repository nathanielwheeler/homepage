import { tasks } from './tasks'
import { NewTasks } from './tasks'

interface App {
  Tasks: tasks.Tasks
  hello: () => void
}

let newApp = function (): App {
  const app: App = {
    Tasks: NewTasks(),
    hello: () => { console.log("hello world") },
  }
  return app
}

// @ts-ignore
window["app"] = newApp()