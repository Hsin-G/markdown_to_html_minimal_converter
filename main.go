package main

import (
	"bytes"
	"fmt"
	"os"	

	// goldmark
	"github.com/yuin/goldmark"
	// some markdown extension
	"github.com/yuin/goldmark/extension"
	// highlighting syntax extension
	"github.com/yuin/goldmark-highlighting" 
)

// config sruct
type config struct{
	CSS_file  string  `json:"css_file"`
	Highlight string  `json:"hightlight_style"`
}

// return an HTML template
func html_template() string {
	htmlT := `<!DOCTYPE html>
<html>
<head>
	<meta charset="UTF-8">
	<style>
%s
	</style>
</head>
<body>
	<div class="content" >
%s
	</div>
</body>
</html>
`
	return htmlT
}

func render(htmlContent string, state bool, conf config) bool{

	var render string
	// transform []byte into a string
	htmlContent = fmt.Sprintf("%s", htmlContent)
	// Read CSS file
	input, err := os.ReadFile(conf.CSS_file)
	if (err != nil){
		fmt.Println("Read css error: ", err)
		render = fmt.Sprintf(html_template(), "", htmlContent)
	}else{
		css := fmt.Sprintf("%s", string(input))
		render = fmt.Sprintf(html_template(), css, htmlContent)
	}
	// with or without -o option
	if state == true {
		err = os.WriteFile(os.Args[3], []byte(render), 0644)
	}else{
		err = os.WriteFile("Output.html", []byte(render), 0644)
	}
	if err != nil {
		fmt.Println("Write error: ", err)
		return false
	}
	return true
}

func main() {
	// check if argument is valid
	state := false 
	if !checkArgs(&state) {
		argsError()
		return
	}
	
	// Read Markdown file
	input, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("Read error: ", err)
		return
	}

	// get config
	conf := get_config()

	// Convert Markdown -> HTML
	var buf bytes.Buffer
	md := goldmark.New(goldmark.WithExtensions(extension.GFM,
	 highlighting.NewHighlighting(highlighting.WithStyle(conf.Highlight))))
	err = md.Convert(input, &buf)
	if err != nil {
		fmt.Println("Error during conversion: ", err)
		return
	}

	// Creat and Write in the HTML file
	outputHTML:= buf.Bytes()
	if !render(string(outputHTML), state, conf){
		return
	}
	if state == true {
		fmt.Println("HTML file: ", os.Args[3])
	}else{
		fmt.Println("HTML file: output.html")
	}
}

