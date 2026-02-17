package main

import "fmt"

/*
A text editor where the user can undo changes, such as text addition,
or formatting. The editor stores snapshots of its state (text content)
after each change, enabling the user to revert to previous states
*/

// type TextEditor struct {
// 	content string 
// }

// func (te *TextEditor) Write(text string){
// 	te.content = text
// }

// func (te *TextEditor) getContent() string {
// 	return te.content;
// }


// func main(){
// 	editor := TextEditor{}

// 	editor.Write("Hello World !")
// 	editor.Write("Hello Everyone !")
// 	// problem -> Undo the last write !

// 	fmt.Println(editor.getContent())
// }


type TextEditor struct {
	content string 
}

func (te *TextEditor) Write(text string){
	te.content = text
}

func (te *TextEditor) getContent() string {
	return te.content;
}

func (te *TextEditor) Save() TextEditorMomento {
	return TextEditorMomento{content: te.content}
}


func main(){
	editor := TextEditor{}

	caretaker := Caretaker{}

	editor.Write("Hello World !")

	caretaker.saveState(&editor)
	editor.Write("Hello Everyone !")
	// problem -> Undo the last write !
	caretaker.saveState(&editor)

	// caretaker.undo(&editor)

	fmt.Println(editor.getContent())
}


type TextEditorMomento struct {
	content string
}

func (teM *TextEditorMomento) GetContent() string {
	return teM.content
}


type Caretaker struct {
	stack []TextEditorMomento
}

func (ct *Caretaker) saveState(textEditor *TextEditor) {
	ct.stack = append(ct.stack, textEditor.Save())
}

func (ct *Caretaker) undo(textEditor *TextEditor) {
	ct.stack = ct.stack[:len(ct.stack)-1]
	textEditor.Write(ct.stack[len(ct.stack)-1].GetContent())
}