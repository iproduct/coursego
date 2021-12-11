package main

import ("html/template"; "log";	"os")

const textTempl =
`{{len .}} books:
{{range .}}----------------------------------------
ID: {{.ID}}
Title: {{.Title | printf "%.64s"}}
{{end}}`

func main() {
	tmpl := template.Must(template.New("mytext").Parse(textTempl))
	//if err != nil {
	//	log.Fatal("Error Parsing template: ", err)
	//	return
	//}
	err1 := tmpl.Execute(os.Stdout, goBooks)
	if err1 != nil {
		log.Fatal("Error executing template: ", err1)
	}
}