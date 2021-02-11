class App {
  hello = helloFunc
}

function helloFunc() {
  console.log("hello world")
}

window["app"] = new App()