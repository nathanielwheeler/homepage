{{define "main"}}
<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>Home</title>
  <script>
    const go = new Go();
    WebAssembly
      .instantiateStreaming(fetch("json.wasm"), go.importObject)
      .then((result) => {
        go-run(result.instance);
      });
  </script>
</head>
<body>
  
</body>
</html>
{{end}}