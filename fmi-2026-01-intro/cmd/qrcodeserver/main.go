package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var templ = template.Must(template.New("qr").Parse(templateStr))

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "Hello World!")
}

// QR http handler function return the HTML template with QR code for the input link
func QR(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	templ.Execute(w, req.FormValue("s"))
}

const templateStr = `
<html>
<head>
<title>QR Link Generator</title>
</head>
<body>
{{if .}}
<img src="https://qrcode.tec-it.com/API/QRCode?data={{.}}" />
<br>
{{.}}
<br>
<br>
{{end}}
<form action="/" name=f method="GET"><input maxLength=1024 size=70
name=s value="" title="Text to QR Encode"><input type=submit
value="Show QR" name=qr>
</form>
</body>
</html>
`

func main() {
	http.HandleFunc("/", QR) //http.Handle("/", http.HandlerFunc(QR))
	fmt.Println("Starting server at port 8001")
	err := http.ListenAndServe(":8001", nil)
	log.Fatal(err)
}
