package main

import (
    "example.com/note/note"
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

	userNote, err := note.New(title, content)

    if err != nil {
        fmt.Println(err)
		return
    }

	userNote.Display()
	err = userNote.Save()

	if err != nil {
		fmt.Println("Save failed.")
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
