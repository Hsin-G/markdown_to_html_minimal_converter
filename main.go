package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"

	"github.com/yuin/goldmark"
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
	md := goldmark.New()
	err = md.Convert(input, &buf)
	if err != nil {
		fmt.Println("Error during conversion: ", err)
		return
	}

	// Write in the HTML file
	if state == true {
		err = os.WriteFile(os.Args[3], buf.Bytes(), 0644)
	}else{
		err = os.WriteFile("Output.html", buf.Bytes(), 0644)
	}
	if err != nil {
		fmt.Println("Write error: ", err)
		return
	}
	if state == true {
		fmt.Println("HTML file: ", os.Args[3])
	}else{
		fmt.Println("HTML file: output.html")
	}
}

