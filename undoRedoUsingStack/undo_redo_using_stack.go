package undoredousingstack

import "fmt"

type Editor struct {
	Current   string
	UndoStack []string
	RedoStack []string
}

func (e *Editor) Perform(action string) {
	e.UndoStack = append(e.UndoStack, e.Current)
	e.RedoStack = []string{}
	e.Current = action
}

func (e *Editor) Undo() {
	if len(e.UndoStack) == 0 {
		fmt.Println("Nothing to Undo")
        return
	}
    
    e.RedoStack = append(e.RedoStack, e.Current)
    
    lastIndex := len(e.UndoStack) - 1
    e.Current = e.UndoStack[lastIndex]
    e.UndoStack = e.UndoStack[:lastIndex]
}

func (e *Editor) Redo(){
    if len(e.RedoStack) == 0{
        fmt.Println("Nothing to Redo")
        return
    }
    
    e.UndoStack = append(e.UndoStack, e.Current)
    
    lastIndex := len(e.RedoStack) - 1
    e.Current = e.RedoStack[lastIndex]
    e.RedoStack = e.RedoStack[:lastIndex]
}