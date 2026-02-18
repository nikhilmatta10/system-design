package main

import "fmt"

type CSVParser struct{}

func (cp *CSVParser) parse() {
	cp.openFile()
	fmt.Println("Parsing a CSV File")
	cp.closeFile()
}

func (cp *CSVParser) openFile() {
	fmt.Println("Opening file")
}

func (cp *CSVParser) closeFile() {
	fmt.Println("Closing file")
}

type HTMLParser struct{}

func (cp *HTMLParser) parse() {
	cp.openFile()
	fmt.Println("Parsing a HTML File")
	cp.closeFile()
}

func (cp *HTMLParser) openFile() {
	fmt.Println("Opening file")
}

func (cp *HTMLParser) closeFile() {
	fmt.Println("Closing file")
}


func main(){
	csvParser := CSVParser{}
	csvParser.parse()
	jsonParser := HTMLParser{}
	jsonParser.parse()
}