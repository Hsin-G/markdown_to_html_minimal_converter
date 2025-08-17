package main

import (
	"fmt"
	"os"
	"strings"
	// json
	"encoding/json"
)

//decode json config
func get_config() config {
	var conf config

	// check if config file exist
	_, check := os.Stat("./conf.json")
	if os.IsNotExist(check){
		check = os.WriteFile("./conf.json", []byte(`{"css_file":"./css/standard.css","hightlight_style":"monokai"}`), 0644)
		if check != nil{
			fmt.Println("Writing in json file Error:", check)
			os.Exit(1)
		}
	}
	// read the json file
	data, err := os.ReadFile("./conf.json")
	if err != nil {
		fmt.Println("Reading json file error:", err)
		conf.Highlight = "monokai"
		return conf
	}
	// decode json 
	err = json.Unmarshal(data, &conf)
	if err != nil {
		fmt.Println("Unmarshal Error: ", err)
		os.Exit(1)
	}
	return conf
}

// check programme argument
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

// Invalid Arguments Error message
func argsError(){
	fmt.Println("Error: Invalid Argument!\n")
	fmt.Println("-Usage:	 ./<executable> <markdown_file>")
	fmt.Println("	 ./<executable> <markdown_file> -o <output_file.html>\n")
}
