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

type Command interface {
	execute()
}

type BoldCommand struct {
	textEditor TextEditor
}

func (bc *BoldCommand) execute() {
	bc.textEditor.boldText()
}

type Button struct {
	command Command
}

func (b *Button) click() {
	b.command.execute()
}

func main() {
	editor := TextEditor{}
	button := Button{&BoldCommand{ textEditor: editor}}
	button.click()
}