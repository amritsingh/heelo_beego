<!DOCTYPE html>

<html>
<head>
  <title>Beego</title>
  <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
</head>

<body>
  {{if .Age}}
    <h1>My name is {{.Name}}</h1>
    <div>
      My age is {{.Age}}
    </div>
  {{else}}
    <form action="/hello" method="GET">
      <label for="name">Enter your name:</label><br>
      <input type="text" id="name" name="name"><br>
      <br>
      <label for="age">Enter your age:</label><br>
      <input type="text" id="age" name="age"><br>
      <br>
      <input type="submit" value="Submit">
    </form>
  {{end}}
</body>
</html>
