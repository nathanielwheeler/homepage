{{define "main"}}
<!DOCTYPE html>
<html lang="en">

<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>Home</title>
  <link rel="stylesheet" href="/styles/style.css">
  <link rel="shortcut icon" href="assets/favicon.ico" type="image/x-icon">
  <script src="/app/main.js"></script>
</head>

<body class="grid-center-container">

  <header class="grid-top"></header>
  <main class="grid-center">
    <div class="grid-cards">
      <div class="card todo">
        <h3 class="todo-title">TODO</h3>
        <div id="tasks">
          {{if .}}
            {{range .}}
              {{if .}}
              {{template "task" .}}
              {{end}}
            {{end}}
          {{end}}
        </div>
        {{template "taskForm"}}
      </div>
    </div>
  </main>
  <footer class="grid-bottom"></footer>

</body>

</html>
{{end}}

{{define "taskForm"}}
<form onsubmit="app.Tasks.Create(event)">
  <input type="text" name="task" id="newTask">
  <button type="submit">submit</button>
</form>
{{end}}

{{define "task"}}
<div>
  <span>
    <input type="checkbox" onchange="app.Tasks.Complete(event, {{.ID}})">{{.Name}}
  </span>
</div>
{{end}}