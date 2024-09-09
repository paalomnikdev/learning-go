package main

import (
    "example.com/note/note"
	"example.com/note/todo"
    "fmt"
	"bufio"
    "os"
    "strings"
)

func getNoteData() (string, string) {
	title := getUserInput("Note title:")
	content := getUserInput("Note content:")

	return title, content
}

func main() {
	title, content := getNoteData()
	todoText := getUserInput("Todo text: ")

	todo, err := todo.New(todoText)

	if err != nil {
		fmt.Println("TODO error:\n")
		fmt.Println(err)
		return
	}

	userNote, err := note.New(title, content)

    if err != nil {
        fmt.Println(err)
		return
    }

	todo.Display()
	err = todo.Save()

	if err != nil {
		fmt.Println("TODO save failed.")
		return
	}

	fmt.Println("TODO saved.")

	userNote.Display()
	err = userNote.Save()

	if err != nil {
		fmt.Println("Note save failed.")
		return
	}

	fmt.Println("Note saved.")
}

func getUserInput(prompt string) (string) {
	fmt.Print(prompt)

	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}
