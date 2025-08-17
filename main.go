package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"		//some markdown extension
	"github.com/yuin/goldmark-highlighting" //highlighting syntax extension
)

func checkArgs(state *bool)bool{
	argc := len(os.Args)

	if argc < 2 || argc > 4 {
		return false
	}
	if !strings.HasSuffix(os.Args[1], ".md") && !strings.HasSuffix(os.Args[1], ".mk"){
		return false 
	}
	if argc > 2 && (os.Args[2] != "-o" || !strings.HasSuffix(os.Args[3], ".html")){
		return false
	}
	if argc > 2{
		*state = true
	}
	return true
}

func argsError(){
	fmt.Println("Error: Invalid Argument!\n")
	fmt.Println("-Usage:	 ./<executable> <markdown_file>")
	fmt.Println("	 ./<executable> <markdown_file> -o <output_file.html>\n")
}

func render(htmlContent string, state bool) bool{
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
	var render string

	htmlContent = fmt.Sprintf("%s", htmlContent)
	input, err := os.ReadFile("./css/standard.css")
	if (err != nil){
		fmt.Println("Read css error: ", err)
		render = fmt.Sprintf(htmlT, "", htmlContent)
	}else{
		css := fmt.Sprintf("%s", string(input))
		render = fmt.Sprintf(htmlT, css, htmlContent)
	}
	
	if state == true {
		err = os.WriteFile(os.Args[3], []byte(render), 0644)
	}else{
		err = os.WriteFile("Output.html", []byte(render), 0644)
	}
	if err != nil {
		fmt.Println("Write error: ", err)
		return false
	}
	//os.WriteFile("test.html", []byte(render), 0644)
	return true
}

func main() {
	//check if argument is valid
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

	// Convert Markdown -> HTML
	var buf bytes.Buffer
	md := goldmark.New(goldmark.WithExtensions(extension.GFM,
	 highlighting.NewHighlighting(highlighting.WithStyle("monokai"))))
	err = md.Convert(input, &buf)
	if err != nil {
		fmt.Println("Error during conversion: ", err)
		return
	}

	// Write in the HTML file
	outputHTML:= buf.Bytes()
	if !render(string(outputHTML), state){
		return
	}
	if state == true {
		fmt.Println("HTML file: ", os.Args[3])
	}else{
		fmt.Println("HTML file: output.html")
	}
}

