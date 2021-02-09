{{define "main"}}
<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>Home</title>
  <link rel="stylesheet" href="/app/style.css">
  <link rel="shortcut icon" href="assets/favicon.ico" type="image/x-icon">
  <script src="/static/wasm_exec.js"></script>
  <script>
    const go = new Go();
    WebAssembly
      .instantiateStreaming(fetch("/app/json.wasm"), go.importObject)
      .then((result) => {
        go.run(result.instance);
      });
  </script>
</head>
<body class="grid-center-container">

  <header class="grid-top"></header>
  <main class="grid-center">
    <div class="grid-cards">
      <div class="todo">
        <h3 class="todo-title">TODO</h3>
        <ul>
          <li>One</li>
          <li>Two</li>
          <li>Three</li>
        </ul>
      </div>
    </div>
  </main>
  <footer class="grid-bottom"></footer>
  
</body>
</html>
{{end}}