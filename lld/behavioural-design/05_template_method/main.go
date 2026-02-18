package main

import "fmt"

type ParserTemplateMethods struct{}

func (pt *ParserTemplateMethods) openFile() {
	fmt.Println("file opening")
}

func (pt *ParserTemplateMethods) closeFile() {
	fmt.Println("file closing")
}

type ParserTemplate interface {
	parse()
}

type JSONParser struct {
	ParserTemplateMethods
}

func (jp *JSONParser) parse() {
	jp.closeFile()
	fmt.Println("parsing json file")
	jp.closeFile()
}

type HTMLParser struct {
	ParserTemplateMethods
}

func (htmlParser *HTMLParser) parse() {
	htmlParser.openFile()
	fmt.Println("parsing html format file")
	htmlParser.closeFile()
}

func main(){
	csvParser := JSONParser{}
	csvParser.parse()
	jsonParser := HTMLParser{}
	jsonParser.parse()
}