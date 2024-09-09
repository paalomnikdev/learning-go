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

type Saver interface {
	Save() error
}

type Displayer interface {
    Display()
}

type Outputtable interface {
    Saver
	Displayer
}

func main() {
	printSomething(1.1)
	printSomething("blah")
	printSomething(1)

	result := add(1, 2)
	result += 1

	fmt.Println("Generic result: ", result)

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
		fmt.Println("Note error:\n")
        fmt.Println(err)
		return
    }

	err = outputData(todo)

	if err != nil {
		return
	}

	outputData(userNote)
}

func add[T int | float64 | string](a, b T) T {
	return a + b
}

func printSomething(value any)  {
	typedVal, ok := value.(int)

	if ok {
		fmt.Println("INTEGER DETECTED!")
		fmt.Println(typedVal + 22)
	} else {
		fmt.Println("NOT INTEGER!")
	}

    switch value.(type) {
    case int:
		fmt.Println("Integer: ", value)
    case float64:
		fmt.Println("Float64: ", value)
    case string:
		fmt.Println("String: ", value)
    default:
        fmt.Println("Unacceptable!")
    }
}

func outputData(data Outputtable) error  {
	data.Display()
	return saveData(data)
}

func saveData(data Saver) error  {
    err := data.Save()

	if err != nil {
		fmt.Println("Data save failed.")
		return err
	}

	fmt.Println("Data saved.")
	return nil
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
