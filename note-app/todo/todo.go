package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

const storageDir = "./todos"

type Todo struct {
	Text string `json:"text"`
}

func (todo Todo) Display()  {
	fmt.Printf("TODO: %v  \n", todo.Text)
}

func (todo Todo) Save() error {
	fileName := strings.ToLower(fmt.Sprintf(
		"%s/%s.json",
		storageDir,
		time.Now().Format("20060102_150405"),
	))

	jsonData, err := json.Marshal(todo)

	if err != nil {
		return err
	}

	return os.WriteFile(fileName, jsonData, 06444)
}

func New(text string) (Todo, error) {
	if text == "" {
		return Todo{}, errors.New("invalid input")
	}

	return Todo{
		Text: text,
	}, nil
}
