package main

const tmplAllBooksStr = `
<!DOCTYPE html>
<html>
<head>
	<title>Bookstore</title>
	<!--Import Google Icon Font-->
	<link href="https://fonts.googleapis.com/icon?family=Material+Icons" rel="stylesheet">
	<!--Import materialize.css-->
	<link type="text/css" rel="stylesheet" href="static/css/materialize.min.css"  media="screen,projection"/>
	<!--Import main.css-->
	<link type="text/css" rel="stylesheet" href="static/css/main.css"  media="screen,projection"/>

	<!--Let browser know website is optimized for mobile-->
	<meta name="viewport" content="width=device-width, initial-scale=1.0"/>
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
	<input maxLength=1024 size=70 name="bookId" value="" title="Book ID" placeholder="Book ID">
	<input type="submit" value="Show Details" name="showDetails">
	</form>

	<!--JavaScript at end of body for optimized loading-->
	<script type="text/javascript" src="static/js/materialize.min.js"></script>
</body>
</html>`
