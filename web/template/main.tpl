{{define "main"}}
<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>Home</title>
  <link rel="stylesheet" href="/app/style.css">
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
<body>
  
</body>
</html>
{{end}}