package main

import (
	"flag"
	"html/template"
	"log"
	"net/http"
)

var addr = flag.String("addr", ":8080", "09-http service address") // Q=17, R=18

var templ = template.Must(template.New("qr").Parse(templateStr))

func main() {
	flag.Parse()
	http.Handle("/", http.HandlerFunc(QR))
	log.Fatal(http.ListenAndServe(*addr, nil))
}

// QR 09-http handler function return the HTML template with QR code for the input link
func QR(w http.ResponseWriter, req *http.Request) {
	log.Println("s = ", req.FormValue("s"))
	templ.Execute(w, req.FormValue("s"))
}

const templateStr = `
<html>
<head>
<title>QR Link Generator</title>
</head>
<body>
{{if .}}
<img src="https://qrcode.tec-it.com/API/QRCode?data={{.}}" style="width:300px;height:300px"/>
<br>
{{.}}
<br>
<br>
{{end}}
<form action="/" name=f method="POST">
<input maxLength=1024 size=70 name="s" value="" title="Text to QR Encode">
<input type=submit value="Show QR" name=qr>
</form>
</body>
</html>
`
