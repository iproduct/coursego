package main

import (
	"fmt"
	"regexp"
)

func main() {
	markup := `<a href="products.html" class="button">Go to products</a>`
	re, _ := regexp.Compile(`href="(.*?)"`)
	values := re.FindStringSubmatch(markup)
	fmt.Println(values)
	if len(values) > 0 {
		fmt.Println("URL:", values[1]) // prints URL: products.html
	}
}
