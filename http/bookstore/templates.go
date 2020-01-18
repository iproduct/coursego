package main

const tmplAllBooksStr = `
<html>
<head>
<title>Bookstore</title>
<style>
table, th, td {
	border-collapse: collapse;
	border:1px solid lightblue;
	padding: 3px 5px;
}
</style>
</head>
<body>
<h1>Bookstore</h1>
{{if .}}
<h2>Books List ({{len .}} books)</h2>
<table>
<tr><th>ID</th><th>Title</th></tr>
{{range .}}
<tr><td>{{.ID}}</td><td>{{.Title}}</td></tr>
{{end}}
</table>
<br>
<br>
<br>
{{end}}
<form action="/" name=f method="GET">
<input maxLength=1024 size=70 name="bookId" value="" title="Book ID">
<input type="submit" value="Show Details" name="showDetails">
</form>
</body>
</html>
`
