interface App {
  hello: () => void
}

let newApp = function (): App {
  const app: App = {
    hello: () => { console.log("hello world") }
  }
  return app
}

window["app"] = newApp()