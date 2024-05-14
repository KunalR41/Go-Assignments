package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Memento Design pattern
type Memento interface{}

// Originator interface
type Originator interface {
	SetState(state string)
	GetState() string
	SaveToMemento() Memento
	RestoreFromMemento(memento Memento)
}

// Originator
type TextEditor struct {
	state string
}

func NewTextEditor() *TextEditor {
	return &TextEditor{}
}

func (t *TextEditor) SetState(state string) {
	t.state = state
}

func (t *TextEditor) GetState() string {
	return t.state
}

func (t *TextEditor) SaveToMemento() Memento {
	return t.state
}

func (t *TextEditor) RestoreFromMemento(memento Memento) {
	t.state = memento.(string)
}

// Caretaker
type History struct {
	mementos []Memento
}

func NewHistory() *History {
	return &History{
		mementos: make([]Memento, 0),
	}
}

func (h *History) Push(memento Memento) {
	h.mementos = append(h.mementos, memento)
}

func (h *History) Pop() Memento {
	if len(h.mementos) == 0 {
		return nil
	}
	lastIndex := len(h.mementos) - 1
	lastMemento := h.mementos[lastIndex]
	h.mementos = h.mementos[:lastIndex]
	return lastMemento
}

func main() {
	textEditor := NewTextEditor()
	history := NewHistory()

	textEditor.SetState("Initial state")
	history.Push(textEditor.SaveToMemento())

	textEditor.SetState("Modified state")
	history.Push(textEditor.SaveToMemento())

	previousState := history.Pop()
	if previousState != nil {
		textEditor.RestoreFromMemento(previousState)
	}

	http.HandleFunc("/state", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		response := map[string]string{"mesage": textEditor.GetState()}
		json.NewEncoder(w).Encode(response)
	})

	fmt.Println("Server is running on port 8888...")
	log.Fatal(http.ListenAndServe(":8888", nil))
}
