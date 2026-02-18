package main

import "fmt"

type TextEditor struct{}

func (te *TextEditor) boldText() {
	fmt.Println("Text has been bolded. ")
}

func (te *TextEditor) italicizeText() {
	fmt.Println("Text has been italicized. ")
}

func (te *TextEditor) underlineText() {
	fmt.Println("Text has been underlined. ")
}

type BoldButton struct {
	textEditor TextEditor
}

func (bb *BoldButton) click() {
	bb.textEditor.boldText();
}

func main() {

	textEditor := TextEditor{}
	boldButton := BoldButton{textEditor: textEditor}
	
	boldButton.click()
}