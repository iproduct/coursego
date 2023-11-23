package main

import (
	"flag"
	"html/template"
	"net/http"
)

var addr = flag.String("addr", ":8080", "go run ./qrcodes.go address")

var templ = template.Must(template.New("qr").Parse(templateStr))

func main() {
	flag.Parse()
	http.Handle("/", http.HandlerFunc(QR))
}

func QR(w http.ResponseWriter, r *http.Request) {
	templ.Execute(w, r.FormValue("qr"))
}

const templateStr = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>QR Link Generator</title>
</head>
<body>
	{{if .}}
		<img src="http://chart.apis.google.com/chart?chs=300x300&cht=qr&choe=UTF-8&chl={{.}}">
		<br>
		<br>
		<br>
		{{.}}
	{{end}}
	<form action="/" name="f" method="GET">
		<input maxLength="1024" size="80" name="qr" value="" placeholder="Text to QR to encode">
		<input type="submit" value="Show QR Code">
	</form>
</body>
</html>
`
